package gui

import (
	"fmt"
	"strings"

	"github.com/ying32/govcl/vcl"
)

//::private::
type TForm2Fields struct {
}

func (f *TForm2) OnButton1Click(sender vcl.IObject) {
	protocol := f.ComboBox1.Text()
	ipPort := f.Edit1.Text()
	// Form1.ListView1.Items().BeginUpdate()
	item := Form1.ListView1.Items().Add()
	item.SetCaption(fmt.Sprintf("%d", Form1.ListView1.Items().Count()))
	item.SubItems().Add(fmt.Sprintf("%s", protocol))
	item.SubItems().Add(fmt.Sprintf("%s", strings.Split(ipPort, ":")[0]))
	item.SubItems().Add(fmt.Sprintf("%s", strings.Split(ipPort, ":")[1]))
	item.MakeVisible(true)
	f.Close()
	// Form1.ListView1.Items().EndUpdate()
}

func (f *TForm2) OnButton2Click(sender vcl.IObject) {
	f.Close()
}
