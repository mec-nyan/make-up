// main.go
package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct{}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Just wait for the user to press a key, then exit.
	switch msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit
	}
	return m, nil
}

func (m model) View() string {
	return "Hello!"
}

func main() {

	p := tea.NewProgram(model{})
	if _, err := p.Run(); err != nil {
		fmt.Println("Oops!", err)
		os.Exit(1)
	}
}
