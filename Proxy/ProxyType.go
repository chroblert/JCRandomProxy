package Proxy

import "sync"

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
