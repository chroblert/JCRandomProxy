package gui

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"../Conf"
	"../Proxy"
	"../ProxyEntry"
	"github.com/hpcloud/tail"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

//::private::
type TForm1Fields struct {
}

// 启动代理
func (f *TForm1) OnButton1Click(sender vcl.IObject) {
	var UseProxyPool bool = true
	if !Form1.RadioButton1.Checked() {
		UseProxyPool = false
	}
	var PPIP string = Form1.Edit1.Text()
	var PPPort string = Form1.Edit2.Text()
	var Port string = Form1.Edit3.Text()
	// 临时
	PPIP = "http://10.103.91.179"
	PPPort = "5010"
	Port = "8081"
	var UseProxy bool = true
	var UseHttpsProxy bool = true
	var MinProxyNum, _ = strconv.Atoi(Form1.Edit4.Text())
	Conf.InitConfig(MinProxyNum, UseProxyPool, Port, UseProxy, UseHttpsProxy, PPIP, PPPort)
	// 启动代理
	go ProxyEntry.Proxymain()
	// 渲染可用代理池
	go RenderValidProxyPool()

	// 启动日志实时输出
	//go logRealTime()//拉低性能，暂时取消
}

// 停止代理
func (f *TForm1) OnButton2Click(sender vcl.IObject) {

}

func (f *TForm1) OnButton3Click(sender vcl.IObject) {
	dlgOpen := vcl.NewOpenDialog(f)
	dlgOpen.SetFilter("文本文件(*.txt)|*.txt|所有文件(*.*)|*.*")
	dlgOpen.SetOptions(dlgOpen.Options().Include(types.OfShowHelp, types.OfAllowMultiSelect))
	dlgOpen.SetTitle("打开")
	// 打开文件成功后
	if dlgOpen.Execute() {
		f.ListBox2.Items().Add(time.Now().Format(fmt.Sprintf("2006-01-02 15:04:05 : %s", "导入代理文件")))
		log.Println("导入文件")
		log.Println(dlgOpen.FileName())
		Conf.CustomProxyFile = dlgOpen.FileName()
		var tmpmap = make(map[string]Proxy.Aproxy)
		tmpmap = Proxy.GetMetaproxyFromFile()
		f.ListView1.Items().BeginUpdate()
		i := 0
		for k := range tmpmap {
			i++
			item := f.ListView1.Items().Add()
			item.SetCaption(fmt.Sprintf("%d", i))
			item.SubItems().Add(tmpmap[k].Protocol)
			item.SubItems().Add(tmpmap[k].Ip)
			item.SubItems().Add(tmpmap[k].Port)
		}
		f.ListView1.Items().EndUpdate()

	}

}

// 导出
func (f *TForm1) OnButton4Click(sender vcl.IObject) {
}

// 添加代理
func (f *TForm1) OnButton5Click(sender vcl.IObject) {
	Form2.ShowModal()
}

// 删除代理
func (f *TForm1) OnButton6Click(sender vcl.IObject) {

}

// 日志实时输出
func logRealTime() {
	t, _ := tail.TailFile("log.txt", tail.Config{Follow: true})
	for line := range t.Lines {
		// fmt.Println(line.Text)
		if Form1.ListBox2.Items().Count() > 30 {
			Form1.ListBox2.Items().Delete(0)
		}
		Form1.ListBox2.Items().Add(line.Text)
	}
}

// 定时渲染可用代理池
func RenderValidProxyPool() {
	ticker := time.NewTicker(time.Duration(2 * time.Second))
	for range ticker.C {
		Form1.ListBox1.Items().Clear()
		for k := range Proxy.Proxymap {
			tmp := Proxy.Proxymap[k]
			Form1.ListBox1.Items().Add(tmp.Protocol + "://" + tmp.Ip + ":" + tmp.Port)
		}
	}

}
