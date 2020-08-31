package Proxy
import (
	"JCRandomProxy-v1.0/Conf"
)

// 定义接收返回的代理的结构体
type (
	PPCount struct {
		Count int64 `json:"count"`
	}
	PP struct {
		Check_count int64 `json:"check_count"`
		Fail_count int64 `json:"Fail_count"`
		Last_status int64 `json:"Last_status"`
		Last_time string `json:"last_time"`
		Proxy string `json:"proxy"`
		Region string `json:"region"`
		Source string `json:"source"`
		Type string `json:"type"`
	}
)

func GetAProxy()  (string,string,error) {
	if Conf.UseProxyPool {
		return GetAProxyA()
	}else{
		return GetAProxyB()
	}
	
}