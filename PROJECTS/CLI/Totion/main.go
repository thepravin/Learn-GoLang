package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/thepravin/totion/models"
)

var (
	cursorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
)

func initializeMode() models.Model {
	// Initialize new file input
	ti := textinput.New()
	ti.Placeholder = "What would you like to call it?"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 40
	ti.Cursor.Style = cursorStyle
	ti.PromptStyle = cursorStyle
	ti.TextStyle = cursorStyle

	return models.NewMessage(ti, false)
}

func main() {
	p := tea.NewProgram(initializeMode())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error : ", err.Error())
		os.Exit(1)
	}
}
