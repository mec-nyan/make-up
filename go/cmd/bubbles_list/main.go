package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model struct {
	list list.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "x":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return docStyle.Render(m.list.View())
}

func main() {
	items := []list.Item{
		item{title: "Blueberry jam", desc: "I like blueberry jam."},
		item{title: "Mate cocido", desc: "Mate cocido is a South American beverage."},
		item{title: "Cup of tea", desc: "Can anything be any more British?"},
		item{title: "Cookie jar", desc: "My cookie jar is almost empty!"},
		item{title: "Neovim", desc: "Best text editor, no doubt about it!"},
		item{title: "Vim", desc: "Best text editor, no doubt about it!"},
		item{title: "Emacs", desc: "Best text editor, no doubt about it!"},
		item{title: "Bluetooth ear buds", desc: "Bluetooth ear buds suck dude."},
		item{title: "Split keyboard", desc: "I need one of those now!"},
		item{title: "AI generated code", desc: "This isn't one of those."},
		item{title: "Notebook", desc: "Real, paper notebooks are good for writing stuff."},
	}

	m := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = "Stuff I like"

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Shit! Something went wrong...:", err)
		os.Exit(1)
	}
}
