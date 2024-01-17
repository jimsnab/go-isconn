package main

import (
	"fmt"
	"net"
	"os"
	"syscall"
	"time"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		panic("isconn [cli|srv]")
	}

	if args[1] == "cli" {
		client()
	} else {
		server()
	}
}

func isConnected(conn net.Conn) bool {
	f, err := conn.(*net.TCPConn).File()
	if err != nil {
		return false
	}

	b := []byte{0}
	_, _, err = syscall.Recvfrom(int(f.Fd()), b, syscall.MSG_PEEK|syscall.MSG_DONTWAIT)
	return err != nil
}

func server() {
	listener, err := net.Listen("tcp", "localhost:7500")
	if err != nil {
		panic(err)
	}
	fmt.Println("accepting one connection")
	conn, err := listener.Accept()
	if err != nil {
		panic(err)
	}

	for {
		fmt.Printf("server:%t\n", isConnected(conn))
		time.Sleep(time.Second)
	}
}

func client() {
	conn, err := net.Dial("tcp", "localhost:7500")
	if err != nil {
		panic(err)
	}

	for {
		fmt.Println("about to read")
		conn.SetReadDeadline(time.Now().Add(time.Millisecond * 250))
		b := []byte{0}
		fmt.Println("reading")
		n, err := conn.Read(b)
		fmt.Println("read", n, err)
		fmt.Printf("client:%t\n", isConnected(conn))
	}
}
