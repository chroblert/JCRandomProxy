package main

import (
	"log"
	"net"
)

func PortForward(client net.Conn, targetaddr string) {
	// Read a header firstly in case you could have opportunity to check request
	// whether to decline or proceed the request
	buffer := make([]byte, 1024)
	n, err := client.Read(buffer)
	if err != nil {
		log.Printf("Unable to read from input, error: %s\n", err.Error())
		return
	}
	// targetaddr = "223.82.106.253:3128"
	targetconn, err := net.Dial("tcp", targetaddr)
	if err != nil {
		log.Println("Unable to connect to: %s, error: %s\n", targetaddr, err.Error())
		client.Close()
		return
	}
	n, err = targetconn.Write(buffer[:n])
	if err != nil {
		log.Printf("Unable to write to output, error: %s\n", err.Error())
		client.Close()
		targetconn.Close()
		return
	}
	go proxyRequest(client, targetconn)
	proxyRequest(targetconn, client)
}
