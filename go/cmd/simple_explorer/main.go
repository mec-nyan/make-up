package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/list"
)

var (
	rounded = lipgloss.Border{
		Top:          "━",
		Bottom:       "━",
		Left:         "┃",
		Right:        "┃",
		TopLeft:      "╭",
		TopRight:     "╮",
		BottomLeft:   "╰",
		BottomRight:  "╯",
		MiddleLeft:   "",
		MiddleRight:  "",
		Middle:       "",
		MiddleTop:    "",
		MiddleBottom: "",
	}

	leftStyle = lipgloss.NewStyle().
		// Margin(1, 2).
		Padding(1, 2)

	rightStyle = lipgloss.NewStyle().
			Border(rounded).
			BorderForeground(lipgloss.Color("#cba6f7")).
			Margin(1, 1, 1, 0).
			Padding(1, 2)
)

type (
	Item struct {
		Name, Desc string
	}
	model struct {
		list          list.List
		width, height int
	}
)

func (i Item) Title() string { return i.Name }

func (i Item) Description() string { return i.Desc }

func (i Item) FilterValue() string { return i.Name }

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		return m, nil
	default:
		return m, tea.Quit
	}
}

func (m model) View() string {
	leftWidth := m.width / 3
	rightWidth := m.width - leftWidth

	innerLeftWidth := leftWidth - leftStyle.GetHorizontalFrameSize() +
		leftStyle.GetPaddingLeft() + leftStyle.GetPaddingRight()

	innerLeftHeight := m.height - leftStyle.GetVerticalFrameSize() +
		leftStyle.GetPaddingTop() + leftStyle.GetPaddingBottom()

	leftStyle = leftStyle.Width(innerLeftWidth).Height(innerLeftHeight)

	innerRightWidth := rightWidth - rightStyle.GetHorizontalFrameSize() +
		rightStyle.GetPaddingLeft() + rightStyle.GetPaddingRight()

	// innerRightHeight := m.height - rightStyle.GetVerticalFrameSize() +
	// 	rightStyle.GetPaddingTop() + rightStyle.GetPaddingBottom()

	rightStyle = rightStyle.Width(innerRightWidth)

	left := leftStyle.Render(m.list.String())

	desc := ""
	if item := m.selectedItem(); item != nil {
		desc = item.Desc
	}
	right := rightStyle.Render(desc)

	return lipgloss.JoinHorizontal(lipgloss.Top, left, right)
}

func (m model) selectedItem() *Item {
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

func SelectedEnum(_ list.Items, i int) string {
	if i == 0 {
		return "> "
	}
	return ""
}

func main() {
	items := list.New()
	for _, i := range getSampleItems() {
		items.Item(i.Name)
	}
	items.Enumerator(SelectedEnum)

	m := model{list: *items}

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Oops! We fucked up:", err)
		os.Exit(1)
	}
}
