package main

import (
	"github.com/lee-lou2/msg/api/rest"
	"github.com/lee-lou2/msg/cmd/collector"
	"github.com/lee-lou2/msg/cmd/dispatcher"
	"github.com/lee-lou2/msg/cmd/sender"
	"github.com/lee-lou2/msg/configs"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 기본 설정
	if err := configs.LoadEnvs(); err != nil {
		log.Fatalln(err)
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	// 프로그램 종료시 실행될 코드
	go func() {
		<-c
		log.Println("👋 프로그램이 종료되었습니다.")
		os.Exit(1)
	}()

	// 프로그램 실행
	go collector.Run()
	go dispatcher.Run()
	go sender.Run()
	rest.Run()
}
