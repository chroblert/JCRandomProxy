package Conf
import (
    "path/filepath"
    "fmt"
    "github.com/go-ini/ini"
    "crypto/tls"
)
var (
	PPIP string
	PPPort string
	UseProxyPool bool
    CustomProxyFile string
	Port string
	UseProxy bool
 )

 func InitConfig() {
     confFile,_ := filepath.Abs("Conf/config.ini")
     cfg,err  := ini.Load(confFile)
     if err != nil {
         panic(err)
     }
     fmt.Println("JCTest",cfg)
     UseProxyPool,_ = cfg.Section("main").Key("UseProxypool").Bool()
	 Port = cfg.Section("main").Key("Port").String()
	 UseProxy,_ = cfg.Section("main").Key("UseProxy").Bool()
     PPIP  = cfg.Section("proxypool").Key("PPIP").String()
	 PPPort = cfg.Section("proxypool").Key("PPPort").String()
	 CustomProxyFile,_ = filepath.Abs(cfg.Section("customproxy").Key("CustomProxyFile").String())
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
