package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/tg123/go-htpasswd"
	"go.uber.org/zap"
	"os"
	"strings"
	"time"
)

var sharedKey []byte
var envName string
var envColor string

func Id() func(c *gin.Context) {
	var authMethodsNames []string
	for _, method := range authOptions {
		authMethodsNames = append(authMethodsNames, method.Name())
	}
	authMethodsJson, err := json.Marshal(authMethodsNames)
	if err != nil {
		logger.Fatal("Impossible to create authMethodsJson", zap.Error(err))
	}
	response := []byte("window.AUTH = " + string(authMethodsJson) + "; window.ENV= {name:'" + envName + "',color:'" + envColor + "'}")
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/javascript")
		c.Writer.WriteHeader(200)
		c.Writer.Write(response)
	}
}

type AuthOption interface {
	Name() string
	GetToken(c *gin.Context) (string, error)
}

var authOptions = make(map[string]AuthOption, 0)
var tokenOption = &BareTokenAuth{}

func initAuth() {
	param := os.Getenv("AUTH_TOKEN_KEY")
	if param == "" {
		logger.Fatal("No auth token key provided")
	}
	buf, err := base64.RawStdEncoding.DecodeString(param)
	if err != nil || len(buf) != 32 {
		logger.Fatal("invalid token key defined")
	}
	sharedKey = buf
	authOptions[tokenOption.Name()] = tokenOption
	simpleloginEnv := os.Getenv("SIMPLE_AUTH")
	if simpleloginEnv != "" {
		htp, err := htpasswd.NewFromReader(strings.NewReader(simpleloginEnv), htpasswd.DefaultSystems, nil)
		if err != nil {
			logger.Error("reading SIMPLE_AUTH variable", zap.Error(err))
		} else {
			sl := &SimpleLogin{file: htp}
			logger.Debug("Loaded htpasswd from " + simpleloginEnv)
			authOptions[sl.Name()] = sl
		}
	}
	envName = os.Getenv("ENV")
	envColor = os.Getenv("ENV_COLOR")
	if envName == "" {
		envName = "DEV"
	}
	if envColor == "" {
		envColor = "green"
	}
}

const ErrInvalidAuth = StringError("invalid auth")

func generateToken(user string, duration time.Duration) string {
	toSign := make([]byte, 8, len(user)+8)
	expire := time.Now().Add(duration)
	binary.LittleEndian.PutUint64(toSign[:8], uint64(expire.Unix()))
	toSign = append(toSign, []byte(user)...)
	hash := hmac.New(sha256.New, sharedKey)
	hash.Write(toSign)
	mac := hash.Sum(nil)
	tokenBytes := make([]byte, 40, len(user)+40)
	binary.LittleEndian.PutUint64(tokenBytes[:8], uint64(expire.Unix()))
	copy(tokenBytes[8:40], mac)
	tokenBytes = append(tokenBytes, []byte(user)...)
	return base64.RawStdEncoding.EncodeToString(tokenBytes)
}

func validateToken(token string) error {
	bytes, err := base64.RawStdEncoding.DecodeString(token)
	if err != nil {
		return ErrInvalidAuth
	}
	if len(bytes) < 44 {
		return ErrInvalidAuth
	}
	exp := binary.LittleEndian.Uint64(bytes[0:8])
	expt := time.Unix(int64(exp), 0)
	mac1 := bytes[8:40]
	id := bytes[40:]
	logger.Debug("Trying to authenticate", zap.String("id", string(id)), zap.String("expiry", expt.String()))
	hash := hmac.New(sha256.New, sharedKey)
	hash.Write(bytes[0:8])
	hash.Write(id)
	mac2 := hash.Sum(nil)
	if hmac.Equal(mac1, mac2) {
		logger.Debug("Trying to authenticate")
		if expt.After(time.Now()) {
			return nil
		} else {
			return ErrInvalidAuth
		}
	} else {
		return ErrInvalidAuth
	}
}

func AuthMiddleware(c *gin.Context) {
	var authHeader = c.Request.Header.Get("authorization")
	if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if validateToken(token) == nil {
			c.Next()
			return
		}
	}
	c.Writer.WriteHeader(403)
	c.Abort()
}

func Login(c *gin.Context) {
	loginMethodName := c.PostForm("method")
	method, ok := authOptions[loginMethodName]
	if !ok {
		logger.Warn("Invalid login method", zap.String("method", loginMethodName))
		c.Writer.WriteHeader(403)
		return
	}
	logger.Debug("Trying login with method: " + loginMethodName)
	token, err := method.GetToken(c)
	if err != nil {
		logger.Debug("Invalid login", zap.Error(err))
		c.Writer.WriteHeader(403)
		return
	}
	logger.Debug("Valid login")
	c.Writer.WriteHeader(200)
	c.Writer.Write([]byte(token))
}

//**** Bare Laniakea Token ****

type BareTokenAuth struct {
}

func (b *BareTokenAuth) Name() string {
	return "BareToken"
}

func (b *BareTokenAuth) GetToken(c *gin.Context) (string, error) {
	var token = c.PostForm("token")
	if token != "" {
		if validateToken(token) == nil {
			return token, nil
		}
		return "", ErrInvalidAuth
	}
	return "", ErrInvalidAuth
}

//**** Simple login/password ****

type SimpleLogin struct {
	file *htpasswd.File
}

func (s *SimpleLogin) Name() string {
	return "SimpleLogin"
}

func (s *SimpleLogin) GetToken(c *gin.Context) (string, error) {
	l := c.PostForm("login")
	p := c.PostForm("password")
	if l != "" && p != "" {
		if s.file.Match(l, p) {
			return generateToken(l, 24*time.Hour), nil
		}
	}
	return "", ErrInvalidAuth
}
