package network

import (
	"log"
	"net"
	"testing"
	"time"
)

func TestUdpRead(t *testing.T) {
	c, err := net.ListenPacket("udp", ":8181")
	if err != nil {
		log.Panic(err)
	}
	// 不设置超时时间会一直堵塞
	now := time.Now()
	c.SetReadDeadline(now.Add(10 * time.Second))
	log.Printf("time:%d\n", time.Now().Unix())
	buf := make([]byte, 1024)
	_, addr, err := c.ReadFrom(buf)
	if err != nil {
		log.Panic(err)
	}
	log.Println(addr)
	// 超时时间可以重新设置
	// 底层参考 runtime.poll_runtime_pollSetDeadline
	c.SetReadDeadline(time.Now().Add(10 * time.Second))
	_, addr, err = c.ReadFrom(buf)
	if err != nil {
		log.Printf("time:%d\n", time.Now().Unix())
		log.Panic(err)
	}
	log.Println(addr)
}

func TestUdpWrite(t *testing.T) {
	addr, err := net.ResolveUDPAddr("udp", "255.255.255.255:8181")
	if err != nil {
		return
	}
	c, err := net.ListenPacket("udp", ":0")
	if err != nil {
		log.Panic(err)
	}
	n, err := c.WriteTo([]byte{0, 0, 0, 0, 0}, addr)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("write ok:%d", n)
}
