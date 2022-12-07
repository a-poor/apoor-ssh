package main

import (
	"apoor-ssh/pkg/model"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	m := model.NewModel()
	app := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := app.Run(); err != nil {
		panic(err)
	}
}
