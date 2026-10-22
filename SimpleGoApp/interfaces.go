package simpleGoApp

import "github.com/nsf/termbox-go"

type UI struct {
}

func NewUI() *UI {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	return &UI{}
}
