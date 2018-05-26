package panel

import (
	"image/color"
	"log"

	"github.com/tarm/serial"
)

// Panel represents a single flipdot panel
type Panel struct {
	Address []byte // nil implies broadcast, okay if just one panel

	Width  int
	Height int

	State [][]bool

	Port *serial.Port
}

// NewPanel returns a new Panel with the given size, attached to the given port. The panel's Close() should be called when done.
func NewPanel(w, h int, portName string, portBaud int) *Panel {
	panel := &Panel{
		Width:  w,
		Height: h,
		State:  make([][]bool, w),
	}
	for i := 0; i < w; i++ {
		panel.State[i] = make([]bool, h)
	}

	if portName == "" || portBaud == 0 {
		log.Printf("Running in debug mode, with no panel connection")
		return panel
	}
	var err error
	panel.Port, err = serial.OpenPort(&serial.Config{Name: portName, Baud: portBaud})
	if err != nil {
		log.Fatal(err)
	}
	return panel
}

// Close the serial port
func (p *Panel) Close() {
	if p.Port != nil {
		p.Port.Close()
	}
	p.Port = nil
}

// Send the state of the board to the associated flip dot panel and refresh
func (p *Panel) Send() {
	p.sendBoard(true)
}

// Queue the state of the board to the panel, show when Refresh() is called (used for multiple panels)
func (p *Panel) Queue() {
	p.sendBoard(false)
}

func (p *Panel) sendBoard(refresh bool) {
	data := make([]byte, p.Width)
	for x := 0; x < p.Width; x++ {
		d := 0
		for y := 0; y < p.Height; y++ {
			d = d<<1 | int(p.GetInt(x, y))
		}
		data[x] = byte(d)
	}
	p.sendData(p.Address, data, refresh)
}

// Refresh causes any queued state to be displayed
func (p *Panel) Refresh() {
	p.sendData(nil, nil, true)
}

func (p *Panel) PrintState() {
	for y := 0; y < p.Height; y++ {
		line := ""
		for x := 0; x < p.Width; x++ {
			if p.Get(x, y) {
				line += "⚫️"
			} else {
				line += "⚪️"
			}
		}
		log.Println(line)
	}
}

func (p *Panel) sendData(address, data []byte, refresh bool) {
	if address == nil {
		address = []byte{0xff}
	}
	if data == nil {
		data = []byte{}
	}
	command := byte(0)
	switch len(data) {
	case 112:
		if refresh {
			command = 0x82
		} else {
			command = 0x81
		}
	case 56:
		if refresh {
			command = 0x85
		} else {
			command = 0x86
		}
	case 28:
		if refresh {
			command = 0x83
		} else {
			command = 0x84
		}
	case 14:
		if refresh {
			command = 0x92
		} else {
			command = 0x93
		}
	case 7:
		if refresh {
			command = 0x87
		} else {
			command = 0x88
		}
	case 0:
		command = 0x82
		address = []byte{}
	}
	if command == 0 {
		log.Fatalf("Unknown byte length %d", len(data))
	}
	message := append([]byte{0x80}, command)
	message = append(message, address...)
	message = append(message, data...)
	message = append(message, 0x8f)

	if p.Port == nil {
		log.Printf("Message: %x", message)
		p.PrintState()
		return
	}

	n, err := p.Port.Write(message)
	if err != nil {
		log.Fatal(err)
	}
	if n != len(address)+len(data)+3 {
		log.Fatalf("Didn't send all bytes, only %d", n)
	}
}

// Get the state of the dot at the given coordinate as a boolean
func (p *Panel) Get(x, y int) bool {
	return p.State[x][y]
}

// GetInt returns the state of the dot at the given coordinate as an int
func (p *Panel) GetInt(x, y int) int {
	if p.State[x][y] {
		return 1
	}
	return 0
}

// Color of the given dot- currently assumes black and white
func (p *Panel) Color(x, y int) color.RGBA {
	if p.Get(x, y) {
		return color.RGBA{255, 255, 255, 0}
	}
	return color.RGBA{0, 0, 0, 0}
}

// Set the given coordinate dot on or off
func (p *Panel) Set(x, y int, state bool) {
	if x < 0 || x >= p.Width {
		log.Printf("WARNING: Skipping Set() with x %d out of range [0, %d)", x, p.Width)
		return
	}
	if y < 0 || y >= p.Height {
		log.Printf("WARNING: Skipping Set() with y %d out of range [0, %d)", x, p.Height)
		return
	}
	p.State[x][y] = state
}

// Clear the dots on or off
func (p *Panel) Clear(state bool) {
	for x := 0; x < p.Width; x++ {
		for y := 0; y < p.Height; y++ {
			p.Set(x, y, state)
		}
	}
}
