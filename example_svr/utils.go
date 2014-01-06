package main

import (
	"fmt"
	"net"
)

func Init(ip string, listenPortBase int) error {
	var err error

	listenConn, err = CreateUdpConn(ip + ":" + fmt.Sprintf("%v", listenPortBase))
	if err != nil {
		return err
	}

	return nil
}

func CreateUdpConn(addrStr string) (*net.UDPConn, error) {
	addr, err := net.ResolveUDPAddr("udp", addrStr)
	if err != nil {
		return nil, err
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
