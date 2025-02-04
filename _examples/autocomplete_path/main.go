package main

import (
	"path/filepath"
	"sort"

	"github.com/duke-git/lancet/v2/slice"
	"github.com/sahilm/fuzzy"

	"github.com/yunginnanet/infinite/components"
)

func main() {
	var f components.Suggester = func(valCtx components.AutocompleteValCtx) ([]string, bool) {
		cursorWord := valCtx.CursorWord()
		files, err := filepath.Glob(cursorWord + "*")
		if err != nil {
			return nil, false
		}

		matches := fuzzy.Find(cursorWord, files)
		if len(matches) == 0 {
			return nil, false
		}

		sort.Stable(matches)

		suggester := slice.Map[fuzzy.Match, string](matches, func(index int, item fuzzy.Match) string {
			return files[item.Index]
		})
		return suggester, true
	}

	c := components.NewAutocomplete(f)

	components.NewStartUp(c).Start()
}
