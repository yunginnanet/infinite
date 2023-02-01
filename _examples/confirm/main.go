package main

import (
	"fmt"

	inf "github.com/yunginnanet/infinite"
	"github.com/yunginnanet/infinite/components/input/confirm"
)

func main() {

	c := inf.NewConfirm(
		confirm.WithPure(),
		confirm.WithDefaultYes(),
		// confirm.WithDisableOutputResult(),
		// confirm.WithPrompt("hello world?"),
		confirm.WithDisplayHelp(),
		// confirm.WithSymbol(emoji.Question),
		// confirm.WithKeyMap(
		//	confirm.KeyMap{Quit: key.NewBinding(
		//		key.WithKeys("c"))},
		// ),
	)

	c.Display()

	if c.Value() {
		fmt.Println("yes, you are.")
	} else {
		fmt.Println("no,you are not.")
	}
}
