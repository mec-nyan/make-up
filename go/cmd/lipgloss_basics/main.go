package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type (
	colorName = string
	hexValue  = string
)

var palette = map[colorName]hexValue{
	"rosewater": "#f5e0dc",
	"flamingo":  "#f2cdcd",
	"pink":      "#f5c2e7",
	"mauve":     "#cba6f7",
	"red":       "#f38ba8",
	"maroon":    "#eba0ac",
	"peach":     "#fab387",
	"yellow":    "#f9e2af",
	"green":     "#a6e3a1",
	"teal":      "#94e2d5",
	"sky":       "#89dceb",
	"sapphire":  "#74c7ec",
	"blue":      "#89b4fa",
	"lavender":  "#b4befe",
	"text":      "#cdd6e4",
	"base":      "#1e1e2e",
}

type button struct {
	content string
	width   int
}

type model struct {
	width int
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit
	case tea.WindowSizeMsg:
		m.width = msg.Width
	}
	return m, nil
}

func (m model) View() string {

	var buttons []button

	for name, value := range palette {
		if name == "base" {
			continue
		}
		btn := lipgloss.NewStyle().
			Foreground(lipgloss.Color(palette["base"])).
			Background(lipgloss.Color(value)).
			Padding(1, 2).
			Margin(1, 0, 0, 1).
			Render(capitalise(name))
		w := lipgloss.Width(btn)
		buttons = append(buttons, button{content: btn, width: w})
	}

	var output []string

	occupied := 0
	var row []button
	for _, btn := range buttons {
		if occupied+btn.width < m.width {
			occupied += btn.width
			row = append(row, btn)
		} else {
			var textRows []string
			for _, r := range row {
				textRows = append(textRows, r.content)
			}
			output = append(output, lipgloss.JoinHorizontal(lipgloss.Top, textRows...)) 
			row = nil
			occupied = 0
		}
	}

	if len(row) > 0 {
		var textRows []string
		for _, r := range row {
			textRows = append(textRows, r.content)
		}
		output = append(output, lipgloss.JoinHorizontal(lipgloss.Top, textRows...)) 
	}

	return lipgloss.JoinVertical(lipgloss.Left, output...)
}

func main() {
	p := tea.NewProgram(model{}, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Oops!")
		os.Exit(1)
	}

}

func capitalise(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}
