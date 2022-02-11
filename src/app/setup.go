package app

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"gopkg.in/yaml.v2"

	"github.com/gin-gonic/gin"
	"github.com/yeyee2901/gin-web/src/context"
	"github.com/yeyee2901/gin-web/src/utils"
)

// LoadConfigFile godoc
// Load konfigurasi aplikasi dari config/config.yaml
func LoadConfigFile() (configStruct *context.Config) {
	f, err := ioutil.ReadFile("config/config.yaml")

	if err != nil {
		panic(errors.New("[FAIL] Fail to parse config file!"))
	}

	err = yaml.Unmarshal(f, &configStruct)

	if err != nil {
		panic(errors.New("[FAIL] Fail to unmarshal config file!"))
	}

	return configStruct
}

// InitHttpClient godoc
// Di golang, object http.Client aman untuk digunakan dalam konteks asinkron
// Object ini digunakan jika ingin interkoneksi antar middleware melalui
// koneksi HTTP. HTTP Client akan di konfigurasi berdasarkan config.yaml
func InitHttpClient(config *context.Config) *http.Client {
	timeout := time.Duration(config.Http.Timeout) * time.Second
	maxClient := config.Http.MaxClient

	transport := &http.Transport{
		MaxConnsPerHost:       maxClient,
		IdleConnTimeout:       timeout,
		ResponseHeaderTimeout: timeout,
	}

	return &http.Client{
		Transport: transport,
		Timeout:   timeout,
	}
}

// InitLogger godoc
// Untuk inisialisasi Logger.
// gin.Mode() == "debug"    -> stdout & gin-stdout.log
// gin.Mode() == "release"  -> gin-stdout.log
func InitLogger(app *gin.Engine) {

	loggerMiddleware := new(gin.LoggerConfig)
	logfile, err := os.OpenFile("logs/gin-stdout.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// set format logging nya
	loggerMiddleware.Formatter = func(formatParam gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] - %d - \"%s %s %s\" - \"%s\" %s %s\n",
			formatParam.ClientIP,
			formatParam.TimeStamp.Format(time.RFC1123),
			formatParam.StatusCode,
			formatParam.Request.Proto,
			formatParam.Method,
			formatParam.Path,
			formatParam.Request.UserAgent(),

			// untuk membaca request error (jika ada)
			utils.ParseHttpError(formatParam.ErrorMessage),

			// untuk membaca request body
			utils.ReadRequestBody(formatParam.Request.Body),
		)

	}

	// set agar Gin log ke stdout & ke file
	if gin.Mode() == gin.DebugMode {
		loggerMiddleware.Output = io.MultiWriter(logfile, os.Stdout)
	} else if gin.Mode() == gin.ReleaseMode {
		loggerMiddleware.Output = io.MultiWriter(logfile)
	} else {
		fmt.Println("[ERROR]: mode aplikasi tidak dikenali. \"debug\" atau \"release\" saja!")
		os.Exit(1)
	}

	// load logger
	app.Use(gin.LoggerWithConfig(*loggerMiddleware))
}

// InitPanicHandler godoc
// Inisialisasi Logger untuk panic
// Ketika ada panic, maka akan di write ke:
// gin.Mode() == "debug"    -> stderr & logs/gin-stderr.log
// gin.Mode() == "release"  -> logs/gin-stderr.log
func InitPanicHandler(app *gin.Engine) {
	panicLogfile, err := os.OpenFile("logs/gin-stderr.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// jika mode aplikasi sudah release, tidak perlu log ke stderr
	var writer io.Writer
	if gin.Mode() == gin.DebugMode {
		writer = io.MultiWriter(panicLogfile, os.Stderr)
	} else if gin.Mode() == gin.ReleaseMode {
		writer = io.MultiWriter(panicLogfile)
	} else {
		fmt.Println("[ERROR]: mode aplikasi tidak dikenali. \"debug\" atau \"release\" saja!")
		os.Exit(1)
	}

	// buat middleware untuk handle panic
	// middleware akan write ke logfile khusus untuk panic dan
	// juga write ke stderror
	panicLoggerMiddleware := gin.CustomRecoveryWithWriter(writer, func(ctx *gin.Context, recovered interface{}) { // kalau ada string nya berarti di log saja
		if err, ok := recovered.(string); ok {
			msg := fmt.Sprintf("INTERNAL SERVER ERROR: %s", err)
			ctx.String(http.StatusInternalServerError, msg)
		}

		ctx.AbortWithStatus(http.StatusInternalServerError)
	})

	// load panic handler
	app.Use(panicLoggerMiddleware)
}
