package Proxy

func ProxyCheckSche(stop chan int) {
	MSafeProxymap.ProxyCheck(stop)
}
