package main

import (
	"context"
	"github.com/yuxing/hellomodule/bookstore/server"
	"github.com/yuxing/hellomodule/bookstore/store/factory"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// go 中，main包是整个程序的入口
// 除此之外，main包还是整个程序主要模块初始化与组装的场所
func main() {
	s, err := factory.New("mem")
	if err != nil {
		panic(err)
	}

	srv := server.NewBookStoreServer(":8080", s)
	errChan, err := srv.ListenAndServe()
	if err != nil {
		log.Panicln("web server start failed: ", err)
		return
	}
	log.Panicln("web server start ok")

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err = <-errChan:
		log.Println("web server run failed: ", err)
		return
	case <- c:
		log.Println("bookstore program is exiting...")
		ctx, cf := context.WithTimeout(context.Background(), time.Second)
		defer cf()
		err = srv.Shutdown(ctx)
	}

	if err != nil {
		log.Println("bookstore program exit error:", err)
		return
	}

	log.Println("bookstore program exit ok")
}
