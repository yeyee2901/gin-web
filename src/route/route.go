package route

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	appContext "github.com/yeyee2901/gin-web/src/context"
	c "github.com/yeyee2901/gin-web/src/controller"
)

// InitEngine godoc
func InitEngine(app *gin.Engine, ctx *appContext.AppContext) {
	// type cast
	appCtx := c.AppContext(*ctx)

	// LOAD MIDDLEWARES
	app.Use(appCtx.AuthMiddleware)

	// GET ROUTES
	var URL_GET = map[string]gin.HandlerFunc{
		"/cek_middleware": appCtx.CekMiddleware,
		"/cek_http": appCtx.CekHttp,
	}

	// POST ROUTES
	var URL_POST = map[string]gin.HandlerFunc{
		"/login": appCtx.Login,
	}

	// POPULATE ROUTE DENGAN LOOPING
	for pathGET, handler := range URL_GET {
		app.GET(pathGET, handler)
	}

	for pathPOST, handler := range URL_POST {
		app.POST(pathPOST, handler)
	}
}

// EnableCors godoc
// Enable CORS (Cross Origin Resource Sharing) untuk aplikasi ini
func EnableCors(app *gin.Engine) {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = append(corsConfig.AllowOrigins, "*")
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.ExposeHeaders = []string{"*"}

	// load the middleware
	app.Use(cors.New(corsConfig))
}
