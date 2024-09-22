package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Server struct {
	param serverParam
}

type serverParam struct {
	host    string
	port    int
	timeout time.Duration
	logger  *log.Logger
}

func NewBuilder(host string, port int) *serverParam {
	return &serverParam{
		host: host,
		port: port,
	}
}

func (sb *serverParam) Timeout(timeout time.Duration) *serverParam {
	sb.timeout = timeout
	return sb
}

func (sb *serverParam) Logger(logger *log.Logger) *serverParam {
	sb.logger = logger
	return sb
}

// ここでようやくServerを返す
func (sb *serverParam) Build() *Server {
	return &Server{
		param: *sb,
	}
}

func (s *Server) Start() error {
	if s.param.logger != nil {
		s.param.logger.Println("server started")
		fmt.Println("server started")
	}
	return nil
}

func BuilderPattern() {
	f, err := os.Create("server.log")
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	logger := log.New(f, "", log.LstdFlags)
	// BuilderPatternではメソッドチェーンを使用するが、Goではメソッドチェーンはあまり使用されないらしい
	svr := NewBuilder("localhost", 8888).
		Timeout(time.Minute).
		Logger(logger).
		Build()
	if err = svr.Start(); err != nil {
		log.Panicln(err)
	}
}
