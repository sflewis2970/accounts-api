package server

import (
	"github.com/rs/cors"
	"github.com/sflewis2970/accounts-api/config"
	"github.com/sflewis2970/accounts-api/controllers"
	"log"
	"net/http"
)

func main() {
	// Initialize logging
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// Get config data
	cfgData := config.NewConfig().LoadCfgData()

	// Create controllers
	controller := controllers.NewController()

	// setup Cors
	log.Print("Setting up CORS...")
	corsOptionsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodPost, http.MethodGet},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})
	corsHandler := corsOptionsHandler.Handler(controller.Router)

	// Server Address info
	addr := cfgData.Host + ":" + cfgData.Port
	log.Print("The address used by the service is: ", addr)

	// Start Server
	log.Print("Web service server is ready...")

	// Listen and Serve
	log.Fatal(http.ListenAndServe(addr, corsHandler))
}
