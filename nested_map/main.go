package main


func getNameCounts(names []string) map[rune]map[string]int {
	new_map := make(map[rune]map[string]int)
	for _, name := range names {
		runes := []rune(name)
		if _, ok := new_map[runes[0]]; !ok {
			new_map[runes[0]] = make(map[string]int)
		}
		new_map[runes[0]][name]++
	}
	return new_map
}
