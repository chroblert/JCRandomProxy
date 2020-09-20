package Proxy

import (
	"JCRandomProxy/Conf"
	"bufio"
	"crypto/md5"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func GetAProxyB() (string, string, error) {
	if len(metaproxymap) != 0 {
		tmp := GetAvailableProxy(metaproxymap)
		return tmp.ip + ":" + tmp.port, tmp.protocol, nil
	}
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())
	// log.Printf("JCTLog:%s",Conf.CustomProxyFile)
	file, err := os.Open(Conf.CustomProxyFile)
	if err != nil {
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
		metaproxymap[fmt.Sprintf("%x", md5.Sum([]byte(ptype+"://"+IP+":"+Port)))] = aproxy{ptype, IP, Port}
		// proxyList = append(proxyList, proxystr)
	}
	// proxystr := proxyList[rand.Intn(len(proxyList))]
	// ptype := strings.Split(proxystr, ",")[0]
	// proxy := strings.Split(proxystr, ",")[1]
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// return proxy, ptype, err
	tmp := GetAvailableProxy(metaproxymap)
	return tmp.ip + ":" + tmp.port, tmp.protocol, nil
}
