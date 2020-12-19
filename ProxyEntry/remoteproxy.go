package ProxyEntry

import (
	"encoding/base64"
	"net"
	"strings"
	"time"

	log "github.com/chroblert/JCRandomProxy/Logs"

	"github.com/chroblert/JCRandomProxy/Conf"
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
	// 2020/12/19: 增加是否开启认证
	if Conf.EnableAuth {
		log.Printf("JCDebug32: %s", string(buffer[:n]))
		// 2020/12/18: 计划增加认证功能
		strHttpReq := string(buffer[:n])
		if strings.Contains(strHttpReq, "Proxy-Authorization") {
			authString := Conf.ProxyUser + ":" + Conf.ProxyPasswd
			encodeString := base64.StdEncoding.EncodeToString([]byte(authString))
			if !(strings.Contains(strHttpReq, encodeString)) {
				log.Printf("JCTest:认证失败\n")
				client.Write([]byte("JCTest: Authorization Failure"))
				return
			}
		} else {
			return
		}
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
