package app

import (
	"context"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	docs "github.com/yeyee2901/gin-web/docs"
	appCtx "github.com/yeyee2901/gin-web/src/context"
	"github.com/yeyee2901/gin-web/src/route"
)

func RunApp() {
	// LOAD CONFIG FILE
	config := LoadConfigFile()

	// INIT HTTP CLIENT
	httpClient := InitHttpClient(config)

	// INIT CONTEXT MANAGER
	configContext := context.WithValue(context.Background(), appCtx.Key("config"), config)
	httpContext := context.WithValue(context.Background(), appCtx.Key("http"), httpClient)
	AppContext := &appCtx.AppContext{
		Config: configContext,
		Http:   httpContext,
	}

	// INIT APP ENGINE (engine kosong tanpa logger dan middleware)
	app := gin.New()
	gin.SetMode(config.Aplikasi.Mode)

	// INIT LOGGER
	InitLogger(app)

	// INIT PANIC HANDLER
	InitPanicHandler(app)

	// ENABLE CORS
	route.EnableCors(app)

	// INIT ROUTING APLIKASI
	route.InitEngine(app, AppContext)

	// INIT SWAGGER
	docs.SwaggerInfo_swagger.BasePath = "/"
	app.GET("/api/doc/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	app.Run(":8080")
}
