package config

// App should be used to store common information about the Web Application.
type App struct {
	Description string
	Name        string
}

// DefaultApp specifies the concrete information.
var DefaultApp = &App{
	Description: "Put your description here",
	Name:        "My First App",
}
