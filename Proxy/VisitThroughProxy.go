package Proxy

import (
	"net/http"
	"net/url"
	"time"

	log "github.com/chroblert/JCRandomProxy/v3/Logs"

	"github.com/chroblert/JCRandomProxy/v3/Conf"
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
		Timeout:   time.Duration(Conf.Timeout) * time.Second, //20200922: 增加超时机制
	}

	// req,err := http.NewRequest("GET","https://myip.ipip.net",nil)
	_, err := cli.Get(targeturl)
	if err != nil {
		log.Println(err)
		return
	}
}
