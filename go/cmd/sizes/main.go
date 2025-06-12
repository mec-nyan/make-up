// main.go
package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	_ = iota
	fullSize
	topBottom
	leftRight
	quarters
)

type model struct {
	width, height, what int
}

var boxStyle = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder()).
	Padding(1, 4)

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "j", "n":
			if m.what < quarters {
				m.what++
			}
		case "k", "p":
			if m.what > 0 {
				m.what--
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	var output string
	switch m.what {
	case fullSize:
		innerHeigt := m.height - boxStyle.GetVerticalFrameSize() + boxStyle.GetPaddingTop() + boxStyle.GetPaddingBottom()
		innerWidth := m.width - boxStyle.GetHorizontalFrameSize() + boxStyle.GetPaddingLeft() + boxStyle.GetPaddingRight()
		fullBoxStyle := boxStyle.Width(innerWidth).Height(innerHeigt)
		output = fmt.Sprintf("rows: %d\ncols: %d\n", m.height, m.width)
		output += fmt.Sprintf("- padding: %d %d %d %d\n", fullBoxStyle.GetPaddingTop(),
			fullBoxStyle.GetPaddingRight(), fullBoxStyle.GetPaddingBottom(), fullBoxStyle.GetPaddingLeft())
		output += fmt.Sprintf("- margin:  %d %d %d %d\n", fullBoxStyle.GetMarginTop(),
			fullBoxStyle.GetMarginRight(), fullBoxStyle.GetMarginBottom(), fullBoxStyle.GetMarginLeft())
		output += fmt.Sprintf("- inner size: rows %d cols %d", innerHeigt, innerWidth)
		return fullBoxStyle.Render(output)
	case topBottom:
	case leftRight:
	case quarters:
	default:
		output =  fmt.Sprintf("Cols: %d\nRows: %d\n", m.width, m.height)
	}
	return output
}

func main() {

	p := tea.NewProgram(model{}, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Oops!", err)
		os.Exit(1)
	}
}
