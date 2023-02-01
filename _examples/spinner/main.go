package main

import (
	"time"

	inf "github.com/yunginnanet/infinite"
	"github.com/yunginnanet/infinite/components"
	"github.com/yunginnanet/infinite/components/spinner"
)

func main() {
	_ = inf.NewSpinner(
		spinner.WithShape(components.Dot),
		// spinner.WithDisableOutputResult(),
	).Display(func(spinner *spinner.Spinner) {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Millisecond * 100)
			spinner.Refreshf("hello world %d", i)
		}

		spinner.Finish("finish")

		spinner.Refresh("is finish?")
	})
}
