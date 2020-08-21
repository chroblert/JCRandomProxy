package Conf
import (
    "path/filepath"
    "fmt"
    "github.com/go-ini/ini"
)
var (
	PPIP string
	PPPort string
	UseProxyPool bool
	CustomProxyFile string
 )

 func InitConfig() {
     confFile,_ := filepath.Abs("Conf/config.ini")
     cfg,err  := ini.Load(confFile)
     if err != nil {
         panic(err)
     }
     fmt.Println("JCTest",cfg)
     UseProxyPool,_ = cfg.Section("main").Key("UseProxypool").Bool()
     PPIP  = cfg.Section("proxypool").Key("PPIP").String()
	 PPPort = cfg.Section("proxypool").Key("PPPort").String()
	 CustomProxyFile,_ = filepath.Abs(cfg.Section("customproxy").Key("CustomProxyFile").String())
 }
