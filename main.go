package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
)

func main() {
	if l := GetSysLogger(); l == nil {
		log.Fatal("init syslog err")
		panic(0)
	}

	go StartReceiver()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Println("Shutdown Server ...")
}
