package main

import (
	inf "github.com/yunginnanet/infinite"
	"github.com/yunginnanet/infinite/color"
	"github.com/yunginnanet/infinite/components"
	"github.com/yunginnanet/infinite/components/selection/multiselect"
	"github.com/yunginnanet/infinite/style"
)

func main() {
	input := components.NewInput()
	input.Prompt = "Filtering: "
	input.PromptStyle = style.New().Bold().Italic().Fg(color.LightBlue)

	_, _ = inf.NewMultiSelect([]string{
		"Buy carrots",
		"Buy celery",
		"Buy kohlrabi",
		"Buy computer",
		"Buy something",
		"Buy car",
		"Buy subway",
	},
		multiselect.WithHintSymbol("x"),
		multiselect.WithUnHintSymbol("√"),
		// multiselect.WithDisableOutputResult(),
		// multiselect.WithCursorSymbol(emoji.PointRight),
		// multiselect.WithDisableFilter(),
		multiselect.WithFilterInput(input),
	).
		Display("select your items!")

	// _, _ = inf.
	//	NewMultiSelect([]string{"f1", "f2", "f3"}).
	//	Display()
}
