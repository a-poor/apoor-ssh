package sections

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	boxActiveColor   = lipgloss.Color("#f0f0f0")
	boxInactiveColor = lipgloss.Color("#303030")

	titleColor = lipgloss.Color("#5eb9ff")

	sectionHeaderFgColor = lipgloss.Color("#f0f0f0")

	tagBgColor = lipgloss.Color("#303030")
	tagFgColor = lipgloss.Color("#f0f0f0")

	linkColor = lipgloss.Color("#5eff89")
)

func FormatTitle(title string) string {
	style := lipgloss.NewStyle().
		Bold(true).
		Margin(1).
		MarginLeft(0).
		Foreground(titleColor).
		Underline(true)
	return style.Render(title)
}

func FormatSectionHeader(h string, w int) string {
	style := lipgloss.NewStyle().
		Bold(true).
		MarginBottom(1).
		Width(w).
		Foreground(sectionHeaderFgColor)
	return style.Render(h)
}

func FormatSectionDivider(w int) string {
	// Create the style...
	style := lipgloss.NewStyle().
		Margin(1).
		Align(lipgloss.Center).
		Width(w)

	// How many dashes do we need?
	txt := strings.Repeat("-", 3)

	// Render and return...
	return style.Render(txt)
}

func FormatBody(body string, w int) string {
	style := lipgloss.NewStyle().
		Width(w).
		MaxWidth(w)
	return style.Render(body)
}

func FormatLink(link string) string {
	style := lipgloss.NewStyle().
		Foreground(linkColor)
	return style.Render(link)
}

func FormatKeyword(w string) string {
	style := lipgloss.NewStyle().
		Italic(true).
		Foreground(lipgloss.Color("#5eb9ff"))
	return style.Render(w)

}
