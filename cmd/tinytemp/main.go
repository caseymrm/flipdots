package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"strconv"

	"github.com/caseymrm/flipdots/panel"
	"github.com/caseymrm/flipdots/text"
)

func temperature(woeid string) int {
	url := "https://query.yahooapis.com/v1/public/yql?format=json&q=select%20item.condition%20from%20weather.forecast%20where%20woeid%20%3D%20" + woeid
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	var response struct {
		Query struct {
			Results struct {
				Channel struct {
					Item struct {
						Condition struct {
							Temp int `json:"temp,string"`
						} `json:"condition"`
					} `json:"item"`
				} `json:"channel"`
			} `json:"results"`
		} `json:"query"`
	}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&response)
	if err != nil {
		log.Fatal(err)
	}
	return response.Query.Results.Channel.Item.Condition.Temp
}

func main() {
	width := flag.Int("w", 7, "width of panel")
	height := flag.Int("h", 7, "width of panel")
	port := flag.String("p", "/dev/tty.usbserial-A505J9SE", "the serial port, empty string to simulate")
	baud := flag.Int("b", 9600, "baud rate of port")
	flag.Parse()

	p := panel.NewPanel(*width, *height, *port, *baud)
	defer p.Close()

	laWeather := temperature("2442047")
	sfWeather := temperature("2487956")
	laStr := strconv.Itoa(laWeather)
	sfStr := strconv.Itoa(sfWeather)

	f := text.GetFont(3)
	if laWeather > 9 && laWeather < 100 {
		f.Draw(p, 0, 0, string(laStr[0]))
		f.Draw(p, 4, 0, string(laStr[1]))
	}
	if sfWeather > 9 && sfWeather < 100 {
		f.Draw(p, 0, 4, string(sfStr[0]))
		f.Draw(p, 4, 4, string(sfStr[1]))
	}
	p.Send()
}
