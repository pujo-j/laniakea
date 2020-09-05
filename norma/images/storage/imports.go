package main

import (
	_ "github.com/dgraph-io/ristretto"
	_ "github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/pgxpool"
	_ "go.uber.org/zap"
	_ "gocloud.dev/blob"
	_ "gocloud.dev/blob/azureblob"
	_ "gocloud.dev/blob/gcsblob"
	_ "gocloud.dev/blob/s3blob"
)
