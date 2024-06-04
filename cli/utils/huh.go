package utils

import "github.com/charmbracelet/huh"

func QueryTarget(value *string) {
	huh.NewInput().Title("Enter target").Prompt("Target: ").Value(value).Run()
}

func QueryUser(value *string) {
	huh.NewInput().Title("Enter user").Prompt("User: ").Value(value).Run()
}

func QueryConfig() *Config {
	var target string
	QueryTarget(&target)
	var user string
	QueryUser(&user)
	config := Config{
		User:   user,
		Target: target,
	}
	WriteConfig(&config)
	return &config
}
