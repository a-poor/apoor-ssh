package main

import (
	"apoor-ssh/pkg/sections"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	m := sections.NewModel()
	app := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := app.Run(); err != nil {
		panic(err)
	}
}
