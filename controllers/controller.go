package controllers

import (
	"github.com/gorilla/mux"
	"github.com/sflewis2970/accounts-api/handlers"
	"log"
)

// Controller structure defines teh layout of the Controller
type Controller struct {
	Router      *mux.Router
	acctHandler *handlers.AccountHandler
}

// Package controllers object
var controller *Controller

func (c *Controller) setupRoutes() {
	// Display log message
	log.Print("Setting up accounts api service routes")

	// Trivia routes
	c.Router.HandleFunc("/api/v1/api/register", c.acctHandler.RegisterUser).Methods("POST")
	c.Router.HandleFunc("/api/v1/api/find", c.acctHandler.GetUser).Methods("GET")
	c.Router.HandleFunc("/api/v1/api/update", c.acctHandler.UpdateUser).Methods("PUT")
	c.Router.HandleFunc("/api/v1/api/remove", c.acctHandler.DeleteUser).Methods("DELETE")
}

// NewController function create a new Controller and initializes new Controller object
func NewController() *Controller {
	// Create controllers component
	log.Print("Creating controllers object...")
	controller = new(Controller)

	// Trivia handler
	controller.acctHandler = handlers.NewAccountHandler()

	// Set controllers routes
	controller.Router = mux.NewRouter()
	controller.setupRoutes()

	return controller
}
