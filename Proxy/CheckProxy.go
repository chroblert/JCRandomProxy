package Proxy

import (
	"JCRandomProxy/Conf"
	"bufio"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// 验证代理服务器是否可用
func CheckProxy(proxyAddr, checkaddr string) bool {
	proxymd5 := fmt.Sprintf("%x", md5.Sum([]byte(proxyAddr)))
	// 代理服务器可用，则添加至map中
	if CheckProxyA(proxyAddr, checkaddr) {
		protocol := strings.Split(proxyAddr, ":")[0]
		ip := strings.Split(strings.Split(proxyAddr, "/")[2], ":")[0]
		port := strings.Split(proxyAddr, ":")[2]
		tmpproxy := aproxy{protocol, ip, port}
		// proxylist = append(proxylist, tmpproxy)
		proxymap[proxymd5] = tmpproxy
		log.Println(proxymap)
		return true
	}
	// 代理服务器不可用，则删除
	delete(proxymap, proxymd5)
	delete(metaproxymap, proxymd5)
	return false
}
func CheckProxyA(proxyAddr, checkaddr string) bool {
	// if !Conf.UseProxyPool {
	// 	return true
	// }

	prox, _ := url.Parse(proxyAddr)
	log.Println("JCTLog: 代理地址: ", prox.Host)
	// Dial and create client connection
	proxc, err := net.DialTimeout("tcp", prox.Host, time.Second*5)
	if err != nil {
		return false
	}
	// 解析最终目标url
	reqURL, err := url.Parse(checkaddr)
	if err != nil {
		return false
	}
	log.Println("JCTLog: reqURL: ", reqURL.String())
	req, err := http.NewRequest(http.MethodGet, reqURL.String(), nil)
	if err != nil {
		log.Println("JCTLog: http.NewRequest: ", err)
		return false
	}

	req.Close = false
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.3")
	err = req.Write(proxc)
	fmt.Println(req)
	if err != nil {
		log.Println("JCTLog: req.Write: ", err)
		return false
	}

	resp, err := http.ReadResponse(bufio.NewReader(proxc), req)
	if err != nil {
		log.Println("JCTLog: http.ReadResponse: ", err)
		return false
	}
	defer resp.Body.Close()
	fmt.Println("===================sss")
	// fmt.Println(resp.Body)
	// fmt.Println(resp.StatusCode)
	fmt.Println(resp.Status)
	// fmt.Println(resp.Proto)
	// fmt.Println(resp.Header)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return false
	}

	// fmt.Println(string(body))
	fmt.Println("===================eee")
	defer resp.Body.Close()
	// fmt.Println(strings.Split(prox.Host,":")[1])
	if strings.Contains(string(body), strings.Split(prox.Host, ":")[0]) {
		fmt.Println("包含", prox.Host)
		return true
	}
	// 删除无效代理
	if Conf.UseProxyPool {
		_, err := http.Get(Conf.PPIP + ":" + Conf.PPPort + "/delete/?proxy=" + prox.Host)
		if err != nil {
			log.Println(err)
		}
		// body, err := ioutil.ReadAll(resp.Body)
		// if err != nil {
		// 	fmt.Println(err)
		// 	return false
		// }
		log.Println("JCTLog: 删除代理: ", prox.Host)
	}
	// if (resp.StatusCode != 200) {
	err = fmt.Errorf("Connect server using proxy error,StatusCode [%d]", resp.StatusCode)
	return false
	// }

}
