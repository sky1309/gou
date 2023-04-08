package net

import (
	"net"
	"time"
)

type Conn struct {
	c net.Conn
}

func (c *Conn) Read(b []byte) (n int, err error) {
	return c.c.Read(b)
}

func (c *Conn) Write(b []byte) (n int, err error) {
	return c.c.Write(b)
}

func (c *Conn) Close() error {
	return c.c.Close()
}

func (c *Conn) LocalAddr() net.Addr {
	return c.c.LocalAddr()
}

func (c *Conn) SetDeadline(t time.Time) error {
	return c.c.SetDeadline(t)
}

func (c *Conn) SetReadDeadline(t time.Time) error {
	return c.c.SetReadDeadline(t)
}

func (c *Conn) SetWriteDeadline(t time.Time) error {
	return c.c.SetWriteDeadline(t)
}
