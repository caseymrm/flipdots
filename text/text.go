package text

import (
	"encoding/json"
	"errors"
	"log"
	"unicode/utf8"

	"github.com/caseymrm/flipdots/panel"
)

type Character rune

type RuneInfo struct {
	Character Character
	Bitmap    []string
}

type Font struct {
	Name          string
	Width, Height int
	Characters    []RuneInfo

	charMap map[Character][][]bool
}

var theFont *Font

// GetFont TODO: multiple size options
func GetFont(height int) *Font {
	if theFont == nil {
		var fontJSON []byte
		switch height {
		case 3:
			fontJSON = getCase3JSON()
		default:
			log.Printf("Warning, no font with height %d, using 7", height)
			fallthrough
		case 7:
			fontJSON = getVictorJSON()
		}
		var font Font
		err := json.Unmarshal(fontJSON, &font)
		if err != nil {
			log.Fatalf("Loading font: %v", err)
		}
		theFont = &font
		theFont.charMap = make(map[Character][][]bool)
		for _, info := range theFont.Characters {
			theFont.charMap[info.Character] = make([][]bool, theFont.Width)
			for _, line := range info.Bitmap {
				for x := 0; x < len(line); x++ {
					theFont.charMap[info.Character][x] = append(theFont.charMap[info.Character][x], line[x] == '0')
				}
			}
		}
	}
	return theFont
}

// MarshalJSON lets us store Character as a string
func (c *Character) MarshalJSON() ([]byte, error) {
	if !utf8.ValidRune(rune(*c)) {
		return nil, errors.New("invalid rune")
	}
	return json.Marshal(string(*c))
}

// UnmarshalJSON lets us restore a Character stored as a string
func (c *Character) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	r, size := utf8.DecodeRuneInString(s)
	if r == utf8.RuneError {
		return errors.New("rune error")
	}
	if size != len(s) {
		return errors.New("rune wrong size")
	}
	*c = Character(r)
	return nil
}

// Draw the text to the given coordinates on this panel
func (f *Font) Draw(panel *panel.Panel, panelX, panelY int, text string) {
	for i := 0; i < len(text); i++ {
		c := text[i]
		pattern, ok := f.charMap[Character(c)]
		if ok {
			for x := 0; x < f.Width; x++ {
				if panelX+x < 0 || panelX+x >= panel.Width {
					continue
				}
				for y := 0; y < f.Height; y++ {
					if panelY+y < 0 || panelY+y >= panel.Height {
						continue
					}
					panel.Set(panelX+x, panelY+y, pattern[x][y])
				}
			}
		} else {
			log.Printf("No rune found for %s in %s", string(c), f.Name)
		}
		panelX += f.Width + 1
	}
}
