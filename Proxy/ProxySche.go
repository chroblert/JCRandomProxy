package Proxy

func ProxyCheckSche(stop chan int) {
	MSafeProxymap.ProxyCheck(stop)
}

func ProxyNumCheckSche(stop chan int) {
	MSafeProxymap.GetProxysSche(stop)
}
