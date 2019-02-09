package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
	"gopkg.in/yaml.v2"
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

var rootCmds = []prompt.Suggest{
	{Text: "exit", Description: "Exit D&D CLI"},
	{Text: "get", Description: "Display various information about your character"},
	{Text: "set", Description: "Set or update various attributes"},
	{Text: "add", Description: "Increment a value. Use this for tracking things that fluctuate frequently, such as HP."},
	{Text: "sub", Description: "Decrement a value. Use this for tracking things that fluctuate frequently, such as HP."},
}

func (c *Character) executor(in string) {
	in = strings.TrimSpace(in)
	blocks := strings.Split(in, " ")
	switch blocks[0] {
	case "help":
		fmt.Println(HELP)
		return
	case "get":
		if len(blocks) > 1 {
			switch blocks[1] {
			case "items":
				c.printItems()
				return
			case "alignment":
				fmt.Println(c.Alignment)
			case "name":
				fmt.Println(c.Name)
			default:
				attr, err := c.getAttr(blocks[1])
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println(attr)
				}
			}
		} else {
			return
		}
	case "exit":
		os.Exit(0)
	default:
		return
	}
}

func (c *Character) printItems() {
	for _, i := range c.Items {
		fmt.Println("- " + i)
	}
}

func (c *Character) getAttr(attr string) (int, error) {
	attrs := map[string]int{
		"hp":           c.HP,
		"ac":           c.AC,
		"speed":        c.Speed,
		"str":          c.Str,
		"strength":     c.Str,
		"dex":          c.Dex,
		"dexterity":    c.Dex,
		"con":          c.Con,
		"constitution": c.Con,
		"int":          c.Int,
		"intelligence": c.Int,
		"wis":          c.Wis,
		"wisdom":       c.Wis,
		"cha":          c.Cha,
		"charisma":     c.Cha,
	}

	for k, v := range attrs {
		if k == attr {
			return v, nil
		}
	}
	return 0, fmt.Errorf("Attribute not found.")
}

func readState(config string) (*Character, error) {
	c := Character{}
	b, err := ioutil.ReadFile(config)
	if err != nil {
		return &c, err
	}

	err = yaml.Unmarshal([]byte(b), &c)
	return &c, err
}

func dumpState(c Character) error {
	d, err := yaml.Marshal(&c)
	if err != nil {
		fmt.Println("Unable to save state!")
		fmt.Println(err.Error())
		return err
	}

	err = ioutil.WriteFile(fmt.Sprintf("%s.yaml", c.Name), d, 0666)
	if err != nil {
		fmt.Println("Unable to save state!")
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func main() {
	fmt.Println("D&D CLI -- Because D&D alone wasn't nerdy enough")
	c, err := readState("config.yaml")
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
