package controllers

import "go-chat-app-monolith/internal/pkg/users"

type Controller struct {
	UserService *users.Service
}
