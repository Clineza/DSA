package views

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	red  = lipgloss.NewStyle().Foreground(lipgloss.Color("9"))
	blue = lipgloss.NewStyle().Foreground(lipgloss.Color("12"))
	gray = lipgloss.NewStyle().Foreground(lipgloss.Color("7"))
)

type Menu struct {
	cursor         int
	algorithms     []string
	dataStructures []string
	choices        []string
}

func NewMenu() Menu {
	algorithms := []string{"Bubble Sort", "Selection Sort"}
	dataStructures := []string{"Linked List", "Queue", "Stack"}
	choices := append(algorithms, dataStructures...)
	return Menu{
		algorithms:     algorithms,
		dataStructures: dataStructures,
		choices:        choices,
	}
}

func (m Menu) Init() tea.Cmd {
	return nil
}

func (m Menu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			switch m.choices[m.cursor] {
			case "Bubble Sort":
				return NewBubbleSort(), nil
			}
		}
	}

	return m, nil
}

func (m Menu) View() string {
	// The header
	selectedColor := gray
	isSelected := false
	s := "\n\n" + gray.Render("which one do you fancy?")
	s += "\n"

	for i, choice := range m.choices {

		if i == 0 {
			s += fmt.Sprintf("\n" + gray.Render("an ") + red.Render("algorithm") + gray.Render("...") + "\n\n")
		} else if i == len(m.algorithms) {
			s += "\nor a " + blue.Render("Data Structure") + "...\n\n"
		}

		cursor := " "
		checked := " "
		if i < len(m.algorithms) {
			if m.cursor == i {
				isSelected = true
				cursor = ">"
				checked = "a"
			} else {
				isSelected = false
			}
			if isSelected == true {
				selectedColor = red
			} else {
				selectedColor = gray
			}
		} else {
			if m.cursor == i {
				isSelected = true
				cursor = ">"
				checked = "d"
			} else {
				isSelected = false
			}
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

/*
one option for routing would be to maybe have a function that is called for an enter event.
so if the user clicks enter on a selection.. then a function is called that returns a view to the model.

only thing is i would have to find some way for the main.go to know the return value of a function being called somewhere else

*/
