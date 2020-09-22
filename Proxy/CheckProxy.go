package Proxy

import (
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

	"../Conf"
)

// 验证代理服务器是否可用
func CheckProxy(proxyAddr, checkaddr string) bool {
	// log.Println(proxyAddr)
	proxymd5 := fmt.Sprintf("%x", md5.Sum([]byte(proxyAddr)))
	// 代理服务器可用，则添加至map中
	if CheckProxyB(proxyAddr, checkaddr) {
		protocol := strings.Split(proxyAddr, ":")[0]
		ip := strings.Split(strings.Split(proxyAddr, "/")[2], ":")[0]
		port := strings.Split(proxyAddr, ":")[2]
		tmpproxy := aproxy{protocol, ip, port}
		// proxylist = append(proxylist, tmpproxy)
		// Proxymap[proxymd5] = tmpproxy
		MSafeProxymap.WriteAproxy(proxymd5, tmpproxy)
		log.Println("当前可用代理池: ", MSafeProxymap.Map)
		return true
	}
	// 代理服务器不可用，则删除
	// delete(Proxymap, proxymd5)
	MSafeProxymap.DeleteAproxy(proxymd5)
	// delete(MetaProxymap, proxymd5)
	MSafeMetaProxymap.DeleteAproxy(proxymd5)
	return false
}
func CheckProxyB(proxyAddr, checkaddr string) bool {
	// 20200922使用新方法校验代理
	httpproxy := proxyAddr
	prox, _ := url.Parse(proxyAddr)
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse(httpproxy)
	}
	ht := &http.Transport{
		Proxy: proxy,
	}
	cli := &http.Client{
		Transport: ht,
	}
	// req,err := http.NewRequest("GET","https://myip.ipip.net",nil)
	log.Println(time.Now().Format("15:04:05"))
	resp, err := cli.Get(checkaddr)
	if err != nil {
		log.Println(err)
		return false
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return false
	}
	if strings.Contains(string(body), strings.Split(prox.Host, ":")[0]) {
		log.Println("包含", prox.Host)
		log.Println("代理xxxxx: ", string(body))
		return true
	}
	//结束
	return false
}
func CheckProxyA(proxyAddr, checkaddr string) bool {

	prox, _ := url.Parse(proxyAddr)
	log.Println("JCTLog: 代理地址: ", prox.Host)
	// Dial and create client connection
	// proxc, err := net.DialTimeout("tcp", prox.Host, time.Second*5)
	proxc, err := net.Dial("tcp", prox.Host)

	if err != nil {

		return false
	}
	defer proxc.Close()
	// 解析最终目标url
	reqURL, err := url.Parse(checkaddr)
	if err != nil {
		return false
	}
	log.Println("JCTLog: 校验地址: ", reqURL.String())
	req, err := http.NewRequest(http.MethodGet, reqURL.String(), nil)
	if err != nil {
		log.Println("JCTLog: http.NewRequest: ", err)
		return false
	}
	req.Close = false
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.3")
	err = req.Write(proxc)
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
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return false
	}
	if strings.Contains(string(body), strings.Split(prox.Host, ":")[0]) {
		log.Println("包含", prox.Host)
		log.Println("代理xxxxx: ", string(body))
		log.Println("请求内容", req.Host, req.URL, req.RequestURI, req.RemoteAddr, req.Header, req.Body)
		return true
	}
	// 删除无效代理
	if Conf.UseProxyPool {
		_, err := http.Get(Conf.PPIP + ":" + Conf.PPPort + "/delete/?proxy=" + prox.Host)
		if err != nil {
			log.Println(err)
		}
		log.Println("JCTLog: 删除代理: ", prox.Host)
	}
	err = fmt.Errorf("Connect server using proxy error,StatusCode [%d]", resp.StatusCode)
	return false

}
