package pokedexcli

import (
	"strings"
)

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	result := strings.Fields(text)

	return result
}
