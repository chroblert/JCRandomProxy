package ProxyEntry

import (
	"log"
	"net"
	"runtime/debug"

	"../Conf"
	"../Proxy"
)

/**
*
* Author: JC0o0l
* email: jerryzvs@163.com
* wechat: JC_SecNotes
 */

func Proxymain(stop chan int) {
	// 监听TCP连接
	l, err := net.Listen("tcp", ":"+Conf.Port)
	log.Println("监听在：", Conf.Port)
	if err != nil {
		log.Panic(err)
	}

	for {

		// 接收停止信号
		select {
		case <-stop:
			log.Println("收到停止信号")
			l.Close()
			stop <- 1
			return
		default:
		}
		// 接收TCP连接，返回一个net.Conn
		client, err := l.Accept()
		if err != nil {
			log.Panic("Panic", err)
		}
		// 收到请求后，调用handle进行处理
		go handle(client)
	}
}

func handle(client net.Conn) {
	defer client.Close()
	defer func() {
		if err := recover(); err != nil {
			debug.PrintStack()
			log.Panic(err)
		}
	}()
	if client == nil {
		return
	}

	log.Println("JCTLog: client tcp tunnel connection: ", client.LocalAddr().String(), "->", client.RemoteAddr().String())
	/* 20200923: 删除该段
	// 随机取出一个代理
	paddr, ptype, _ := Proxy.GetAProxy()
	proxyAddr := ptype + "://" + paddr
	// 验证代理是否有效
	checkaddr := "https://myip.ipip.net"
	if Proxy.CheckProxy(proxyAddr, checkaddr) {
		log.Println(" 代理有效 ", proxyAddr)
		// 有效，使用端口转发
		Rproxy(client, paddr)
	} else {
		log.Println(" 代理无效 ", proxyAddr)
		// 判断该代理是否在可用代理池，若在，则删除
		// 无效，使用自身代理
		Lproxy(client)
	}
	*/
	// 使用代理
	visit(client)
}

// 20200923: 将使用代理独立出来
func visit(client net.Conn) {
	// 取出一个代理
	aproxyaddr, ok := Proxy.MSafeProxymap.GetARandProxy()
	// 取出代理失败，则使用本地代理
	if !ok {
		Lproxy(client)
	} else {
		// 取出代理成功，则使用可用代理
		protocol := aproxyaddr.Protocol
		ip := aproxyaddr.Ip
		port := aproxyaddr.Port
		paddr := protocol + "://" + ip + ":" + port
		Rproxy(client, paddr)
	}
}
