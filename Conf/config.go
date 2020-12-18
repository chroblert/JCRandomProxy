package Conf

import (
	"crypto/tls"
	"path/filepath"

	"github.com/go-ini/ini"
)

var (
	// main设置
	UseProxyPool    bool = true
	CustomProxyFile string
	SaveProxyFile   string = "proxy.lst"
	Port            string = "8081"
	UseProxy        bool
	UseHttpsProxy   bool
	// Proxypool代理池设置
	PPIP   string = "http://localhost"
	PPPort string = "5010"
	// 校验代理设置
	MinProxyNum    int    = 0
	MaxProxyNum    int    = 5
	ProxyCheckAddr string = "https://myip.ipip.net"
	Timeout        int    = 5
	StopUrl        string = "http://myip.ipip.net"
	CheckInterval  int    = 2
	EnableCheck    bool   = true
	CustomSucFlag  string = "false"
	// 日志设置
	LogPath  string = "logss/app.log"
	LogCount int    = 5
	MaxSize  int64  = 1024 * 1024 * 256
	MaxAge   int    = 3
	// 认证设置
	ProxyUser   string = "admin"
	ProxyPasswd string = "admin"
)

func InitConfig(aTimeout, aMinProxyNum, aMaxProxyNum int, aUseProxyPool bool, aPort string, aUseProxy bool, aUseHttpsProxy bool, aPPIP string, aPPPort string) {
	Timeout = aTimeout
	MinProxyNum = aMinProxyNum
	MaxProxyNum = aMaxProxyNum
	UseProxyPool = aUseProxyPool
	Port = aPort
	UseProxy = aUseProxy
	UseHttpsProxy = aUseHttpsProxy
	PPIP = aPPIP
	PPPort = aPPPort

}
func InitConfigFromFile() {
	confFile, _ := filepath.Abs("Conf/config.ini")
	cfg, err := ini.Load(confFile)
	if err != nil {
		panic(err)
	}
	// main设置
	UseProxyPool, _ = cfg.Section("main").Key("UseProxypool").Bool()
	Port = cfg.Section("main").Key("Port").String()
	UseProxy, _ = cfg.Section("main").Key("UseProxy").Bool()
	UseHttpsProxy, _ = cfg.Section("main").Key("UseHttpsProxy").Bool()
	ProxyUser = cfg.Section("main").Key("ProxyUser").String()
	ProxyPasswd = cfg.Section("main").Key("ProxyPasswd").String()
	// proxypool设置
	PPIP = cfg.Section("proxypool").Key("PPIP").String()
	PPPort = cfg.Section("proxypool").Key("PPPort").String()
	CustomProxyFile, _ = filepath.Abs(cfg.Section("customproxy").Key("CustomProxyFile").String())
	// checkproxy设置
	ProxyCheckAddr = cfg.Section("checkproxy").Key("ProxyCheckAddr").String()
	StopUrl = cfg.Section("checkproxy").Key("StopUrl").String()
	MinProxyNum, _ = cfg.Section("checkproxy").Key("MinProxyNum").Int()
	MaxProxyNum, _ = cfg.Section("checkproxy").Key("MaxProxyNum").Int()
	Timeout, _ = cfg.Section("checkproxy").Key("Timeout").Int()
	CheckInterval, _ = cfg.Section("checkproxy").Key("CheckInterval").Int()
	EnableCheck, _ = cfg.Section("checkproxy").Key("EnableCheck").Bool()
	CustomSucFlag = cfg.Section("checkproxy").Key("CustomSucFlag").String()

	// log设置
	LogPath = cfg.Section("log").Key("LogPath").String()
	LogCount, _ = cfg.Section("log").Key("LogCount").Int()
	MaxSize, _ = cfg.Section("log").Key("MaxSize").Int64()
	MaxAge, _ = cfg.Section("log").Key("MaxAge").Int()

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
