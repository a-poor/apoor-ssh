package sections

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Nav struct {
	Width    int      // Box (max) width
	Height   int      // Box max height
	Current  int      // Index of the current section
	Sections []string // Section names
	Active   bool     // Is the nav section active?
}

func NewNav(sections []string) *Nav {
	return &Nav{
		Width:    20,
		Height:   10,
		Current:  0,
		Sections: sections,
		Active:   false,
	}
}

func (n *Nav) SetWidth(w int) {
	n.Width = w
}

func (n *Nav) SetHeight(h int) {
	n.Height = h
}

func (n *Nav) SetCurrent(i int) {
	if i >= 0 && i < len(n.Sections) {
		n.Current = i
	}
}

func (n *Nav) SetActive(a bool) {
	n.Active = a
}

func (n *Nav) Update(msg tea.Msg) (*Nav, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if n.Current > 0 {
				n.Current--
			}
		case "down", "j":
			if n.Current < len(n.Sections)-1 {
				n.Current++
			}
		}
	}

	return n, nil
}

func (n *Nav) View() string {
	// Create the border...
	boxBorderColor := boxInactiveColor
	if n.Active {
		boxBorderColor = boxActiveColor
	}

	// Define the styles...
	boxStyle := lipgloss.NewStyle().
		Width(n.Width).
		MaxHeight(n.Height).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(boxBorderColor).
		Padding(1, 1)
	deselectedStyle := lipgloss.NewStyle().
		Faint(true)
	selectedStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#9621a3")).
		Bold(true)

	// Get the formatted text...
	var text []string
	for i, s := range n.Sections {
		if i == n.Current {
			text = append(
				text,
				selectedStyle.Render(
					"> "+s+" <",
				),
			)
		} else {
			text = append(
				text,
				deselectedStyle.Render(
					"  "+s+" ",
				),
			)
		}
	}

	// Join the text...
	joined := lipgloss.JoinVertical(lipgloss.Top, text...)

	// Return the data...
	return boxStyle.Render(joined)
}
