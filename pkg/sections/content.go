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
	Current  int
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
	*c.VP, _ = c.VP.Update(msg)
	return c, nil
}

func (c *Content) RenderSection(i int) string {
	return c.Sections[i].Render(c.Width, 0)
}

func (c *Content) GetSectionHeight(i int) int {
	return lipgloss.Height(c.RenderSection(i))
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

func (c *Content) SetCurrent(i int) {
	if i >= 0 && i < len(c.Sections) {
		c.Current = i
	}
}

func (c *Content) GetCurrent() int {
	return c.Current
}

func (c *Content) SetWidth(w int) {
	c.Width = w
}

func (c *Content) SetHeight(h int) {
	c.Height = h
}
