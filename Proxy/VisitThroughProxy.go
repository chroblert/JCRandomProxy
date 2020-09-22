package Proxy

import (
	"log"
	"net/http"
	"net/url"
	"time"
)

func VisitThroughProxy(httpproxy, targeturl string) {
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse(httpproxy)
	}
	ht := &http.Transport{
		Proxy: proxy,
	}
	cli := &http.Client{
		Transport: ht,
		Timeout:   5 * time.Second, //20200922: 增加超时机制
	}

	// req,err := http.NewRequest("GET","https://myip.ipip.net",nil)
	log.Println(time.Now().Format("15:04:05"))
	_, err := cli.Get(targeturl)
	if err != nil {
		log.Println(err)
		return
	}
}
