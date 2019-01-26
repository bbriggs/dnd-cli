package main

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
)

type Character struct {
	Name      string   `yaml:"name"`
	HP        int      `yaml:"hp"`
	AC        int      `yaml:"ac"`
	Speed     int      `yaml:"speed"`
	Str       int      `yaml:"str"`
	Dex       int      `yaml:"dex"`
	Con       int      `yaml:"con"`
	Int       int      `yaml:"int"`
	Wis       int      `yaml:"wis"`
	Cha       int      `yaml:"cha"`
	Items     []string `yaml:"items"`
	Alignment string   `yaml:"alignment"`
}

const HELP = "Generic help text"

func (c *Character) LivePrefix() (string, bool) {
	if c.Name == "" {
		return "(none) > ", true
	}
	return c.Name + "> ", true
}

var suggestions = []prompt.Suggest{
	{Text: "exit", Description: "Exit D&D CLI"},
	{Text: "stats", Description: "Display your base stats"},
	{Text: "hp", Description: "Show your current hit points"},
	{Text: "help", Description: "Get help"},
	{Text: "name", Description: "Display your character's name"},
	{Text: "items", Description: "Check your bag to see what you're carrying"},
	{Text: "alignment", Description: "In case you forgot how evil you're supposed to be"},
}

func completer(in prompt.Document) []prompt.Suggest {
	w := in.GetWordBeforeCursor()
	if w == "" {
		return []prompt.Suggest{}
	}
	return prompt.FilterHasPrefix(suggestions, w, true)
}

func (c *Character) executor(in string) {
	in = strings.TrimSpace(in)
	blocks := strings.Split(in, " ")
	switch blocks[0] {
	case "help":
		fmt.Println(HELP)
		return
	case "exit":
		os.Exit(0)
	case "hp":
		fmt.Println(c.HP)
		return
	case "name":
		fmt.Println(c.Name)
		return
	case "items":
		for _, element := range c.Items {
			fmt.Println("-", element)
		}
		return
	case "alignment":
		fmt.Println(c.Alignment)
		return
	default:
		return
	}
}

func readConfig(config string) (*Character, error) {
	c := Character{}
	b, err := ioutil.ReadFile(config)
	if err != nil {
		return &c, err
	}

	err = yaml.Unmarshal([]byte(b), &c)
	return &c, err
}

func main() {
	fmt.Println("D&D CLI -- Because D&D alone wasn't nerdy enough")
	c, err := readConfig("config.yaml")
	if err != nil {
		panic(err)
	}
	p := prompt.New(
		c.executor,
		completer,
		prompt.OptionLivePrefix(c.LivePrefix),
		prompt.OptionTitle("dnd-util"),
	)
	p.Run()
}
