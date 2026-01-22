package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase[T Categorizable] struct {
		name          string
		items         []T
		expectedCount map[string]int
	}

	runCases := []testCase[Weapon]{
		{
			name: "weapons grouped by melee and ranged",
			items: []Weapon{
				{Name: "Iron Sword", Cat: "melee"},
				{Name: "Oak Bow", Cat: "ranged"},
				{Name: "Steel Axe", Cat: "melee"},
			},
			expectedCount: map[string]int{
				"melee":  2,
				"ranged": 1,
			},
		},
		{
			name:          "empty slice returns empty map",
			items:         []Weapon{},
			expectedCount: map[string]int{},
		},
	}

	submitCasesPotions := []testCase[Potion]{
		{
			name: "potions grouped by healing and damage",
			items: []Potion{
				{Name: "Small Healing", Cat: "healing"},
				{Name: "Big Healing", Cat: "healing"},
				{Name: "Fire Bottle", Cat: "damage"},
			},
			expectedCount: map[string]int{
				"healing": 2,
				"damage":  1,
			},
		},
		{
			name: "mixed categories with single items",
			items: []Potion{
				{Name: "Speed Tonic", Cat: "buff"},
				{Name: "Stone Skin", Cat: "defense"},
				{Name: "Lightning Flask", Cat: "damage"},
			},
			expectedCount: map[string]int{
				"buff":    1,
				"defense": 1,
				"damage":  1,
			},
		},
	}

	passCount := 0
	failCount := 0

	for _, test := range runCases {
		result := groupByCategory(test.items)
		if !checkCountsWeapon(test.items, result, test.expectedCount) {
			failCount++
			printWeaponFailure(t, test.name, test.items, result, test.expectedCount)
		} else {
			passCount++
			printWeaponSuccess(test.name, test.items, result, test.expectedCount)
		}
	}

	skipped := len(submitCasesPotions)
	if withSubmit {
		for _, test := range submitCasesPotions {
			result := groupByCategory(test.items)
			if !checkCountsPotion(test.items, result, test.expectedCount) {
				failCount++
				printPotionFailure(t, test.name, test.items, result, test.expectedCount)
			} else {
				passCount++
				printPotionSuccess(test.name, test.items, result, test.expectedCount)
			}
		}
		skipped = 0
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}
}

func checkCountsWeapon(items []Weapon, grouped map[string][]Weapon, expected map[string]int) bool {
	totalExpected := 0
	for _, c := range expected {
		totalExpected += c
	}
	if totalExpected != len(items) {
		return false
	}

	for cat, expCount := range expected {
		actualCount := len(grouped[cat])
		if actualCount != expCount {
			return false
		}
	}

	for cat := range grouped {
		if _, ok := expected[cat]; !ok && len(grouped[cat]) > 0 {
			return false
		}
	}

	return true
}

func checkCountsPotion(items []Potion, grouped map[string][]Potion, expected map[string]int) bool {
	totalExpected := 0
	for _, c := range expected {
		totalExpected += c
	}
	if totalExpected != len(items) {
		return false
	}

	for cat, expCount := range expected {
		actualCount := len(grouped[cat])
		if actualCount != expCount {
			return false
		}
	}

	for cat := range grouped {
		if _, ok := expected[cat]; !ok && len(grouped[cat]) > 0 {
			return false
		}
	}

	return true
}

func printWeaponFailure(t *testing.T, name string, items []Weapon, grouped map[string][]Weapon, expected map[string]int) {
	failMsg := fmt.Sprintf(`---------------------------------
Test: %s

Input weapons:
`, name)
	for _, w := range items {
		failMsg += fmt.Sprintf("  - %s (Category: %s)\n", w.Name, w.Cat)
	}

	failMsg += "\nExpected category counts:\n"
	for cat, count := range expected {
		failMsg += fmt.Sprintf("  %s: %d\n", cat, count)
	}

	failMsg += "\nActual category counts:\n"
	for cat, ws := range grouped {
		failMsg += fmt.Sprintf("  %s: %d\n", cat, len(ws))
	}

	failMsg += "Fail\n"
	t.Errorf("%s", failMsg)
}

func printWeaponSuccess(name string, items []Weapon, grouped map[string][]Weapon, expected map[string]int) {
	fmt.Printf(`---------------------------------
Test: %s

Input weapons:
`, name)
	for _, w := range items {
		fmt.Printf("  - %s (Category: %s)\n", w.Name, w.Cat)
	}

	fmt.Println("\nExpected category counts:")
	for cat, count := range expected {
		fmt.Printf("  %s: %d\n", cat, count)
	}

	fmt.Println("\nActual category counts:")
	for cat, ws := range grouped {
		fmt.Printf("  %s: %d\n", cat, len(ws))
	}

	fmt.Println("Pass")
}

func printPotionFailure(t *testing.T, name string, items []Potion, grouped map[string][]Potion, expected map[string]int) {
	failMsg := fmt.Sprintf(`---------------------------------
Test: %s

Input potions:
`, name)
	for _, p := range items {
		failMsg += fmt.Sprintf("  - %s (Category: %s)\n", p.Name, p.Cat)
	}

	failMsg += "\nExpected category counts:\n"
	for cat, count := range expected {
		failMsg += fmt.Sprintf("  %s: %d\n", cat, count)
	}

	failMsg += "\nActual category counts:\n"
	for cat, ps := range grouped {
		failMsg += fmt.Sprintf("  %s: %d\n", cat, len(ps))
	}

	failMsg += "Fail\n"
	t.Errorf("%s", failMsg)
}

func printPotionSuccess(name string, items []Potion, grouped map[string][]Potion, expected map[string]int) {
	fmt.Printf(`---------------------------------
Test: %s

Input potions:
`, name)
	for _, p := range items {
		fmt.Printf("  - %s (Category: %s)\n", p.Name, p.Cat)
	}

	fmt.Println("\nExpected category counts:")
	for cat, count := range expected {
		fmt.Printf("  %s: %d\n", cat, count)
	}

	fmt.Println("\nActual category counts:")
	for cat, ps := range grouped {
		fmt.Printf("  %s: %d\n", cat, len(ps))
	}

	fmt.Println("Pass")
}

var withSubmit = true
