package main

import (
	"flag"
	"log"
	"math/rand"
	"time"

	"github.com/caseymrm/flipdots/panel"
)

func main() {
	width := flag.Int("w", 7, "width of panel")
	height := flag.Int("h", 7, "width of panel")
	port := flag.String("p", "/dev/tty.usbserial-A505J9SE", "the serial port")
	baud := flag.Int("b", 9600, "baud rate of port")
	flag.Parse()

	p := panel.NewPanel(*width, *height, *port, *baud)
	defer p.Close()

	for i := 0; i < 20; i++ {
		for x := 0; x < p.Width; x++ {
			for y := 0; y < p.Height; y++ {
				p.Set(x, y, rand.Intn(2) == 0)
			}
		}
		log.Printf("Sending panel %d/20", i+1)
		p.Send()
		time.Sleep(200 * time.Millisecond)
	}
}
