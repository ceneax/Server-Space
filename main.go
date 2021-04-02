package main

import (
	//"flag"

	//"xcdh.space/space/config"
	"xcdh.space/space/router"
)

func main() {
	// @title 星辰大海 API
	// @version 1.0
	// @description 星辰大海后端API
	// @termsOfService http://swagger.io/terms/

	// @contact.name API支持者
	// @contact.url http://www.swagger.io/support
	// @contact.email revolveyao@qq.com

	// @license.name Apache 2.0
	// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

	// @host localhost:8080
	// @BasePath /
	// @securityDefinitions.apikey ApiKeyAuth
	// @in header
	// @name Authorization
	router.Run()
}
