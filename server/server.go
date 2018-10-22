// Package server baana-app API.
//
// the purpose of this application is to provide APIs for
//
//
//     Schemes: http, https
//     Host: localhost
//     BasePath: /v1/api
//     Version: 0.0.1
//     Contact:
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package server

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
	bsvc "gitlab.com/ajithnn/baana/service"
)

// Run .
func Run(svc *bsvc.Service) {

	defer svc.Terminate()

	// Reading routes.json  to generate routes for Gin
	routes, err := ioutil.ReadFile("config/routes.json")
	if err != nil {
		panic("Invalid route/ no routes.json  found.")
		// os.Exit(1)
	}

	// Uncomment when releasing to production.
	//gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.Static("/public", "./public")
	svc.LoadRoutes(r, routes)
	r.Run(":8080")

}
