package model

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	user string
	w, h int
}

type ModelConf struct {
	User   string
	Width  int
	Height int
}

func NewModel(conf ModelConf) *Model {
	return &Model{
		user: conf.User,
		w:    conf.Width,
		h:    conf.Height,
	}
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
	// Make the paragraphs...
	p1 := "I'm a software engineer living in Los Angeles, CA."

	p2 := "I currently work as a full-stack developer at Command Credit -- a business credit reseller in Rhinebeck, NY."

	p3 := "I enjoy problem solving, working with data (building pipelines, managaing large datasets, and extracting insights), and build tools to help people (including web apps, APIs, CLIs, and more!)."

	p4 := fmt.Sprintf(
		"This SSH server was built with Go, Bubble Tea, and Wish. You can learn more at: %s",
		fmtLink("https://github.com/a-poor/apoor-ssh"),
	)

	// Contact info...
	contacts := []struct {
		k, v string
	}{
		{
			k: "Website",
			v: "http://austinpoor.com",
		},
		{
			k: "GitHub",
			v: "https://github.com/a-poor",
		},
		{
			k: "LinkedIn",
			v: "http://linkedin.com/in/austin_poor",
		},
	}
	joinedContact := fmtContactKVs(contacts)

	// Join the content...
	content := lipgloss.JoinVertical(
		lipgloss.Left,
		fmtNav(m.w, pageTitle, m.user),
		fmtTitle("Hi, I'm Austin!"),
		// "\n",
		indent(1, fmtHeader("About Me")),
		indent(2, fmtP(p1, m.w-5)),
		indent(2, fmtP(p2, m.w-5)),
		indent(2, fmtP(p3, m.w-5)),
		indent(2, fmtP(p4, m.w-5)),
		"\n",
		indent(1, fmtHeader("Want to find out more?")),
		indent(2, joinedContact),
	)

	// ...and render it with a margin
	return lipgloss.NewStyle().
		MaxHeight(m.h).
		MaxWidth(m.w).
		Render(content)
}
