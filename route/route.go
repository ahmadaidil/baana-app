package route

import (
	"github.com/ahmadaidil/baana-app/controllers"
	"gitlab.com/ajithnn/baana/service"
)

// Init .
func Init() {
	service.ControllerFuncs = make(map[string]service.HandlerInit)

	service.ControllerFuncs["User"] = controllers.User

}
