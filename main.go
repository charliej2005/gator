package main

import (
	"fmt"

	"github.com/charliej2005/gator/internal/config"
)

func main() {
	// placeholder main that accesses the config and updates it to my name :)
	cfg, _ := config.Read()
	fmt.Printf("dbUrl: %v, username: %v\n", cfg.DbURL, cfg.CurrentUserName)
	username := "charlie"
	fmt.Println("setting username")
	_ = cfg.SetUser(username)
	cfg, _ = config.Read()
	fmt.Printf("dbUrl: %v, username: %v\n", cfg.DbURL, cfg.CurrentUserName)
	// implement proper functionality later
}
