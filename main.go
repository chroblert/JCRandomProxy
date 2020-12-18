// 由res2go IDE插件自动生成。
package main

import (
	// "io"

	log "github.com/chroblert/JCRandomProxy-GUI/Logs"
	// "os"
	"github.com/chroblert/JCRandomProxy-GUI/Conf"
	"github.com/chroblert/JCRandomProxy-GUI/gui"
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
	vcl.Application.Run()
}
