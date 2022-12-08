package model

import "github.com/charmbracelet/lipgloss"

const pageTitle = "$ ssh ssh.austinpoor.com"

func fmtNav(w int, title, user string) string {
	sLeft := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ffffff")).
		Background(lipgloss.Color("#9621a3")).
		Padding(0, 1).
		Render(title)

	var sRight string
	if user != "" {
		sRight = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#ffffff")).
			Background(lipgloss.Color("#cf112d")).
			Padding(0, 1).
			Render(user)
	}

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

func fmtTitle(t string) string {
	return lipgloss.NewStyle().
		Margin(1, 1, 1, 1).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#ff5eb9")).
		Padding(1, 2).
		Bold(true).
		Render(t)
}

func fmtHeader(h string) string {
	return lipgloss.NewStyle().
		Bold(true).
		Border(lipgloss.NormalBorder(), false, false, true, false).
		Render(h)
}

func fmtContactKey(s string) string {
	return lipgloss.NewStyle().
		Background(lipgloss.Color("#b60bd9")).
		MarginRight(1).
		MarginBottom(1).
		Bold(true).
		Render("" + s + ":")
}

func fmtContactVal(s string) string {
	return lipgloss.NewStyle().
		Underline(true).
		Foreground(lipgloss.Color("#5eff89")).
		MarginBottom(1).
		Render(s)
}

func fmtContactKs(ks []string) string {
	strs := make([]string, len(ks))
	for i, k := range ks {
		strs[i] = fmtContactKey(k)
	}
	return lipgloss.JoinVertical(
		lipgloss.Right,
		strs...,
	)
}

func fmtContactVs(vs []string) string {
	strs := make([]string, len(vs))
	for i, v := range vs {
		strs[i] = fmtContactVal(v)
	}
	return lipgloss.JoinVertical(
		lipgloss.Left,
		strs...,
	)
}

func fmtContactKVs(kvs []struct{ k, v string }) string {
	ks := make([]string, len(kvs))
	vs := make([]string, len(kvs))
	for i, kv := range kvs {
		ks[i] = kv.k
		vs[i] = kv.v
	}
	return lipgloss.NewStyle().
		Margin(0, 0).
		Render(
			lipgloss.JoinHorizontal(
				lipgloss.Top,
				fmtContactKs(ks),
				fmtContactVs(vs),
			),
		)
}

func indent(n int, s string) string {
	return lipgloss.NewStyle().
		MarginLeft(n).
		Render(s)
}

func fmtP(p string, w int) string {
	return lipgloss.NewStyle().
		Width(w).
		MarginBottom(1).
		Render("%> " + p)
}

func fmtLink(link string) string {
	return lipgloss.NewStyle().
		Underline(true).
		Foreground(lipgloss.Color("#5e89ff")).
		Render(link)
}
