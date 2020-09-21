package Proxy

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"log"

	"../Conf"

	"os"
	"strings"
)

func GetAProxyB() (string, string, error) {
	log.Println(MetaProxymap)
	if len(MetaProxymap) != 0 {
		tmp := GetAvailableProxy(MetaProxymap)
		delete(MetaProxymap, fmt.Sprintf("%x", md5.Sum([]byte(tmp.Protocol+"://"+tmp.Ip+":"+tmp.Port))))
		return tmp.Ip + ":" + tmp.Port, tmp.Protocol, nil
	}
	MetaProxymap, err := GetMetaproxyFromFile()
	if err != nil {
		return "", "", err
	}
	// return proxy, ptype, err
	tmp := GetAvailableProxy(MetaProxymap)
	return tmp.Ip + ":" + tmp.Port, tmp.Protocol, nil
}
func GetMetaproxyFromFile() (map[string]aproxy, error) {
	var tmpmetaproxymap = make(map[string]aproxy)
	// 设置随机数种子
	// rand.Seed(time.Now().UnixNano())
	file, err := os.Open(Conf.CustomProxyFile)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer file.Close()
	// var proxyList []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		proxystr := scanner.Text()
		if !strings.Contains(proxystr, ",") {
			return nil, fmt.Errorf("格式错误")
		}
		ptype := strings.Split(proxystr, ",")[0]
		proxy := strings.Split(proxystr, ",")[1]
		if !strings.Contains(proxystr, ":") {
			return nil, fmt.Errorf("格式错误")
		}
		IP := strings.Split(proxy, ":")[0]
		Port := strings.Split(proxy, ":")[1]
		tmpmetaproxymap[fmt.Sprintf("%x", md5.Sum([]byte(ptype+"://"+IP+":"+Port)))] = aproxy{ptype, IP, Port}
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
		return nil, err
	}
	if len(tmpmetaproxymap) < 1 {
		return nil, fmt.Errorf("空，未读取到代理")
	}
	return tmpmetaproxymap, nil
}
