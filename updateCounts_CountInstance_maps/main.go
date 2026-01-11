package main

func updateCounts(messagedUsers []string, validUsers map[string]int) {
	for _, message := range messagedUsers {
		if _, ok := validUsers[message]; ok {
			validUsers[message]++
		}
	}
}
