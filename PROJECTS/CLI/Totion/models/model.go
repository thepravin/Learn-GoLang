package models

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	newFileInput             textinput.Model
	notetextarea             textarea.Model
	isCreateFileInputVisible bool
	RootDir                  string
	currentFile              *os.File // sotre the file pointer
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd

	switch msg := msg.(type) {

	// It is a key press?
	case tea.KeyMsg:

		// what was the actual key press
		switch msg.String() {

		// These keys should exit the program
		case "ctrl+c", "q":
			return m, tea.Quit
		case "ctrl+n":
			// fmt.Println("Key : ", msg)
			m.isCreateFileInputVisible = true
			return m, nil

		case "ctrl+s", "esc":
			// textarea value -> write it in that file and close it

			if m.currentFile == nil { // If NO file is open, do nothing
				break
			}

			if err := m.currentFile.Truncate(0); err != nil {
				fmt.Println("Can not save the file :(")
				return m, nil
			}

			if _, err := m.currentFile.Seek(0, 0); err != nil {
				fmt.Println("Can not save the file :(")
				return m, nil
			}

			// write into file
			if _, err := m.currentFile.WriteString(m.notetextarea.Value()); err != nil {
				fmt.Println("Can not save the file :(")
				return m, nil
			}

			if err := m.currentFile.Close(); err != nil {
				fmt.Println("Can not close the file :(")
				return m, nil
			}

			m.currentFile = nil
			m.notetextarea.SetValue("")

			return m, nil

		case "enter":
			// todo : create file
			fileName := m.newFileInput.Value()
			if fileName != "" {
				filePath := fmt.Sprintf("%s/%s.md", m.RootDir, fileName)

				// check file already prsent or not
				if _, err := os.Stat(filePath); err == nil {
					// file exist
					return m, nil
				}

				file, err := os.Create(filePath)
				if err != nil {
					log.Fatal("Error : %v", err)
				}
				m.currentFile = file
				m.isCreateFileInputVisible = false
				m.newFileInput.SetValue("")
			}
			return m, nil
		}

	}

	if m.isCreateFileInputVisible {
		m.newFileInput, cmd = m.newFileInput.Update(msg)
	}

	if m.currentFile != nil {
		m.notetextarea, cmd = m.notetextarea.Update(msg)
	}

	return m, cmd
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
	if m.isCreateFileInputVisible {
		view = m.newFileInput.View()
	}

	if m.currentFile != nil {
		view = m.notetextarea.View()
	}

	return fmt.Sprintf("\n%s\n\n%s\n\n%s", welcome, view, help)
}

// helper function
func ModelInitializationBridge(fileName textinput.Model, isCreateFileInputVisible bool, rootDir string, notesTextArea textarea.Model) Model {
	return Model{
		newFileInput:             fileName,
		isCreateFileInputVisible: isCreateFileInputVisible,
		RootDir:                  rootDir,
		notetextarea:             notesTextArea,
	}
}
