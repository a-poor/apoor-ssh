package letters

import "github.com/charmbracelet/lipgloss"

func getLetter(letter string) string {
	switch letter {
	case "A", "a":
		return A
	case "B", "b":
		return B
	case "C", "c":
		return C
	case "D", "d":
		return D
	case "E", "e":
		return E
	case "F", "f":
		return F
	case "G", "g":
		return G
	case "H", "h":
		return H
	case "I", "i":
		return I
	case "J", "j":
		return J
	case "K", "k":
		return K
	case "L", "l":
		return L
	case "M", "m":
		return M
	case "N", "n":
		return N
	case "O", "o":
		return O
	case "P", "p":
		return P
	case "Q", "q":
		return Q
	case "R", "r":
		return R
	case "S", "s":
		return S
	case "T", "t":
		return T
	case "U", "u":
		return U
	case "V", "v":
		return V
	case "W", "w":
		return W
	case "X", "x":
		return X
	case "Y", "y":
		return Y
	case "Z", "z":
		return Z
	case " ":
		return Space
	}
	return ""
}

func FormatWord(text string) string {
	chars := []rune(text)
	var letters []string
	for i, letter := range chars {
		letters = append(letters, getLetter(string(letter)))
		if i < len(chars)-1 {
			letters = append(letters, Sep)
		}
	}
	return lipgloss.JoinHorizontal(lipgloss.Top, letters...)
}
