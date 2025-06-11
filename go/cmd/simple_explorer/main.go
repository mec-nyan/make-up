package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	Item struct {
		Name, Desc string
	}
	model struct {
		list list.Model
	}
)

func (i Item) Title() string       { return i.Name }

func (i Item) Description() string { return i.Desc }

func (i Item) FilterValue() string { return i.Name }

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return m.list.View()
}

func getSampleItems() []Item {
	return []Item{
		{Name: "Cup of tea", Desc: "Nothing like a hot cup a tea."},
		{Name: "Computer", Desc: `Some people would say "com-pju-?-e"`},
		{Name: "Keyboard", Desc: `I prefer the split ones`},
		{Name: "Plant", Desc: `I love them! And thy look really nice in my home or garden. They also purify the air in your flat.`},
		{Name: "Beans", Desc: `Beans are tasty! I like a good plate of beans in the winter.`},
		{Name: "Coffee", Desc: `I can't drink it anymore, it makes my tummy ache. I never really liked it anyway`},
		{Name: "Sugar", Desc: `Sugar may taste good, but it isn't very good for your health. Actually, it is pretty bad!`},
		{Name: "Stairs", Desc: `Stairs are very useful when you need to go up somewhere. They can also be helpful when you need to go down lol.`},
	}
}

func main() {
	var items []list.Item
	for _, i := range getSampleItems() {
		items = append(items, i)
	}

	const defaultWidth = 20
	l := list.New(items, list.NewDefaultDelegate(), defaultWidth, 20)
	l.Title = "Stuff"

	m := model{list: l}

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Println("Oops! We fucked up:", err)
		os.Exit(1)
	}
}
