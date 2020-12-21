// 由res2go IDE插件自动生成。
package main

import (
	// "io"

	"strconv"

	log "github.com/chroblert/JCRandomProxy/v3/Logs"
	// "os"
	"github.com/chroblert/JCRandomProxy/v3/Conf"
	"github.com/chroblert/JCRandomProxy/v3/gui"
	"github.com/ying32/govcl/vcl"
)

var mw interface{}

func init() {
	// logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	// if err != nil {
	// 	panic(err)
	// }
	// mw = io.MultiWriter(os.Stdout, logFile)
	// log.SetOutput(mw.(io.Writer))
	// log.SetFlags(log.LstdFlags | log.Lshortfile)
	Conf.InitConfigFromFile()
	// fmt.Println(Conf.MaxSize)
	log.InitLogs(Conf.LogPath, Conf.MaxSize, Conf.MaxAge, Conf.LogCount)
}
func main() {
	vcl.Application.SetScaled(true)
	vcl.Application.SetTitle("JCRP【随机代理】")
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.CreateForm(&gui.Form1)
	vcl.Application.CreateForm(&gui.Form2)
	UpdateForm()
	vcl.Application.Run()

}

// 使用配置文件中的值更新窗口
func UpdateForm() {
	// log.Println("xxxx: ", Conf.PPIP)
	gui.Form1.SetCaption("JCRP【随机代理】 v3.3.7 - by JC0o0l")
	gui.Form1.Edit1.SetText(Conf.PPIP)
	gui.Form1.Edit2.SetText(Conf.PPPort)
	gui.Form1.Edit3.SetText(Conf.Port)
	gui.Form1.Edit4.SetText(strconv.Itoa(Conf.MinProxyNum))
	gui.Form1.Edit5.SetText(strconv.Itoa(Conf.MaxProxyNum))
	gui.Form1.Edit6.SetText(strconv.Itoa(Conf.Timeout))
	if Conf.UseProxyPool {
		gui.Form1.RadioButton1.SetChecked(true)
	} else {
		gui.Form1.RadioButton1.SetChecked(false)
		gui.Form1.RadioButton2.SetChecked(true)
	}

	if Conf.EnableCheck {
		gui.Form1.ToggleBox1.SetChecked(true)
	} else {
		gui.Form1.ToggleBox1.SetChecked(false)
	}

	if Conf.EnableAuth {
		gui.Form1.ToggleBox2.SetChecked(true)
	} else {
		gui.Form1.ToggleBox2.SetChecked(false)
	}

}
