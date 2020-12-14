package Proxy

import (
	"math/rand"
	"sort"
	"sync"
	"time"

	log "../Logs"

	"../Conf"
)

type SafeProxymap struct {
	sync.RWMutex
	Map map[string]Aproxy
}
type SafeMetaProxymap struct {
	sync.RWMutex
	Map map[string]Aproxy
}

func NewSafeProxymap() *SafeProxymap {
	var spm = new(SafeProxymap)
	spm.Map = make(map[string]Aproxy)
	return spm
}

func (spm *SafeProxymap) ReadAproxy(k string) Aproxy {
	spm.RLock()
	value := spm.Map[k]
	spm.RUnlock()
	return value
}

func (spm *SafeProxymap) WriteAproxy(k string, v Aproxy) {
	spm.Lock()
	spm.Map[k] = v
	spm.Unlock()
}
func (spm *SafeProxymap) Length() int {
	spm.RLock()
	value := len(spm.Map)
	spm.RUnlock()
	return value
}
func (spm *SafeProxymap) DeleteAproxy(k string) {
	spm.Lock()
	delete(spm.Map, k)
	spm.Unlock()
}

func (spm *SafeProxymap) AProxyExist(k string) bool {
	spm.RLock()
	_, ok := spm.Map[k]
	spm.RUnlock()
	return ok
}

// 从可用代理池中随机获取一个代理
func (spm *SafeProxymap) GetARandProxy() (Aproxy, bool) {
	rand.Seed(time.Now().UnixNano())
	spm.RLock()
	defer spm.RUnlock()
	if tmp := len(spm.Map); tmp > 0 {
		keys := make([]string, 0, tmp)
		for k := range spm.Map {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		// log.Println("keys: ", keys)
		return spm.Map[keys[rand.Intn(len(keys))]], true
	}
	return Aproxy{}, false
}

// 校验可用代理池
// 校验间隔：Conf.CheckInterval 以分钟为单位
func (spm *SafeProxymap) ProxyCheck(stop chan int) {
	// 每两分钟校验一次可用代理池
	ticker := time.NewTicker(time.Duration(Conf.CheckInterval) * time.Minute)
	for {
		select {
		case <-stop:
			log.Println("停止校验可用代理池")
			// stop <- 1
			return
		case <-ticker.C:
			spm.RLock()
			var keys []string
			// 遍历SafeProxymap中所有的代理
			for k := range spm.Map {
				keys = append(keys, k)
			}
			spm.RUnlock()
			spm.Lock()
			for _, k := range keys {
				tmpaproxy := spm.Map[k]
				// 20200923：要不要开启一个协程去校验
				protocol := tmpaproxy.Protocol
				ip := tmpaproxy.Ip
				port := tmpaproxy.Port
				proxyadd := protocol + "://" + ip + ":" + port
				res := CheckProxyC(proxyadd, "https://myip.ipip.net")
				if !res {
					//删除代理
					delete(spm.Map, k)
				}
			}
			spm.Unlock()
		}
	}

}

// 20200928: 定时检测可用代理池中代理的数量，并获取代理
func (spm *SafeProxymap) GetProxysSche(stop chan int) {
	ticker := time.NewTicker(time.Duration(Conf.CheckInterval) * time.Minute)
	for {
		select {
		case <-stop:
			log.Println("停止检测并获取可用代理")
			// stop <- 1
			return
		case <-ticker.C:
			if spm.Length() <= Conf.MinProxyNum {
				GetProxys(stop)
			}
		}
	}
}
func NewSafeMetaProxymap() *SafeMetaProxymap {
	var smpm = new(SafeMetaProxymap)
	smpm.Map = make(map[string]Aproxy)
	return smpm
}
func (smpm *SafeMetaProxymap) ReadAproxy(k string) Aproxy {
	smpm.RLock()
	value := smpm.Map[k]
	smpm.RUnlock()
	return value
}

func (smpm *SafeMetaProxymap) WriteAproxy(k string, v Aproxy) {
	smpm.Lock()
	smpm.Map[k] = v
	smpm.Unlock()
}
func (smpm *SafeMetaProxymap) Length() int {
	smpm.RLock()
	value := len(smpm.Map)
	smpm.RUnlock()
	return value
}
func (smpm *SafeMetaProxymap) DeleteAproxy(k string) {
	smpm.Lock()
	delete(smpm.Map, k)
	smpm.Unlock()
}

// 从元代理池中随机获取一个代理
func (smpm *SafeMetaProxymap) GetARandProxy() (Aproxy, bool) {
	rand.Seed(time.Now().UnixNano())
	smpm.RLock()
	defer smpm.RUnlock()
	if tmp := len(smpm.Map); tmp > 0 {
		keys := make([]string, 0, tmp)
		for k := range smpm.Map {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		// log.Println("keys: ", keys)
		return smpm.Map[keys[rand.Intn(len(keys))]], true
	}
	return Aproxy{}, false
}
