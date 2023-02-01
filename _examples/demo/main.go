package main

import (
	"fmt"
	"time"

	"github.com/duke-git/lancet/v2/random"
	"github.com/duke-git/lancet/v2/strutil"

	inf "github.com/yunginnanet/infinite"
	"github.com/yunginnanet/infinite/color"
	"github.com/yunginnanet/infinite/components"
	"github.com/yunginnanet/infinite/components/selection/confirm"
	"github.com/yunginnanet/infinite/components/selection/multiselect"

	"github.com/yunginnanet/infinite/components/spinner"
	"github.com/yunginnanet/infinite/style"
)

func main() {

	options := []string{
		"Jay Chou - Daoxiang",

		"Jay Chou-Sunny Day",

		"Hemingway-The Old Man and the Sea",

		"Xu Wei-I am as free as the wind",

		"Wu Bai-White Pigeon",

		"Wu Bai-Maple Leaf",

		"Eason Chan-Best Bad Friend",

		"Eason Chan-Under Mount Fuji",

		"Eason Chan-Elimination",

		"Eason Chan-naked",

		"New Pants Band - Fireworks",

		"New Pants - Do You Want To Dance",

		"New Pants Band-Life is hot because of you",

		"The New Pants Band-People without ideals are not sad",

		"Nezha Band - Naohai"}

	inf.NewSpinner(
		spinner.WithPrompt(" Loading..."),
		spinner.WithDisableOutputResult(),
	).Display(func(spinner *spinner.Spinner) {
		time.Sleep(time.Millisecond * 100 * 12)
		spinner.Info("A total of %d songs were found", len(options))
	})

	input := components.NewInput()
	input.Prompt = "Filtering: "
	input.PromptStyle = style.New().Bold().Italic().Fg(color.LightBlue)

	selected, _ := inf.NewMultiSelect(options,
		multiselect.WithFilterInput(input),
	).Display("Please select the song you want to download")

	yes, _ := inf.NewConfirmWithSelection(
		confirm.WithPrompt(fmt.Sprintf("Do you want to download this %d song", len(selected))),
	).Display()

	if !yes {
		return
	}

	inf.NewProgressGroup(len(selected)).
		AppendRunner(func(progress *components.Progress) func() {
			title := strutil.After(options[selected[progress.Id-1]], "-")
			total := random.RandInt(10, 20)
			progress.WithTotal(int64(total))
			progress.WithDefaultGradient()
			progress.WithWidth(80)
			progress.WithTitleView(func(done bool) string {
				if done {
					return fmt.Sprintf("Download %s succeeded", title)
				}
				return fmt.Sprintf("Download %s ...", title)
			})
			return func() {
				for i := 0; i < total+1; i++ {
					progress.IncrOne()
					sleep()
				}
			}
		}).Display()

}

func sleep() {
	time.Sleep(time.Millisecond * 100)
}
