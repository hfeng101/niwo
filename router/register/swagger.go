package register

import (
	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"

	_ "github.com/hfeng101/niwo/docs"
)

//swgger文档入口
func SwaggerRouteRegister(app *iris.Application) {
	config := &swagger.Config{
		//or URL is http://localhost:8080/swagger/index.html
		URL: "http://localhost:8080/swagger/doc.json", //The url pointing to API definition
	}
	// use swagger middleware to
	app.Get("/swagger/{any:path}", swagger.CustomWrapHandler(config, swaggerFiles.Handler))
}
