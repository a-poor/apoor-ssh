package sections

import tea "github.com/charmbracelet/bubbletea"

type GoToSectionMsg int

func (m GoToSectionMsg) AsCmd() func() tea.Msg {
	return func() tea.Msg {
		return m
	}
}
