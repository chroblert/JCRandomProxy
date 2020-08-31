package main
import (
	"net"
	"log"
	"runtime/debug"
	"fmt"
	"bytes"
	"net/url"
	"strings"
)

func lproxy(client net.Conn){
	defer func() {
		if err := recover(); err != nil {
			log.Panic(err)
			debug.PrintStack()
		}
	}()
	if client == nil {
		return
	}
	defer client.Close()

	var b [1024]byte
	// 读取应用层的所有数据
	n, err := client.Read(b[:])
	if err != nil || bytes.IndexByte(b[:], '\n') == -1 {
		// 传输层的连接没有应用层的内容，如net.Dial()
		log.Println(err)
		return
	}
	var method, host, address string
	fmt.Sscanf(string(b[:bytes.IndexByte(b[:], '\n')]), "%s%s", &method, &host)
	log.Println(method, host)
	hostPortURL, err := url.Parse(host)
	if err != nil {
		log.Println(err)
		return
	}
	// https
	if hostPortURL.Opaque == "443" {
		address = hostPortURL.Scheme + ":443"
	} else {
		// http
		if strings.Index(hostPortURL.Host, ":") == -1 {
			address = hostPortURL.Host + ":80"
		} else {
			address = hostPortURL.Host
		}
	}
	log.Println("JCTLog: hostPortURL", address)
	// 建立一个到代理服务器的传输通道
	server, err := Dial("tcp", address)
	if err != nil {
		log.Println("JCTLog: Dial: ", err)
		return
	}
	// 在应用层完成数据转发后，关闭传输层的通道
	defer server.Close()
	log.Println("JCTLog: server tcp tunnel connection: ", server.LocalAddr().String(), "->", server.RemoteAddr().String())

	if method == "CONNECT" {
		// https
		fmt.Fprint(client, "HTTP/1.1 200 Connection Established\r\n\r\n")
	} else {
		// http
		// log.Println("JCTLog: ", "server write ", method)
		server.Write(b[:n])
	}
	// 进行转发
	go func() {
		proxyRequest(client,server)
	}()
	proxyRequest(server,client)
	log.Println("JCTLog: 结束： ")
}


// 建立一个传输通道
// network : 网络类型，tcp
// addr: 最终目标服务器地址
func Dial(network, addr string) (net.Conn, error) {
	return net.Dial(network,addr)
}

