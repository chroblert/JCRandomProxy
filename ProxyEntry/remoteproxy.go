package ProxyEntry

import (
	log "../Logs"
	"net"
	"time"

	"../Conf"
)

func Rproxy(client net.Conn, targetaddr string) {
	// Read a header firstly in case you could have opportunity to check request
	// whether to decline or proceed the request
	defer client.Close()
	buffer := make([]byte, 1024)
	n, err := client.Read(buffer)
	if err != nil {
		log.Printf("Unable to read from input, error: %s\n", err.Error())
		return
	}
	// targetaddr = "223.82.106.253:3128"
	// 20200922: 使用带有超时的拨号
	targetconn, err := net.DialTimeout("tcp", targetaddr, time.Duration(Conf.Timeout)*time.Second)
	if err != nil {
		log.Printf("Unable to connect to: %s, error: %s\n", targetaddr, err.Error())
		return
	}
	defer targetconn.Close()
	n, err = targetconn.Write(buffer[:n])
	if err != nil {
		log.Printf("Unable to write to output, error: %s\n", err.Error())
		return
	}
	go proxyRequest(client, targetconn)
	proxyRequest(targetconn, client)
}
