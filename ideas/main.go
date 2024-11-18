package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"os"
)

var (
	red  = lipgloss.NewStyle().Foreground(lipgloss.Color("9"))
	blue = lipgloss.NewStyle().Foreground(lipgloss.Color("12"))
	gray = lipgloss.NewStyle().Foreground(lipgloss.Color("7"))
)

type model struct {
	cursor         int
	choices        []string
	algorithms     []string
	dataStructures []string
	selected       map[int]struct{}
}

func initialModel() model {
	return model{}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	// The header
	selectedColor := gray
	isSelected := false
	s := "\n\n" + gray.Render("which one do you fancy?")
	s += "\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			isSelected = true
			cursor = ">"
		} else {
			isSelected = false
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		if i == 0 {
			s += fmt.Sprintf("\n" + gray.Render("an ") + red.Render("algorithm") + gray.Render("...") + "\n\n")
		} else if i == len(m.algorithms) {
			s += "\nor a " + blue.Render("Data Structure") + "...\n\n"
		}

		if i < len(m.algorithms) {
			if isSelected == true {
				selectedColor = red
			} else {
				selectedColor = gray
			}
		} else {
			if isSelected == true {
				selectedColor = blue
			} else {
				selectedColor = gray
			}
		}

		s += fmt.Sprintf(selectedColor.Render("%s [%s] %s"), cursor, checked, choice)
		s += "\n"
	}

	s += "\n"

	// The footer
	s += "\n" + gray.Render("Press q to quit.") + "\n"

	// Send the UI for rendering
	return s
}

func main() {
	algorithms := []string{"Bubble Sort", "Selection Sort"}
	dataStructures := []string{"Linked List", "Queue", "Stack"}
	choices := append(algorithms, dataStructures...)
	m := model{
		algorithms:     algorithms,
		dataStructures: dataStructures,
		choices:        choices,
		selected:       make(map[int]struct{}),
	}

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

/*

so i need an array of projectiles.
each projectile will have a value equal to what's the base

*/
