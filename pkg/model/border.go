package model

import "github.com/charmbracelet/lipgloss"

var shadowBorder = lipgloss.Border{
	Top:         "",
	Bottom:      "▀",
	Left:        "",
	Right:       "█",
	TopLeft:     "",
	TopRight:    "", // "◣",
	BottomLeft:  "", // "◥",
	BottomRight: "▀",
}
