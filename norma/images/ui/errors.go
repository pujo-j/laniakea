package main

type StringError string

func (s StringError) Error() string {
	return string(s)
}
