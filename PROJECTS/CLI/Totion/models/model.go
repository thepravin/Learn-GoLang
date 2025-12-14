package models

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	MSG string
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// It is a key press?
	case tea.KeyMsg:

		// what was the actual key press
		switch msg.String() {

		// These keys should exit the program
		case "ctrl+c", "q":
			return m, tea.Quit

		}

	}

	return m, nil
}

func (m Model) View() string {
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("16")).
		Background(lipgloss.Color("205")).
		PaddingLeft(2).
		PaddingRight(2)

	welcome := style.Render("WELCOME TO TOTION 📝")

	help := "Ctrl+N: new file | Ctrl+L: list | Esc: back/save | Ctrl+Q: quit"

	view := ""

	return fmt.Sprintf("\n%s\n\n%s\n\n%s", welcome, view, help)
}

func NewMessage(initialMsg string) Model {
	return Model{
		MSG: initialMsg,
	}
}
