package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

// ★型の違う引数を可変長で渡すテクニック★

// ex. Goも可変長で引数を渡すことができるが、型が固定されてしまう
func manyArgs(args ...string) {
	for _, arg := range args {
		fmt.Print(arg)
	}
}

type Server1 struct {
	host    string
	port    int
	timeout time.Duration
	logger  *log.Logger
}

func (s *Server1) Start() error {
	// server start
	return nil
}

// timeoutを受け取り設定・更新する関数をクロージャとして返す
func WithTimeout(timeout time.Duration) func(*Server1) {
	return func(s *Server1) {
		s.timeout = timeout
	}
}

// loggerを受け取り設定・更新する関数をクロージャとして返す
func WithLogger(logger *log.Logger) func(*Server1) {
	return func(s *Server1) {
		s.logger = logger
	}
}

// Optionを定義することで、可変長部分の型をanyにしないで済む
// anyにしないことで開発者が限定する型を可変個で渡せるようになる
type Option func(*Server1)

func New(host string, port int, options ...Option) *Server1 {
	svr := &Server1{
		host: host,
		port: port,
	}
	for _, opt := range options {
		opt(svr)
	}
	return svr
}

func FunctionalOptions() {
	manyArgs("Hello,", "World", "!!!")

	f, err := os.Create("server.log")
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	logger := log.New(f, "", log.LstdFlags)
	svr := New("localhost", 8888,
		WithTimeout(time.Minute),
		WithLogger(logger))
	if err := svr.Start(); err != nil {
		log.Println(err)
	}
}
