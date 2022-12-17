package sections

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	Width     int
	Height    int
	Nav       *Nav
	Content   *Content
	NavActive bool
	UserName  string
}

func NewModel() *Model {
	// Create the sub-components...
	nav := NewNav(navTitles)
	content := NewContent([]Renderer{
		helloSection,
		aboutMeSection,
		// projectsSection,
		blogSection,
		contactMeSection,
		// resumeSection,
		thisSiteSection,
	})

	// Return the model
	return &Model{
		Width:   10,
		Height:  30,
		Nav:     nav,
		Content: content,
	}
}

func (m *Model) WithUserName(name string) *Model {
	m.UserName = name
	return m
}

func (m *Model) WithWidth(width int) *Model {
	m.Width = width
	return m
}

func (m *Model) WithHeight(height int) *Model {
	m.Height = height
	return m
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) UpdateContentPos() {
	// Get the current nav section...
	section := m.Nav.Current

	// Get the height of each section of content...
	var total int
	for i := 0; i < section; i++ {
		total += m.Content.GetSectionHeight(i)
	}

	// Set the content viewport to the current section...
	m.Content.VP.YOffset = total - section
	if m.Content.VP.YOffset < 0 {
		m.Content.VP.YOffset = 0
	}
}

func (m *Model) UpdateNavPos() {
	// Get the current content viewport height...
	yo := m.Content.VP.YOffset

	// Get the height of each section of content...
	var total int
	ch := len(m.Content.Sections) - 1
	for i := 0; i < len(m.Content.Sections); i++ {
		h := m.Content.GetSectionHeight(i)
		total += h - 1
		offset_diff := 1
		if total-offset_diff >= yo {
			ch = i
			break
		}
	}

	// Set the nav current section to the content section...
	m.Nav.SetCurrent(ch)
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Handle the message...
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "right":
			if m.NavActive {
				m.NavActive = false
				m.Nav.Active = false
				m.Content.Active = true
			}
		case "left", "esc":
			if !m.NavActive {
				m.NavActive = true
				m.Nav.Active = true
				m.Content.Active = false
			}
		default:
			if m.NavActive {
				m.Nav, _ = m.Nav.Update(msg)
				m.UpdateContentPos()
			} else {
				m.Content, _ = m.Content.Update(msg)
				m.UpdateNavPos()
			}
		}

		// // Pass the message to the nav...
		// m.Nav, _ = m.Nav.Update(msg)

		// // Pass the message to the content...
		// m.Content, _ = m.Content.Update(msg)

	case tea.WindowSizeMsg:
		// Store the new window size
		m.Width = msg.Width
		m.Height = msg.Height

		// Update the nav's max-height...
		m.Nav.SetHeight(m.Height - 1)

		// Update the content's width...
		nw := lipgloss.Width(m.Nav.View())
		m.Content.SetWidth(m.Width - nw)
		m.Content.SetHeight(m.Height - 1)
	}

	// Keep going...
	return m, nil
}

func (m *Model) View() string {
	body := lipgloss.JoinHorizontal(
		lipgloss.Top,
		m.Nav.View(),
		m.Content.View(),
	)
	body = lipgloss.NewStyle().
		Height(m.Height - 1).
		MaxHeight(m.Height - 1).
		Width(m.Width).
		Render(body)

	footer := FormatFooter(m.Width, PageTitle, m.UserName)
	return lipgloss.JoinVertical(
		lipgloss.Left,
		body,
		footer,
	)
}
