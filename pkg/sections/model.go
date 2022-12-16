package sections

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ModelConf struct{}

type Model struct {
	Width   int
	Height  int
	Nav     *Nav
	Content *Content
}

func NewModel() *Model {
	return &Model{
		Width:  10,
		Height: 30,
		Nav:    NewNav(navTitles),
		Content: NewContent([]Renderer{
			helloSection,
			aboutMeSection,
			// projectsSection,
			blogSection,
			contactMeSection,
			// resumeSection,
			thisSiteSection,
		}),
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Handle the message...
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}

		// Pass the message to the nav...
		m.Nav, _ = m.Nav.Update(msg)

		// Pass the message to the content...
		m.Content, _ = m.Content.Update(msg)

	case tea.WindowSizeMsg:
		// Store the new window size
		m.Width = msg.Width
		m.Height = msg.Height

		// Update the nav's max-height...
		m.Nav.SetHeight(m.Height)

		// Update the content's width...
		nw := lipgloss.Width(m.Nav.View())
		m.Content.SetWidth(m.Width - nw)
		m.Content.SetHeight(m.Height)
	}

	// Keep going...
	return m, nil
}

func (m *Model) View() string {
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		m.Nav.View(),
		// "Hello,\n...\n...world?",
		m.Content.View(),
	)
}
