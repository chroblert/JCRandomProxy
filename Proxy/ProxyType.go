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
