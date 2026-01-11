package main

import (
	"strings"
)

func countDistinctWords(messages []string) int {

	temp_map := make(map[string]struct{})

	for _, message := range messages {
		for _, word := range strings.Fields(strings.ToLower(message)) {
			temp_map[word] = struct{}{}
		}
	}

	return len(temp_map)
}
