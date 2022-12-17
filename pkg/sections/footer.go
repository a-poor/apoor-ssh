package sections

import "github.com/charmbracelet/lipgloss"

const PageTitle = "$ ssh ssh.austinpoor.com"

func FormatFooter(w int, title, user string) string {
	sLeft := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ffffff")).
		Background(lipgloss.Color("#9621a3")).
		Padding(0, 1).
		Render(title)

	u := user
	if u == "" {
		u = "(anon)"
	}
	sRight := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ffffff")).
		Background(lipgloss.Color("#112dcf")).
		Padding(0, 1).
		Render(u)

	sMiddle := lipgloss.NewStyle().
		Width(w - lipgloss.Width(sLeft) - lipgloss.Width(sRight)).
		Background(lipgloss.Color("#373b41")).
		Render("")

	// Join them together...
	sJoin := lipgloss.JoinHorizontal(
		lipgloss.Top,
		sLeft,
		sMiddle,
		sRight,
	)

	// Clip and return...
	return lipgloss.NewStyle().
		MaxWidth(w).
		Render(sJoin)
}
