package main

import (
	"fmt"

	term "github.com/nsf/termbox-go"
)

func main() {
	err := term.Init()
	if err != nil {
		panic(err)
	}
	defer term.Close()

	fmt.Println("Press F1 to exit")
	for {
		switch ev := term.PollEvent(); ev.Type {
		case term.EventKey:
			switch ev.Key {
			case term.KeyEsc:
				fmt.Print("<ESC>")
			case term.KeyF1:
				return
			case term.KeyEnter:
				fmt.Println("<ENTER>")
			case term.KeyTab:
				fmt.Print("<TAB>")
			default:
				fmt.Printf("%c", ev.Ch)
			}
		case term.EventError:
			panic(ev.Err)
		}
	}
}
