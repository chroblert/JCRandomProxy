package Proxy

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"../Conf"
)

func CheckProxyC(proxyAddr, checkaddr string) bool {
	httpproxy := proxyAddr
	prox, err := url.Parse(proxyAddr)
	if err != nil {
		return false
	}
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse(httpproxy)
	}
	ht := &http.Transport{
		Proxy: proxy,
	}
	cli := &http.Client{
		Transport: ht,
		Timeout:   time.Duration(Conf.Timeout) * time.Second,
	}
	// log.Println(time.Now().Format("15:04:05"))
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
		log.Printf("代理%s有效\n", proxyAddr)
		return true
	}
	return false
}
