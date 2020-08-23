package main

import (
	"log"
	"net"
	"net/http"
	"net/url"
	"runtime/debug"
	"bytes"
	"strings"
	"fmt"
	"bufio"
	"time"
	"io"
	"io/ioutil"
	"JCRandomProxy/Conf"
	"JCRandomProxy/Proxy"
)

func init(){
	log.SetFlags(log.LstdFlags|log.Lshortfile)
	Conf.InitConfig()
}

func main(){
	l,err := net.Listen("tcp",":8888")
	if err != nil {
		log.Panic(err)
	}

	for {
		client,err := l.Accept()
		if err != nil {
			log.Panic("Panic",err)
		}
		go handle(client)
	}
}

func handle(client net.Conn){
	defer func(){
		if err := recover(); err != nil {
			log.Panic(err)
			debug.PrintStack()
		}
	}()
	if client == nil {
		return
	}

	log.Println("JCTLog: client tcp tunnel connection: ",client.LocalAddr().String(),"->",client.RemoteAddr().String())
	defer client.Close()

	var b [1024]byte
	// 读取应用层的所有数据
	n,err := client.Read(b[:])
	if err != nil || bytes.IndexByte(b[:],'\n') == -1 {
		// 传输层的连接没有应用层的内容，如net.Dial()
		log.Println(err)
		return
	}
	var method,host,address string
	fmt.Sscanf(string(b[:bytes.IndexByte(b[:],'\n')]),"%s%s",&method,&host)
	log.Println(method,host)
	hostPortURL,err := url.Parse(host)
	if err != nil {
		log.Println(err)
		return
	}
	// https
	if hostPortURL.Opaque == "443" {
		address = hostPortURL.Scheme + ":443"
	}else{
		// http
		if strings.Index(hostPortURL.Host,":") == -1{
			address = hostPortURL.Host + ":80"
		}else{
			address = hostPortURL.Host
		}
	}
	log.Println("JCTLog: hostPortURL",address)
	server,err := Dial("tcp",address)
	if err != nil {
		log.Println("JCTLog: Dial: ",err)
		return
	}
	// 在应用层完成数据转发后，关闭传输层的通道
	defer server.Close()
	log.Println("JCTLog: server tcp tunnel connection: ",server.LocalAddr().String(),"->",server.RemoteAddr().String())
	
	if method == "CONNECT" {
		fmt.Fprint(client,"HTTP/1.1 200 Connection established\r\n\r\n")
	}else{
		log.Println("JCTLog: ","server write ",method)
		server.Write(b[:n])
	}
	// 进行转发
	go func(){
		log.Println("JCTLog: go转发前：")
		io.Copy(server,client)
		log.Println("JCTLog: go转发后：")
	}()
	log.Println("JCTLog: 开始转发：")
	io.Copy(client,server) // 
	// var tt []byte
	// tt,err = ioutil.ReadAll(server)
	// log.Println("ddddd",string(tt))
	// io.Copy(client,bytes.NewReader(tt[:]))
	log.Println("JCTLog: 结束： ")
}

// 建立一个传输通道
// network : 网络类型，tcp
// addr: 最终目标服务器地址
func Dial(network,addr string) (net.Conn,error){
	var proxyAddr string
	proxyAddr = "http://10.103.90.8:10080"
	proxyAddr = "http://49.4.123.243:8080"
	// 随机取出一个代理
	paddr,ptype,_ := Proxy.GetAProxy()
	proxyAddr = ptype + "://" + paddr

	// 建立到代理服务器的传输层通道
	c,err := func() (net.Conn,error){
		prox,_ := url.Parse(proxyAddr)
		log.Println("JCTLog: 代理地址: ",prox.Host)
		// Dial and create client connection
		proxc,err := net.DialTimeout("tcp",prox.Host,time.Second * 5)
		if err != nil {
			return nil,err
		}
		// 在这里返回c可以正常使用
		if CheckProxy(proxyAddr,addr) {
			return proxc,err
		}
		return nil,err
		
	}()
	if c == nil || err != nil {
		log.Println("JCTLog: 代理异常: ",c,err)
		// log.Println("JCTLog: 本地直接转发: ")
		return net.Dial(network,addr)
	}
	log.Println("JCTLog: 代理正常，tunnel信息 ",c.LocalAddr().String(),"->",c.RemoteAddr().String())
	return c,err
}

// 验证代理服务器是否可用
func CheckProxy(proxyAddr,addr string) bool{
		prox,_ := url.Parse(proxyAddr)
		log.Println("JCTLog: 代理地址: ",prox.Host)
		// Dial and create client connection
		proxc,err := net.DialTimeout("tcp",prox.Host,time.Second * 5)
		if err != nil {
			return false
		}
		// 解析最终目标url
		reqURL ,err := url.Parse("http://"+addr)
		if err != nil {
			return false
		}
		log.Println("JCTLog: reqURL: " ,reqURL.String())
		req,err := http.NewRequest(http.MethodGet,reqURL.String(),nil)
		if err != nil {
			log.Println("JCTLog: http.NewRequest: ",err)
			return false
		}

		req.Close = false
		req.Header.Set("User-Agent","Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.3")
		err = req.Write(proxc)
		fmt.Println(req)
		if err != nil {
			log.Println("JCTLog: req.Write: ",err)
			return false
		}
		
		resp,err := http.ReadResponse(bufio.NewReader(proxc),req)
		if err != nil {
			log.Println("JCTLog: http.ReadResponse: ",err)
			return false
		}
		defer resp.Body.Close()	
		fmt.Println("===================sss")
		fmt.Println(resp.Body)
		fmt.Println(resp.StatusCode)
		fmt.Println(resp.Status)
		fmt.Println(resp.Proto)
		fmt.Println(resp.Header)

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return false
		}

		fmt.Println(string(body))
		fmt.Println("===================eee")
		defer resp.Body.Close()		
		if resp.StatusCode != 200 {
			err = fmt.Errorf("Connect server using proxy error,StatusCode [%d]",resp.StatusCode)
			return false
		}
		return true

}