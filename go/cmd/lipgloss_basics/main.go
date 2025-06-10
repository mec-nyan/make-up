package main

import (
	"fmt"
	"strings"

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

func main() {

	var buttons []string

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
		buttons = append(buttons, btn)
	}

	fmt.Println(lipgloss.JoinHorizontal(lipgloss.Top, buttons[:len(buttons)/2]...))
	fmt.Println(lipgloss.JoinHorizontal(lipgloss.Top, buttons[len(buttons)/2:]...) + "\n")
}

func capitalise(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}
