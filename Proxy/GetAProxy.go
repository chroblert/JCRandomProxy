package Proxy

import (
	"github.com/chroblert/JCRandomProxy/Conf"
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
	Protocol  string
	Ip        string
	Port      string
	FailLimit int
}
type aproxy = Aproxy

// 使用具有读写锁的map
var MSafeProxymap = NewSafeProxymap()

// 从文件中读取的元代理池
var MSafeMetaProxymap = NewSafeMetaProxymap()

func GetProxys(stop chan int) {
	if Conf.UseProxyPool {
		GetProxysA()
	} else {
		GetProxysB(stop)
	}
}
