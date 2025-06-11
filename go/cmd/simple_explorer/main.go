package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	leftStyle = lipgloss.NewStyle().
		Width(30).
		Height(40).
		Padding(1, 2, 1, 2)

	rightStyle = lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#89dceb")).
		Padding(1, 2).
		Height(40).
		Width(50)
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
	left := leftStyle.Render(m.list.View())

	desc := ""
	if item := m.selectedItem(); item != nil {
		desc = item.Desc
	}
	right := rightStyle.Render(desc)

	return lipgloss.JoinHorizontal(lipgloss.Top, left, right)
}

func (m model) selectedItem() *Item {
	if selected, ok := m.list.SelectedItem().(Item); ok {
		return &selected
	}
	return nil
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

	const defaultWidth = 30
	l := list.New(items, list.NewDefaultDelegate(), defaultWidth, 40)
	l.Title = "Stuff"

	m := model{list: l}

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Println("Oops! We fucked up:", err)
		os.Exit(1)
	}
}
