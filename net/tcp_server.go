package net

import (
	"fmt"
	"net"

	"github.com/sky1309/gou/log"
)

type OnConnect func(c *Conn)

type OnDisconnect func(c *Conn)

type OnReceive func(c *Conn, data []byte)

type TcpServer struct {
	onConnect    OnConnect
	onDisconnect OnDisconnect
	onReceive    OnReceive
}

func (s *TcpServer) Init(onConnect OnConnect, onDisconnect OnDisconnect, onReceive OnReceive) {
	s.onConnect = onConnect
	s.onDisconnect = onDisconnect
	s.onReceive = onReceive
}

func (s *TcpServer) Listen(port int) error {
	log.Info("listen tcp port %d", port)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	go func() {
		for {
			c, e := listener.Accept()
			if e != nil {
				log.Error("socket listen accept err %s", err)
				continue
			}
			log.Info("new conn accepted")
			conn := &Conn{c: c}
			s.onConnect(conn)

			go s.handleConn(conn)
		}
	}()
	return nil
}

func (s *TcpServer) handleConn(conn *Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Error("conn read err %s", err)
			break
		}
		s.onReceive(conn, buf[:n])
	}
	s.onDisconnect(conn)
}
