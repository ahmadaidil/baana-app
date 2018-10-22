package controllers

import (
	"github.com/ahmadaidil/baana-app/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"gitlab.com/ajithnn/baana/service"
)

// UserController .
type UserController struct {
	*models.User
	db *gorm.DB
	*service.Query
}

// User .
func User(db *gorm.DB) interface{} {
	handle := &UserController{&models.User{}, db, &service.Query{}}
	return handle
}

// Create -> swagger:operation POST /users createUser
//
// Creates a new User with the given parameters.
//
//
// ---
// produces:
// - application/json
// parameters:
// responses:
func (handler *UserController) Create(c *gin.Context) {
}

// Update -> swagger:operation PUT /users/{id} updateUser
//
// Update a User with the given parameters.
//
//
// ---
// produces:
// - application/json
// parameters:
// responses:
func (handler *UserController) Update(c *gin.Context) {
}

// Read -> swagger:operation GET /users/ getUsers
//
// Get Users with the given query filters.
//
//
// ---
// produces:
// - application/json
// parameters:
// responses:
func (handler *UserController) Read(c *gin.Context) {
}

// Delete -> swagger:operation DELETE /users/{id} deleteUser
//
// Delete Users  with the given query filter.
//
//
// ---
// produces:
// - application/json
// parameters:
// responses:
func (handler *UserController) Delete(c *gin.Context) {
}
