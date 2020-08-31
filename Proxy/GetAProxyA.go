package Proxy
import (
	"log"
	"JCRandomProxy-v1.0/Conf"
	"net/http"
	"io/ioutil"
	"encoding/json"
)
func GetAProxyA() (string,string,error){
	// fmt.Println(Conf.PPIP)
	ppCountUrl := Conf.PPIP + ":" + Conf.PPPort + "/get_status/"
	ppGetUrl := Conf.PPIP + ":" + Conf.PPPort + "/get/"
	// 查看当前有多少代理
	req, err := http.NewRequest("GET",ppCountUrl,nil)
	if err != nil {
		log.Println(err)
		return "","",err
	}
	req.Header.Add("accept","application/json")
	res,_ := http.DefaultClient.Do(req)
	resbody,_ := ioutil.ReadAll(res.Body)
	// 解析json数据
	ppCount := &PPCount{}
	err = json.Unmarshal([]byte(resbody),ppCount)
	if err != nil {
		log.Println("error: ",err)
		return "","",err
	}
	// 判断是否有可用的代理
	if ppCount.Count < 1 {
		log.Println("当前没有可用代理")
		return "","",err
	}
	// 获取一个代理
	proxy := &PP{}
	req, _ = http.NewRequest("GET", ppGetUrl, nil)
    req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	defer res.Body.Close()
	res, _ = http.DefaultClient.Do(req)
	resbody, _ = ioutil.ReadAll(res.Body)
	err = json.Unmarshal([]byte(resbody), proxy)
	if err != nil {
		log.Println("error:", err)
		return "","",err
    }
	// 判断代理的类型
	var ptype string
	if proxy.Type != "" {
		ptype = proxy.Type
	} else{
		ptype = "http"
	}
	// 返回
	return proxy.Proxy,ptype,err
}