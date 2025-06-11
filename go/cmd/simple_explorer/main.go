package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type Item struct {
	Name, Description string
}

func getSampleItems() []Item {
	return []Item{
		{Name: "Cup of tea", Description: "Nothing like a hot cup a tea."},
		{Name: "Computer", Description: `Some people would say "com-pju-?-e"`},
		{Name: "Keyboard", Description: `I prefer the split ones`},
		{Name: "Plant", Description: `I love them! And thy look really nice in my home or garden. They also purify the air in your flat.`},
		{Name: "Beans", Description: `Beans are tasty! I like a good plate of beans in the winter.`},
		{Name: "Coffee", Description: `I can't drink it anymore, it makes my tummy ache. I never really liked it anyway`},
		{Name: "Sugar", Description: `Sugar may taste good, but it isn't very good for your health. Actually, it is pretty bad!`},
		{Name: "Stairs", Description: `Stairs are very useful when you need to go up somewhere. They can also be helpful when you need to go down lol.`},
	}
}

func main() {
	p := tea.NewProgram(nil)
	if _, err := p.Run(); err != nil {
		fmt.Println("Oops! We fucked up:", err)
		os.Exit(1)
	}
}
