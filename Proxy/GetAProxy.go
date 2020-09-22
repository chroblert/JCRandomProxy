package Proxy

import (
	"log"
	"math/rand"
	"time"

	"../Conf"
)

// 定义接收返回的代理的结构体
type (
	PPCount struct {
		Count int64 `json:"count"`
	}
	PP struct {
		Check_count int64  `json:"check_count"`
		Fail_count  int64  `json:"Fail_count"`
		Last_status int64  `json:"Last_status"`
		Last_time   string `json:"last_time"`
		Proxy       string `json:"proxy"`
		Region      string `json:"region"`
		Source      string `json:"source"`
		Type        string `json:"type"`
	}
)

type Aproxy struct {
	Protocol string
	Ip       string
	Port     string
}
type aproxy = Aproxy

// var proxylist []aproxy
// 经过验证的可用代理池
var Proxymap = make(map[string]aproxy)

// 使用具有读写锁的map
// var MSafeProxymap = NewSafeProxymap()

// 从文件中读取的代理
var MetaProxymap = make(map[string]aproxy)

// var MSafeMetaProxymap = NewSafeMetaProxymap()

func GetAProxy() (string, string, error) {
	// 先判断可用代理池中的可用代理数量是否大于等于10
	// 若大于等于10，则从可用代理池中随机取出一个
	if len(Proxymap) >= Conf.MinProxyNum {
		// tmpProxy := proxylist[rand.Intn(len(proxylist))]
		tmpProxy := GetAvailableProxy(Proxymap)
		log.Println(tmpProxy)
		return tmpProxy.Ip + ":" + tmpProxy.Port, tmpProxy.Protocol, nil
	}
	if Conf.UseProxyPool {
		return GetAProxyA()
	} else {
		return GetAProxyB()
	}

}

// 随机的从可用代理池中取出一个代理
func GetAvailableProxy(tmp map[string]aproxy) aproxy {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())
	keys := make([]string, 0, len(tmp))
	for k := range tmp {
		keys = append(keys, k)
	}
	// log.Println("keys: ", keys)
	return tmp[keys[rand.Intn(len(keys))]]
}
