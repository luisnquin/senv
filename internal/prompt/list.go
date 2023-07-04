package prompt

import (
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/manifoldco/promptui/list"
)

func ListSelector(label string, options []string) (string, bool) {
	s := promptui.Select{
		Label:             label,
		Items:             options,
		Searcher:          searcher(options),
		StartInSearchMode: true,
		HideHelp:          true,
		Size:              10,
		HideSelected:      true,
	}

	_, selected, err := s.Run()
	if err != nil {
		return "", false
	}

	return selected, true
}

func searcher(targets []string) list.Searcher {
	return func(input string, index int) bool {
		if strings.Contains(strings.ToLower(targets[index]), input) {
			return true
		}

		return false
	}
}
