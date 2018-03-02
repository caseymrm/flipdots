# Flipdots in Go

This library handles controlling [Flip-Dot Boards by Alfa-Zeta](https://flipdots.com/en/products-services/flip-dot-boards-xy5/) from Go. It should work on Mac, Linux, and Windows, but has been primarily tested on Mac with a 7x7 USB panel.

## Simple example:

```golang
package main

import (
    "math/rand"
    "time"

    "github.com/caseymrm/flipdots/panel"
)

func main() {
    p := panel.NewPanel(7, 7, "/dev/tty.usbserial-A505J9SE", 9600)
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
```

## Letter drawing:

```golang
package main

import (
	"flag"
	"time"

	"github.com/caseymrm/flipdots/panel"
	"github.com/caseymrm/flipdots/text"
)

func main() {
    p := panel.NewPanel(7, 7, "/dev/tty.usbserial-A505J9SE", 9600)
	defer p.Close()

	f := text.GetFont()
	str := "Hello world"
	for i := 0; i < len(str); i++ {
		f.Draw(p, 0, 0, string(str[i]))
		p.Send()
		time.Sleep(200 * time.Millisecond)
	}
}
```
![Output](https://github.com/caseymrm/flipdots/raw/master/static/simple.gif)
