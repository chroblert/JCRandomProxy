// 由res2go IDE插件自动生成。
package main

import (
	"io"
	"log"
	"os"

	"./gui"
	"github.com/ying32/govcl/vcl"
)

var mw interface{}

func init() {
	logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	mw = io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw.(io.Writer))
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
func main() {
    vcl.Application.SetScaled(true)
    vcl.Application.SetTitle("project1")
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
    vcl.Application.CreateForm(&gui.Form1)
    vcl.Application.CreateForm(&gui.Form2)
	vcl.Application.Run()
}
