package Proxy

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"../Conf"
)

// 从ProxyPool中获取代理
func GetProxysA() {
	// 获取proxypool代理池的状态
	// ppStatusUrl := Conf.PPIP + ":" + Conf.PPPort + "/get_status/"
	for i := MSafeProxymap.Length(); i < Conf.MaxProxyNum; i = MSafeProxymap.Length() {
		tmpAproxy, err := GetAproxyA()
		if err != nil {
			log.Println("从proxypool代理池中获取代理失败: ", err)
			continue
		}
		tmpproxyaddr := tmpAproxy.Protocol + "://" + tmpAproxy.Ip + ":" + tmpAproxy.Port
		tmpproxyaddrmd5 := fmt.Sprintf("%x", md5.Sum([]byte(tmpproxyaddr)))
		if CheckProxyC(tmpproxyaddr, Conf.ProxyCheckAddr) {
			MSafeProxymap.WriteAproxy(tmpproxyaddrmd5, tmpAproxy)
		} else {
			// 删除无效代理
			DeleteProxyA(tmpproxyaddr)
		}

	}

}

// 从proxypool代理池中删去某个代理
func DeleteProxyA(proxyaddr string) {
	_, err := http.Get(Conf.PPIP + ":" + Conf.PPPort + "/delete/?proxy=" + proxyaddr)
	if err != nil {
		log.Println(err)
	}
	log.Printf("删除代理 %s", proxyaddr)
}

func GetAproxyA() (Aproxy, error) {
	// 从proxypool代理池中获取一个代理
	ppGetproxyUrl := Conf.PPIP + ":" + Conf.PPPort + "/get/"
	req, err := http.NewRequest("GET", ppGetproxyUrl, nil)
	if err != nil {
		log.Println(err)
		return Aproxy{}, err
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return Aproxy{}, err
	}
	resbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return Aproxy{}, err
	}
	defer res.Body.Close()
	ppproxy := &PP{}
	err = json.Unmarshal(resbody, ppproxy)
	if err != nil {
		log.Println(err)
		return Aproxy{}, err
	}
	var protocol string
	if ppproxy.Type != "" {
		protocol = ppproxy.Type
	} else {
		protocol = "http"
	}
	ip := strings.Split(ppproxy.Proxy, ":")[0]
	port := strings.Split(ppproxy.Proxy, ":")[1]
	return Aproxy{protocol, ip, port}, nil
}
