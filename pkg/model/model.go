package model

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	w, h int
}

func NewModel() *Model {
	return &Model{}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.w, m.h = msg.Width, msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m *Model) View() string {
	// Format the title
	title := lipgloss.NewStyle().
		MarginBottom(1).
		Padding(1, 2).
		Bold(true).
		// Foreground(lipgloss.Color("12")).
		Background(lipgloss.Color("#9621a3")).
		Border(shadowBorder).
		BorderForeground(lipgloss.Color("#eb7cf7")).
		Render("Hi, I'm Austin!")

	// Make the paragraphs...
	p1 := lipgloss.NewStyle().
		Width(m.w - 5).
		MarginBottom(1).
		Render("> I'm a software engineer living in Los Angeles, CA.")

	p2 := lipgloss.NewStyle().
		Width(m.w - 5).
		MarginBottom(1).
		Render(`> I like to work with data and build tools to help people.`)

	// Contact info styles...
	cks := lipgloss.NewStyle().
		MarginRight(2).
		Italic(true)

	cvs := lipgloss.NewStyle().
		Underline(true)

	// Contact info...
	contacts := []struct {
		key, val string
	}{
		{
			key: "Website:",
			val: "http://austinpoor.com",
		},
		{
			key: "GitHub:",
			val: "https://github.com/a-poor",
		},
		{
			key: "LinkedIn:",
			val: "http://linkedin.com/in/austin_poor",
		},
	}
	contactKeys := mapFn(contacts, func(c struct{ key, val string }) string {
		return cks.Render(c.key)
	})
	contactVals := mapFn(contacts, func(c struct{ key, val string }) string {
		return cvs.Render(c.val)
	})

	joinedKeys := lipgloss.JoinVertical(lipgloss.Right, contactKeys...)
	joinedVals := lipgloss.JoinVertical(lipgloss.Left, contactVals...)
	joinedContact := lipgloss.JoinHorizontal(lipgloss.Top, joinedKeys, joinedVals)

	contactTitle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, false, true, false).
		BorderForeground(lipgloss.Color("#404040")).
		Render(`Want to learn more?`)

	// Join the content...
	content := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		p1,
		p2,
		contactTitle,
		joinedContact,
	)

	// ...and render it with a margin
	return lipgloss.NewStyle().
		MaxHeight(m.h).
		MaxWidth(m.w).
		Margin(1).
		Render(content)
}

func mapFn[T1 any, T2 any](d []T1, fn func(T1) T2) []T2 {
	r := make([]T2, len(d))
	for i, v := range d {
		r[i] = fn(v)
	}
	return r
}
