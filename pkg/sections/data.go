package sections

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var greeting = `# Hi there! I'm Austin!`

var navTitles = []string{
	"Hello!",
	"About Me",
	"Projects",
	"Blog",
	"Contact",
	"Resume",
	"This Site",
}

var helloSection = RenderFunc(func(w, h int) string {
	title := FormatTitle(greeting)
	div := FormatSectionDivider(w)
	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		div,
	)
})

var aboutMeSection = RenderFunc(func(w, h int) string {
	div := FormatSectionDivider(w)
	header := FormatSectionHeader("## 👋  About Me", w)
	rawBody := `Hi, I'm Austin! I'm a software engineer living in Los Angeles, CA. 

I really enjoy working with data, tackling difficult problems, and designing interfaces to communicate complex ideas.

Some of my favorite technologies and areas-of-interest include Go, TypeScript, Rust, SQL, gRPC, and ML.`
	body := FormatBody(rawBody, w)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		body,
		div,
	)
})

var projectsSection = RenderFunc(func(w, h int) string {
	div := FormatSectionDivider(w)
	header := FormatSectionHeader("## 🚀  Projects", w)
	rawBody := `(Coming soon...)`
	body := FormatBody(rawBody, w)
	return lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		body,
		div,
	)
})

var blogSection = RenderFunc(func(w, h int) string {
	div := FormatSectionDivider(w)
	header := FormatSectionHeader("## 👨‍💻  Blog", w)
	rawBody := `I occasionally write some blog posts.

Eventually, I plan to include some info on recent blog posts here.

For now, check out my blog on Medium: %s`
	body := FormatBody(
		fmt.Sprintf(
			rawBody,
			FormatLink("https://medium.com/@apoor"),
		),
		w,
	)
	return lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		body,
		div,
	)
})

var contactMeSection = RenderFunc(func(w, h int) string {
	// The data...
	contactInfo := []struct{ k, v string }{
		{"GitHub", "https://github.com/a-poor"},
		{"LinkedIn", "https://linkedin.com/in/austin_poor"},
		{"Blog", "https://medium.com/@apoor"},
		{"Website", "https://austinpoor.com"},
		{"Mastodon", "https://mastodon.social/@austinpoor"},
	}

	// Formatted content...
	div := FormatSectionDivider(w)
	header := FormatSectionHeader("## 📨  Contact", w)

	// Create the styles...
	keyStyle := lipgloss.NewStyle().
		Italic(true).
		Background(tagBgColor).
		Foreground(tagFgColor).
		Padding(0, 1).
		Margin(0, 1, 1)
	valStyle := lipgloss.NewStyle().
		Underline(true).
		Foreground(linkColor)
	containerStyle := lipgloss.NewStyle().
		MaxWidth(w)

	// Format the body...
	var parts []string
	for _, info := range contactInfo {
		parts = append(parts, lipgloss.JoinHorizontal(
			lipgloss.Left,
			"-"+keyStyle.Render(info.k+":"),
			valStyle.Render(info.v),
		))
	}
	body := containerStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Left,
			parts...,
		),
	)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		body,
		div,
	)
})

var resumeSection = RenderFunc(func(w, h int) string {
	div := FormatSectionDivider(w)
	header := FormatSectionHeader("## 💼  Resume", w)
	rawBody := `(Coming soon...)`
	body := FormatBody(rawBody, w)
	return lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		body,
		div,
	)
})

var thisSiteSection = RenderFunc(func(w, h int) string {
	rawBody := `I wrote this app with Go, Bubble Tea, and Wish.

You can find the source code here: %s`

	header := FormatSectionHeader("## 🛠️  This Site", w)

	body := FormatBody(
		fmt.Sprintf(
			rawBody,
			FormatLink("https://github.com/a-poor/apoor-ssh"),
		),
		w,
	)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		body,
	)
})
