package main

import (
	"fmt"

	inf "github.com/yunginnanet/infinite"
	"github.com/yunginnanet/infinite/components/selection/confirm"
)

func main() {

	val, _ := inf.NewConfirmWithSelection(
		// confirm.WithDisableOutputResult(),
		// confirm.WithDisableShowHelp(),
		confirm.WithDefaultYes(),
	).Display()

	if val {
		fmt.Println("yes, you are.")
	} else {
		fmt.Println("no,you are not.")
	}
}
