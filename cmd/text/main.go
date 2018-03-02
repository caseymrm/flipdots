package main

import (
	"flag"
	"time"

	"github.com/caseymrm/flipdots/panel"
	"github.com/caseymrm/flipdots/text"
)

func main() {
	width := flag.Int("w", 7, "width of panel")
	height := flag.Int("h", 7, "width of panel")
	port := flag.String("p", "/dev/tty.usbserial-A505J9SE", "the serial port, empty string to simulate")
	baud := flag.Int("b", 9600, "baud rate of port")
	flag.Parse()

	p := panel.NewPanel(*width, *height, *port, *baud)
	defer p.Close()

	f := text.GetFont()
	str := "Hello world"
	for i := 0; i < len(str); i++ {
		f.Draw(p, 0, 0, string(str[i]))
		p.Send()
		time.Sleep(200 * time.Millisecond)
	}
}
