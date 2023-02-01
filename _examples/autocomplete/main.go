package main

import (
	"fmt"
	"sort"

	"github.com/duke-git/lancet/v2/slice"
	"github.com/sahilm/fuzzy"

	"github.com/yunginnanet/infinite/components"
)

func main() {

	suggesterOptions := []string{
		"package", "main", "func", "[]", "{}", "string", "int", "好", "我很好", "好啊",
	}

	input := components.NewInput()
	// input.Model.SetValue("你好啊啊啊")
	c := components.NewAutocomplete(func(valCtx components.AutocompleteValCtx) ([]string, bool) {
		matches := fuzzy.Find(valCtx.CursorWord(), suggesterOptions)
		if len(matches) == 0 {
			return nil, false
		}

		sort.Stable(matches)

		suggester := slice.Map[fuzzy.Match, string](matches, func(index int, item fuzzy.Match) string {
			return suggesterOptions[item.Index]
		})

		return suggester, true
	}).WithInput(input).WithSuggestionViewRender(components.TabSuggestionRender)

	err := components.NewStartUp(c).Start()
	if err != nil {
		fmt.Println(err)
	}
}
