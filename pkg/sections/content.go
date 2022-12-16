package sections

import (
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	defaultTerminalWidth  = 24
	defaultTerminalHeight = 80
)

type Content struct {
	Width    int
	Height   int
	Sections []Renderer
	VP       *viewport.Model
	Active   bool
}

func NewContent(sections []Renderer) *Content {
	w, h := defaultTerminalWidth-50, defaultTerminalHeight-10
	vp := viewport.New(w, h)
	return &Content{
		Width:    w,
		Height:   h,
		Sections: sections,
		VP:       &vp,
		Active:   true,
	}
}

func (c *Content) Update(msg tea.Msg) (*Content, tea.Cmd) {
	// Handle the message...
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "right":
			if !c.Active {
				c.Active = true
			}
		case "left", "esc":
			if c.Active {
				c.Active = false
			}
		default:
			if c.Active {
				*c.VP, _ = c.VP.Update(msg)
			}
		}

	case GoToSectionMsg:
		i := int(msg)
		var off int
		for j, s := range c.Sections {
			if j == i {
				break
			}
			txt := s.Render(c.Width, 0)
			off += lipgloss.Height(txt)
		}
		c.VP.SetYOffset(off)
	}

	// Update the viewport...
	if c.Active {
		vp, err := c.VP.Update(msg)
		if err != nil {
			return c, err
		}
		c.VP = &vp
	}

	return c, nil
}

func (c *Content) View() string {
	// Format the body for the viewport...
	var body string
	for _, s := range c.Sections {
		body += s.Render(c.Width, 0)
	}

	padH := c.VP.Height - 1
	c.VP.SetContent(body + strings.Repeat("\n", padH))
	c.VP.Height = c.Height - 2

	// Render and return...
	return lipgloss.NewStyle().
		MaxWidth(c.Width).
		Width(c.Width - 2).
		PaddingLeft(2).
		MaxHeight(c.Height).
		Render(c.VP.View())
}

func (c *Content) SetWidth(w int) {
	c.Width = w
}

func (c *Content) SetHeight(h int) {
	c.Height = h
}
