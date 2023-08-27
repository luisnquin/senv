package prompt

import (
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/manifoldco/promptui/list"
)

type option struct {
	Name     string
	Selected bool
}

const defaultSelectPromptSize = 10

func ListSelector(label string, options []string, selected string) (string, bool) {
	opts := make([]option, len(options))

	for i, opt := range options {
		opts[i] = option{
			Name:     opt,
			Selected: opt == selected,
		}
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}:",
		Active:   `▸ {{ .Name | cyan | underline }} {{if .Selected}}{{ "(current)" | magenta }}{{end}}`,
		Inactive: `  {{ .Name | cyan }} {{if .Selected}}{{ "(current)" | magenta }}{{end}}`,
		Selected: "▸ {{ .Name | cyan }}",
	}

	s := promptui.Select{
		Label:             label,
		Items:             opts,
		Searcher:          searcher(opts),
		Templates:         templates,
		StartInSearchMode: true,
		HideHelp:          true,
		Size:              defaultSelectPromptSize,
		HideSelected:      true,
	}

	i, _, err := s.Run()
	if err != nil {
		return "", false
	}

	return opts[i].Name, true
}

func searcher(targets []option) list.Searcher {
	return func(input string, index int) bool {
		if strings.Contains(strings.ToLower(targets[index].Name), input) {
			return true
		}

		return false
	}
}
