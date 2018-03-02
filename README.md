# Flipdots code in Go

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

![Output](https://github.com/caseymrm/flipdots/raw/master/static/simple.gif)
