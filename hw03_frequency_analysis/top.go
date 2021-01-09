package hw03_frequency_analysis //nolint:golint,stylecheck

import (
	"regexp"
	"sort"
)

func parseText(text string) []string {
	r := regexp.MustCompile("[А-Яа-яA-Za-z0-9'*?()$.,!-:]+")

	return r.FindAllString(text, -1)
}

func wordsCount(words []string) map[string]int {
	list := make(map[string]int)

	for _, m := range words {
		list[m]++
	}

	return list
}

func wordsTrimTop(words map[string]int, minLen int) []string {
	list := make([]string, 0, minLen)

	for w := range words {
		list = append(list, w)
	}

	sort.Slice(list, func(i, j int) bool {
		return words[list[i]] > words[list[j]]
	})

	if len(list) >= minLen {
		return list[:minLen]
	}

	return list
}

func Top10(text string) []string {
	if text == "" {
		return []string{}
	}

	words := parseText(text)
	wordsCountMap := wordsCount(words)

	return wordsTrimTop(wordsCountMap, 10)
}
