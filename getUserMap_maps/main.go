package main

import "errors"

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	new_map := make(map[string]user)
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	}

	for i := 0; i < len(names); i++ {
		new_map[names[i]] = user{names[i], phoneNumbers[i]}
	}

	return new_map, nil

}

type user struct {
	name        string
	phoneNumber int
}
