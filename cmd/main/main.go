package main

import "go-chat-app-monolith/internal/app"

func main() {
	appObj := app.NewApp()
	appObj.Run()
}
