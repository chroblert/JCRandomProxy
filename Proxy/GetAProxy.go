package Proxy

import (
	"JCRandomProxy/Conf"
	"math/rand"
)

// 定义接收返回的代理的结构体
type (
	PPCountItem struct {
		Https int64 `json:"https"`
		Total int64 `json:"total"`
	}
	PPCount struct {
		Count PPCountItem `json:"count"`
	}
	PP struct {
		Anonymous   string `json:"anonymous"`
		Check_count int64  `json:"check_count"`
		Fail_count  int64  `json:"Fail_count"`
		Https       bool   `json:"https"`
		Last_status bool   `json:"Last_status"`
		Last_time   string `json:"last_time"`
		Proxy       string `json:"proxy"`
		Region      string `json:"region"`
		Source      string `json:"source"`
		Type        string `json:"type"`
	}
)

type aproxy struct {
	protocol string
	ip       string
	port     string
}

var minProxyNum int = 10
var maxProxyNum int = 30

// var proxylist []aproxy
// 经过验证的可用代理池
var proxymap = make(map[string]aproxy)

// 从文件中读取的代理
var metaproxymap = make(map[string]aproxy)

func GetAProxy() (string, string, error) {
	// 先判断可用代理池中的可用代理数量是否大于等于10
	// 若大于等于10，则从可用代理池中随机取出一个
	if len(proxymap) >= 10 {
		// tmpProxy := proxylist[rand.Intn(len(proxylist))]
		tmpProxy := GetAvailableProxy(proxymap)
		return tmpProxy.protocol, tmpProxy.ip + ":" + tmpProxy.port, nil
	}
	if Conf.UseProxyPool {
		return GetAProxyA()
	} else {
		return GetAProxyB()
	}

}

// 随机的从可用代理池中取出一个代理
func GetAvailableProxy(tmp map[string]aproxy) aproxy {
	keys := make([]string, 0, len(tmp))
	for k := range tmp {
		keys = append(keys, k)
	}
	return tmp[keys[rand.Intn(len(keys))]]
}
