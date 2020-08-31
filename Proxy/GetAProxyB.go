package Proxy
import (
	"time"
	"math/rand"
	"os"
	"strings"
	"JCRandomProxy/Conf"
	"log"
	"bufio"
)
func GetAProxyB() (string,string,error) {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())
	// log.Printf("JCTLog:%s",Conf.CustomProxyFile)
    file, err := os.Open(Conf.CustomProxyFile)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
	var proxyList []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		proxystr := scanner.Text()
        proxyList = append(proxyList,proxystr)
	}
	proxystr := proxyList[rand.Intn(len(proxyList))]
	ptype := strings.Split(proxystr,",")[0]
	proxy := strings.Split(proxystr,",")[1]
    if err := scanner.Err(); err != nil {
		log.Fatal(err)
    }
	return proxy,ptype,err
}