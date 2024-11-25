package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"main/views"
	"os"
)

type model struct {
	currentView tea.Model
}

func initialModel() model {
	return model{
		currentView: views.NewMenu(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	model, cmd := m.currentView.Update(msg)
	return model, cmd
}

func (m model) View() string {
	view := m.currentView.View()
	return view
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

/*
	the code works but i need to split the program into different views
1. Start View.
2. algorithm view
3. data structure view

each data structure and algorithm will be its own view
each view should have everything necessary to select (controls) and to view the text or interact.

so pretty much main.go should simply be for routing reasons and should hold the states of the views along with the data for the views themselves.





*/
