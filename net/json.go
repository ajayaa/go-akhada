package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

type message struct {
	Type string
}

type members struct {
	lock        sync.RWMutex
	member_list map[net.Addr]time.Time
}

func (m *members) add_member(addr net.Addr) {
	m.lock.Lock()
	m.member_list[addr] = time.Now()
	m.lock.Unlock()
}

func (m *members) remove_member(addr net.Addr) {
	m.lock.Lock()
	delete(m.member_list, addr)
	m.lock.Unlock()
}

func (m *members) print() {
	m.lock.RLock()
	for addr := range m.member_list {
		fmt.Println(addr.String(), m.member_list[addr])
	}
}

type join_message struct {
	Type string
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s port-number\n", os.Args[0])
		os.Exit(1)
	}
	m := members{member_list: make(map[net.Addr]time.Time)}
	port := os.Args[1]
	port = fmt.Sprintf(":%s", port)
	ln, err := net.Listen("tcp", port)
	if err != nil {
		panic("port bound already")
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic("some error!")
		}
		go handleConnection(conn, m)
	}
}

func handleConnection(conn net.Conn, m members) {
	defer conn.Close()
	m.add_member(conn.RemoteAddr())
	m.print()
	remoteAddr := conn.RemoteAddr().String()
	fmt.Fprintf(os.Stdout, "conn opened from: %s\n", remoteAddr)
	var buffer = make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Fprintf(os.Stdout, "conn closed from: %s\n", remoteAddr)
			return
		}
		mes := message{}
		json.Unmarshal(buffer[:n], &mes)
		fmt.Fprintf(os.Stdout, "%s%s%s", remoteAddr, ": ", mes)
		if mes.Type == "ping" {
			return_msg, _ := json.Marshal(message{Type: "pong"})
			conn.Write([]byte(return_msg))
		}
	}
	m.remove_member(conn.RemoteAddr())
	m.print()
}
