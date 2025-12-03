package main

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	Quit    key.Binding
	Help    key.Binding
	Execute key.Binding
	Up      key.Binding
	Down    key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Quit, k.Help, k.Execute, k.Up, k.Down}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Quit, k.Help, k.Execute},
		{k.Up, k.Down},
	}
}

var keys = keyMap{
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
	Help: key.NewBinding(
		key.WithKeys("h", "?"),
		key.WithHelp("h/?", "help"),
	),
	Execute: key.NewBinding(
		key.WithKeys("enter", "space", "return"),
		key.WithHelp("enter", "execute day"),
	),
	Up: key.NewBinding(
		key.WithKeys("k", "up"),
		key.WithHelp("k/up", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("j", "down"),
		key.WithHelp("j/down", "move down"),
	),
}
