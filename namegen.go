package namegen

import (
	"math/rand"
)

func randomItem(items []string) string {
	return items[rand.Intn(len(items))]
}

func GetNameByType(origin string) string {
	firstNames := []string{}
	lastNames := []string{}

	if origin == "english" {
		firstNames = englishFirstNames
		lastNames = englishLastNames
	} else if origin == "spanish" {
		firstNames = spanishFirstNames
		lastNames = spanishLastNames
	}

	firstName := randomItem(firstNames)
	lastName := randomItem(lastNames)
	return firstName + " " + lastName
}
