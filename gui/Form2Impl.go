package gui

import (
	"crypto/md5"
	"fmt"
	"strings"

	"../Proxy"
	"github.com/ying32/govcl/vcl"
)

//::private::
type TForm2Fields struct {
}

func (f *TForm2) OnButton1Click(sender vcl.IObject) {
	protocol := f.ComboBox1.Text()
	ipPort := f.Edit1.Text()
	ip := strings.Split(ipPort, ":")[0]
	port := strings.Split(ipPort, ":")[1]
	// Form1.ListView1.Items().BeginUpdate()
	item := Form1.ListView1.Items().Add()
	item.SetCaption(fmt.Sprintf("%d", Form1.ListView1.Items().Count()))
	item.SubItems().Add(fmt.Sprintf("%s", protocol))
	item.SubItems().Add(fmt.Sprintf("%s", ip))
	item.SubItems().Add(fmt.Sprintf("%s", port))
	item.MakeVisible(true)
	f.Close()
	// Proxy.MetaProxymap[fmt.Sprintf("%x", md5.Sum([]byte(protocol+"://"+ip+":"+port)))] = Proxy.Aproxy{protocol, ip, port}
	Proxy.MSafeMetaProxymap.WriteAproxy(fmt.Sprintf("%x", md5.Sum([]byte(protocol+"://"+ip+":"+port))), Proxy.Aproxy{protocol, ip, port})
	// Form1.ListView1.Items().EndUpdate()
}

func (f *TForm2) OnButton2Click(sender vcl.IObject) {
	f.Close()
}
