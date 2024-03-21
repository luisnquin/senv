package cmd

import (
	"bufio"
	"bytes"
	"os"
	"sort"
	"strings"

	"github.com/armon/go-radix"
	"github.com/samber/lo"
)

// Returns an empty string if not found.
func getActiveEnvironment(envFilePath string) string {
	data, err := os.ReadFile(envFilePath)
	if err != nil {
		return ""
	}

	s := bufio.NewScanner(bytes.NewReader(data))

	for s.Scan() {
		line := s.Text()

		if len(rxSenvDotEnvComment.FindAllString(line, -1)) == 2 {
			return rxSenvDotEnvComment.ReplaceAllString(line, "")
		}
	}

	return ""
}

func groupAndSortByPrefix(entries map[string]any) []map[string]any {
	rTree := buildRadixTreeFromMap(entries)
	prefixEntries, withoutGroup := groupEntriesByPrefix(entries, rTree)

	result := make([]map[string]any, 0, len(prefixEntries))

	prefixes := getSortedPrefixes(prefixEntries)

	for _, prefix := range prefixes {
		v := prefixEntries[prefix]
		if len(v) == 1 {
			for k2, v2 := range v {
				withoutGroup[k2] = v2
			}
		} else {
			result = append(result, v)
		}
	}

	if len(withoutGroup) > 0 {
		result = append(result, withoutGroup)
	}

	return result
}

func buildRadixTreeFromMap(entries map[string]any) *radix.Tree {
	rTree := radix.New()

	for k := range entries {
		var b strings.Builder

		fragments := strings.Split(k, "_")
		if len(fragments) > 1 {
			fragments = fragments[:len(fragments)-1]
		}

		for _, s := range fragments {
			b.WriteString(s)
			b.WriteString("_")

			rTree.Insert(b.String(), nil)
		}
	}

	return rTree
}

func groupEntriesByPrefix(entries map[string]any, rTree *radix.Tree) (map[string]map[string]any, map[string]any) {
	prefixEntries := make(map[string]map[string]any)
	withoutGroup := make(map[string]any)

	for k, v := range entries {
		if prefix, _, ok := rTree.LongestPrefix(k); ok {
			if _, ok := prefixEntries[prefix]; !ok {
				prefixEntries[prefix] = map[string]any{k: v}
			} else {
				prefixEntries[prefix][k] = v
			}
		} else {
			withoutGroup[k] = v
		}
	}

	return prefixEntries, withoutGroup
}

func getSortedPrefixes(prefixEntries map[string]map[string]any) []string {
	prefixes := lo.Keys(prefixEntries)
	sort.Strings(prefixes)

	return prefixes
}
