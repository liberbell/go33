package simpleGoApp

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

type UI struct {
}

func NewUI() *UI {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	return &UI{}
}

func (ui *UI) Render(data *Data) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	for i, element := range data.Elements {
		printAt(0, i, fmt.Sprintf("Element %d: %d", i+1, element), termbox.ColorWhite, termbox.ColorBlack)
	}
}
