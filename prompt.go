package main

import (
	"github.com/c-bata/go-prompt"
	"strings"
)

var rootCmds = []prompt.Suggest{
	{Text: "exit", Description: "Exit D&D CLI"},
	{Text: "get", Description: "Display various information about your character"},
	{Text: "set", Description: "Set or update various attributes"},
	{Text: "add", Description: "Increment a value. Use this for tracking things that fluctuate frequently, such as HP."},
	{Text: "sub", Description: "Decrement a value. Use this for tracking things that fluctuate frequently, such as HP."},
}

var secondSuggestions = []prompt.Suggest{
	{Text: "alignment"},
	{Text: "class"},
	{Text: "eyes"},
	{Text: "hair"},
	{Text: "name"},
	{Text: "race"},
	{Text: "size"},
	{Text: "skin"},
	{Text: "weight"},
	{Text: "ac"},
	{Text: "acrobatics"},
	{Text: "age"},
	{Text: "animalHandling"},
	{Text: "arcana"},
	{Text: "athletics"},
	{Text: "cha"},
	{Text: "charisma"},
	{Text: "con"},
	{Text: "constitution"},
	{Text: "cp"},
	{Text: "deception"},
	{Text: "dex"},
	{Text: "dexterity"},
	{Text: "ep"},
	{Text: "exp"},
	{Text: "gp"},
	{Text: "height"},
	{Text: "history"},
	{Text: "hp"},
	{Text: "insight"},
	{Text: "int"},
	{Text: "intelligence"},
	{Text: "intimidation"},
	{Text: "level"},
	{Text: "medicine"},
	{Text: "nature"},
	{Text: "passiveWisdom"},
	{Text: "perception"},
	{Text: "persuasion"},
	{Text: "pp"},
	{Text: "religion"},
	{Text: "sleightOfHand"},
	{Text: "sp"},
	{Text: "speed"},
	{Text: "stealth"},
	{Text: "str"},
	{Text: "strength"},
	{Text: "survival"},
	{Text: "tempHP"},
	{Text: "wis"},
	{Text: "wisdom"},
	{Text: "xp"},
}

var intAttrs = []string{
	"ac",
	"acrobatics",
	"age",
	"animalHandling",
	"arcana",
	"athletics",
	"cha",
	"charisma",
	"con",
	"constitution",
	"cp",
	"deception",
	"dex",
	"dexterity",
	"ep",
	"exp",
	"gp",
	"height",
	"history",
	"hp",
	"insight",
	"int",
	"intelligence",
	"intimidation",
	"investigation",
	"level",
	"medicine",
	"nature",
	"passiveWisdom",
	"perception",
	"persuasion",
	"pp",
	"religion",
	"sleightOfHand",
	"sp",
	"speed",
	"stealth",
	"str",
	"strength",
	"survival",
	"tempHP",
	"wis",
	"wisdom",
	"xp"}

var stringSliceAttrs = []string{
	"traits",
	"ideals",
	"bonds",
	"flaws",
	"features",
	"items",
	"equipment"}

var stringAttrs = []string{
	"alignment",
	"class",
	"eyes",
	"hair",
	"name",
	"race",
	"size",
	"skin",
	"weight"}

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
			return prompt.FilterHasPrefix(secondSuggestions, second, true)
		}
	case "set":
		if len(args) == 2 {
			second := args[1]
			return prompt.FilterHasPrefix(secondSuggestions, second, true)
		}
	default:
		return rootCmds
	}

	return rootCmds
}
