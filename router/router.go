package Router

import (
	"github.com/kataras/iris/v12"
	"github.com/hfeng101/niwo/router/register"
)

type Routing struct {

}

func RegistryRoutes(app *iris.Application){
	register.LoginRouteRegister(app)
	register.LogoutRouteRegister(app)

	register.UserRouteRegister(app)

	register.SwaggerRouteRegister(app)
}
