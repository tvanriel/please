package app

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const COLOUR_PRIMARY = "#4c9be8"
const COLOUR_DANGER = "#d9534f"
const COLOUR_WARNING = "#ffc107"
const COLOUR_SUCCESS = "#5cb85c"

var titleStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(COLOUR_PRIMARY)).Bold(true).PaddingTop(1).PaddingLeft(2).PaddingBottom(1).PaddingRight(0)
var textStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("white")).Bold(false)
var boldStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("white")).Bold(true)
var successStyle = textStyle.Copy().Foreground(lipgloss.Color(COLOUR_SUCCESS))
var dangerStyle = textStyle.Copy().Foreground(lipgloss.Color(COLOUR_DANGER))
var warningStyle = textStyle.Copy().Foreground(lipgloss.Color(COLOUR_WARNING))

var listStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(1)

func renderLine(i int, text string) string {
	return fmt.Sprintf(
		"%s. %s",
		boldStyle.Render(strconv.Itoa(i+1)),
		textStyle.Render(text),
	)
}

func renderTodoList(list []string) {
	formattedLines := make([]string, len(list))

	for i := range list {
		if len(list[i]) > 0 {
			formattedLines[i] = renderLine(i, list[i])
		}
	}

	fmt.Println(listStyle.Render(strings.Join(formattedLines, "\n")))
}
