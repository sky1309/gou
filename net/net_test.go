package net

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"

	"github.com/sky1309/gou/log"
)

func TestSocket(t *testing.T) {
	s := Socket{}
	s.Init(func(c *Conn) {
		log.Info("new connection")
	}, func(c *Conn) {
		log.Info("connection disconnect")
	}, func(c *Conn, data []byte) {
		log.Info("connection receive data %v", string(data))
	})

	s.Listen(8000)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	<-c
	log.Info("socket over")
}

func TestClient(t *testing.T) {
	c, err := net.Dial("tcp", ":8000")
	if err != nil {
		log.Error(err.Error())
		return
	}
	for i := 0; i < 5; i++ {
		c.Write([]byte(fmt.Sprintf("abcde%d", i)))
		time.Sleep(time.Second)
	}
}
