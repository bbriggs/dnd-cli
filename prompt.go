package main

import (
	"github.com/c-bata/go-prompt"
	"strings"
)

func completer(d prompt.Document) []prompt.Suggest {
	if d.TextBeforeCursor() == "" {
		return []prompt.Suggest{}
	}
	args := strings.Split(d.TextBeforeCursor(), " ")
	return argumentsCompleter(args)
}

func argumentsCompleter(args []string) []prompt.Suggest {
	if len(args) <= 1 {
		return prompt.FilterHasPrefix(rootCmds, args[0], true)
	}

	first := args[0]
	switch first {
	case "get":
		if len(args) == 2 {
			second := args[1]
			subcommands := []prompt.Suggest{
				{Text: "name"},
				{Text: "hp"},
				{Text: "ac"},
				{Text: "speed"},
				{Text: "strength"},
				{Text: "dexterity"},
				{Text: "constitution"},
				{Text: "intelligence"},
				{Text: "wisdom"},
				{Text: "charisma"},
				{Text: "items"},
				{Text: "alignment"},
				// Aliases
				{Text: "str"},
				{Text: "dex"},
				{Text: "con"},
				{Text: "int"},
				{Text: "wis"},
				{Text: "cha"},
			}
			return prompt.FilterHasPrefix(subcommands, second, true)
		}
	default:
		return []prompt.Suggest{}
	}

	return []prompt.Suggest{}
}
