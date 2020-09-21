package Conf

import (
	"crypto/tls"
	"log"
	"path/filepath"

	"github.com/go-ini/ini"
)

var (
	PPIP            string
	PPPort          string
	UseProxyPool    bool
	CustomProxyFile string
	SaveProxyFile   string = "proxy.lst"
	Port            string
	UseProxy        bool
	UseHttpsProxy   bool
	MinProxyNum     int
)

func InitConfig(aMinProxyNum int, aUseProxyPool bool, aPort string, aUseProxy bool, aUseHttpsProxy bool, aPPIP string, aPPPort string) {

	// confFile, _ := filepath.Abs("Conf/config.ini")
	// cfg, err := ini.Load(confFile)
	// if err != nil {
	// 	panic(err)
	// }
	// log.Println("JCTest", cfg)
	MinProxyNum = aMinProxyNum
	UseProxyPool = aUseProxyPool
	Port = aPort
	UseProxy = aUseProxy
	UseHttpsProxy = aUseHttpsProxy
	PPIP = aPPIP
	PPPort = aPPPort
	// log.Println(MinProxyNum)
	// CustomProxyFile, _ = filepath.Abs(cfg.Section("customproxy").Key("CustomProxyFile").String())
	// log.Println(UseHttpsProxy)

}
func InitConfigFromFile() {

	confFile, _ := filepath.Abs("Conf/config.ini")
	cfg, err := ini.Load(confFile)
	if err != nil {
		panic(err)
	}
	log.Println("JCTest", cfg)
	UseProxyPool, _ = cfg.Section("main").Key("UseProxypool").Bool()
	Port = cfg.Section("main").Key("Port").String()
	UseProxy, _ = cfg.Section("main").Key("UseProxy").Bool()
	UseHttpsProxy, _ = cfg.Section("main").Key("UseHttpsProxy").Bool()
	PPIP = cfg.Section("proxypool").Key("PPIP").String()
	PPPort = cfg.Section("proxypool").Key("PPPort").String()
	CustomProxyFile, _ = filepath.Abs(cfg.Section("customproxy").Key("CustomProxyFile").String())
	// log.Println(UseHttpsProxy)

}

type Cfg struct {
	Port    *string
	Raddr   *string
	Log     *string
	Monitor *bool
	Tls     *bool
}

type TlsConfig struct {
	PrivateKeyFile  string
	CertFile        string
	Organization    string
	CommonName      string
	ServerTLSConfig *tls.Config
}

func NewTlsConfig(pk, cert, org, cn string) *TlsConfig {
	return &TlsConfig{
		PrivateKeyFile: pk,
		CertFile:       cert,
		Organization:   org,
		CommonName:     cn,
		ServerTLSConfig: &tls.Config{
			CipherSuites: []uint16{
				tls.TLS_RSA_WITH_RC4_128_SHA,
				tls.TLS_RSA_WITH_3DES_EDE_CBC_SHA,
				tls.TLS_RSA_WITH_AES_128_CBC_SHA,
				tls.TLS_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_RSA_WITH_AES_128_CBC_SHA256,
				tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_ECDSA_WITH_RC4_128_SHA,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
				tls.TLS_ECDHE_RSA_WITH_RC4_128_SHA,
				tls.TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA,
				tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
				tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
				tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
				tls.TLS_FALLBACK_SCSV,
			},
			PreferServerCipherSuites: true,
		},
	}
}
