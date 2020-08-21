package main
import (    
	"bytes"
    "fmt"
    "io"
	"log"
    "net"
    "net/url"
	"strings"
	"errors"
	"JCRandomProxy/Conf"
	"JCRandomProxy/Proxy"
)
var EOF = errors.New("EOF")
var ErrShortWrite = errors.New("short write")
func main() {
	Conf.InitConfig()
	log.SetFlags(log.LstdFlags|log.Lshortfile)
	
	l, err := net.Listen("tcp", ":8081")    
	if err != nil {
        log.Panic(err)
	}    
	
	for {
		client, err := l.Accept()        
		if err != nil {
            log.Panic(err)
		}        
		go handleClientRequest(client)
	}

}
func handleClientRequest(client net.Conn) {    
	if client == nil {        
		log.Printf("没有client接入")
		return
	}    
	defer client.Close()    
	var b [1024]byte
	n, err := client.Read(b[:])  
	// log.Println("JCTLog ALL: \n",string(b[:n]))  
	if err != nil {
		log.Println(err)       
		return
	}    
	// 获取一个随机代理，代理类型
	
	proxy,ptype,_ := Proxy.GetAProxy()
	// proxy,ptype := "1113.100.209.65:3128","http"
	log.Printf("JCTLog ALL: 获取一个随机代理: %s %s",ptype,proxy)
	var method, host, address string
	// 以空格分割，读取client请求的第一行的方法与主机
	fmt.Sscanf(string(b[:bytes.IndexByte(b[:], '\n')]), "%s%s", &method, &host)
	log.Println("JCTLog ALL: Method Host ",method,host)
	log.Println("JCTLog ALL: Request :\n",string(b[:n]))
	hostPortURL, err := url.Parse(host)    
	if err != nil {
		log.Println(err)        
		return
	}    
	log.Println("JCTLog ALL: 目标主机:端口  ",hostPortURL)
	if hostPortURL.Opaque == "443" { //https访问
        address = hostPortURL.Scheme + ":443"
    } else { //http访问
        if strings.Index(hostPortURL.Host, ":") == -1 { //host不带端口， 默认80
            address = hostPortURL.Host + ":80"
        } else {
            address = hostPortURL.Host
        }
	}    
	address = proxy
	// address = "113.100.209.65:3128"
	log.Printf("JCTLog ALL: 建立到上级代理%s的TCP连接",address)   
	server,err :=net.Dial("tcp",address)
	if err != nil {
		log.Println(err)
		log.Println("JCTLog ALL: 建立TCP连接失败！")        
		return
	}    
	log.Printf("JCTLog ALL: 到上级代理%s的TCP连接建立成功",address)
	if method == "CONNECT" {
		// https请求
		log.Println("JCTLog HTTPS: 这是一个HTTPS请求")
		fmt.Fprint(client, "HTTP/1.1 200 Connection established\r\n\r\n")
		go func(){
			/**
			* ??? io.Copy(server,client)返回的值是498
			*    而io.Copy(&buf,client)返回的值是195
			*    为何会有不同?
			*
			*/
			log.Printf("JCTLog HTTPS: HTTPS请求内容为：\n%s",string(b[:n]))
			// 将client直接转发给server
			n,err := io.Copy(server,client)
			// n,err :=  (server).(*net.TCPConn).ReadFrom(client)
			fmt.Println("JCTLog client.Read Length ",n)
			if err != nil {
				if err != io.EOF {
					fmt.Println("client.Read error:", err)
				}
			}
		}()
   		io.Copy(client, server)
    } else {
		// // http请求
		// log.Printf("JCTLog HTTP: 向上级代理发送HTTP请求")
		// server.Write(b[:n])
		// fmt.Printf("%d",n)
		// log.Printf("JCTLog HTTP: HTTP请求内容为：\n%s",string(b[:n]))
		// log.Printf("JCTLog HTTP: 将上级代理的响应内容转发到client端")
		// n, err = server.Read(b[:]) 
		// log.Printf("JCTLog HTTP: 上级代理的响应内容为：\n%s",string(b[:n]))
		// io.Copy(client, bytes.NewReader(b[:]))
		// // n, err = server.Read(b[:]) 
		// // log.Printf("JCTLog: 上级代理的响应内容如为2：%s",string(b[:n]))
		// log.Printf("JCTLog HTTP: 完毕")



		go func(){
			log.Printf("JCTLog HTTP: HTTP请求内容为：\n%s",string(b[:n]))
			// 将client直接转发给server
			n,err := io.Copy(server,client)
			// n,err :=  (server).(*net.TCPConn).ReadFrom(client)
			fmt.Println("JCTLog client.Read Length ",n)
			if err != nil {
				if err != io.EOF {
					fmt.Println("client.Read error:", err)
				}
			}
		}()
   		io.Copy(client, server)

	}    
    
}



