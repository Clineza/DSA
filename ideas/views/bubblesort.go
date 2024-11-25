package views

import (
	tea "github.com/charmbracelet/bubbletea"
)

type BubbleSort struct {
	cursor   int
	selected map[int]struct{}
}

func NewBubbleSort() BubbleSort {
	return BubbleSort{
		selected: make(map[int]struct{}),
	}
}

func (bs BubbleSort) Init() tea.Cmd {
	return nil
}

func (bs BubbleSort) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return bs, tea.Quit
		}
	}

	return bs, nil
}

func (bs BubbleSort) View() string {
	// The header
	s := "\n\n" + gray.Render("which one do you fancy?")

	s += "\n\nThis is the BubbleSort View"

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
