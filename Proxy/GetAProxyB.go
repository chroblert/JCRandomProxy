package Proxy

import (
	"bufio"
	"crypto/md5"
	"fmt"

	log "../Logs"

	"os"
	"strings"

	"../Conf"
)

// 从文件中读取代理到MetaSafeProxymap中【即元代理池】
func GetMetaproxyBFromFile() error {
	file, err := os.Open(Conf.CustomProxyFile)
	if err != nil {
		log.Println(err)
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		proxystr := scanner.Text()
		if !strings.Contains(proxystr, ",") {
			return fmt.Errorf("格式错误")
		}
		protocol := strings.Split(proxystr, ",")[0]
		proxy := strings.Split(proxystr, ",")[1]
		if !strings.Contains(proxystr, ":") {
			return fmt.Errorf("格式错误")
		}
		ip := strings.Split(proxy, ":")[0]
		port := strings.Split(proxy, ":")[1]
		tmpmd5 := fmt.Sprintf("%x", md5.Sum([]byte(protocol+"://"+ip+":"+port)))
		tmpaproxy := Aproxy{protocol, ip, port, 4}
		MSafeMetaProxymap.WriteAproxy(tmpmd5, tmpaproxy)
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// 从文件中获取代理
func GetProxysB(stop chan int) {
	for i := MSafeProxymap.Length(); i < Conf.MaxProxyNum; i = MSafeProxymap.Length() {
		// 结束条件：MSafeMetaProxymap中代理数量为0或MSafeProxymap中代理数量达标或者收到停止信号
		if MSafeMetaProxymap.Length() == 0 {
			log.Println("当前元代理池中的数量为空，停止获取可用代理")
			return
		}
		// 收到停止信号
		select {
		case <-stop:
			log.Println("收到停止信号：停止获取可用代理")
			// stop <- 1
			return
		default:
		}
		tmpAproxy, err := GetAproxyB()
		if err != nil {
			log.Println("从元代理池中获取代理失败: ", err)
			continue
		}
		tmpproxyaddr := tmpAproxy.Protocol + "://" + tmpAproxy.Ip + ":" + tmpAproxy.Port
		tmpproxyaddrmd5 := fmt.Sprintf("%x", md5.Sum([]byte(tmpproxyaddr)))
		if !Conf.EnableCheck || CheckProxyC(tmpproxyaddr, Conf.ProxyCheckAddr) {
			MSafeProxymap.WriteAproxy(tmpproxyaddrmd5, tmpAproxy)
		} else {
			// 删除无效代理
			DeleteProxyB(tmpproxyaddr)
		}

	}
}

// 从元代理池中删去某个代理
func DeleteProxyB(proxyaddr string) {
	tmpmd5 := fmt.Sprintf("%x", md5.Sum([]byte(proxyaddr)))
	MSafeMetaProxymap.DeleteAproxy(tmpmd5)
	log.Printf("删除代理 %s", proxyaddr)
}

// 从元代理池中随机获取代理
func GetAproxyB() (Aproxy, error) {
	//
	tmpAproxy, ok := MSafeMetaProxymap.GetARandProxy()
	if ok {
		return tmpAproxy, nil
	}
	return Aproxy{}, fmt.Errorf("%s", "error")
}
