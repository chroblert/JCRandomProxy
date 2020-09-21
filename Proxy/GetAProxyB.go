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
	MetaProxymap = GetMetaproxyFromFile()
	// return proxy, ptype, err
	tmp := GetAvailableProxy(MetaProxymap)
	return tmp.Ip + ":" + tmp.Port, tmp.Protocol, nil
}
func GetMetaproxyFromFile() map[string]aproxy {
	var tmpmetaproxymap = make(map[string]aproxy)
	// 设置随机数种子
	// rand.Seed(time.Now().UnixNano())
	file, err := os.Open(Conf.CustomProxyFile)
	if err != nil {
		log.Println(file)
		log.Fatal(err)
	}
	defer file.Close()
	// var proxyList []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		proxystr := scanner.Text()
		ptype := strings.Split(proxystr, ",")[0]
		proxy := strings.Split(proxystr, ",")[1]
		IP := strings.Split(proxy, ":")[0]
		Port := strings.Split(proxy, ":")[1]
		tmpmetaproxymap[fmt.Sprintf("%x", md5.Sum([]byte(ptype+"://"+IP+":"+Port)))] = aproxy{ptype, IP, Port}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return tmpmetaproxymap
}
