package ProxyEntry

import (
	log "../Logs"
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
		log.Println(err)
		return
	}
	for {
		// 接收停止信号
		select {
		case <-stop:
			log.Println("收到停止信号")
			l.Close()
			// stop <- 1
			return
		default:
		}
		// 接收TCP连接，返回一个net.Conn
		client, err := l.Accept()
		if err != nil {
			log.Println("Panic", err)
			return
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
			log.Println(err)
			return
		}
	}()
	if client == nil {
		return
	}

	log.Println("JCTLog: client tcp tunnel connection: ", client.LocalAddr().String(), "->", client.RemoteAddr().String())
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
		// protocol := aproxyaddr.Protocol
		ip := aproxyaddr.Ip
		port := aproxyaddr.Port
		// paddr := protocol + "://" + ip + ":" + port
		paddr := ip + ":" + port
		Rproxy(client, paddr)
	}
}
