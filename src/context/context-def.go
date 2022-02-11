package context

import "context"

// untuk mendapatkan nilai dari context manager, maka harus pakai ini
//  ctxKey := Key("some context")
//  ctxInterface := ctxObject.Value(ctxKey)
//
//  // check type nya
//  value, ok := ctxInterace.( SomeType )
//  if !ok {
//    panic("failed to load context")
//  }
//
//  fmt.Printf("val type %T", value)
type Key string

// App context keseluruhan
type AppContext struct {
    // Config context sesuai yang di load di inisialisasi
	Config context.Context

    // Http context sesuai yang di load di inisialisasi
	Http   context.Context
}

// CONFIG CONTEXT
type Config struct {
    // Konfigurasi aplikasi general
	Aplikasi   aplikasi      `yaml:"aplikasi"`

    // Konfigurasi token (expired dll)
	Token      token         `yaml:"token"`

    // API key untuk sign token
	ApiKey     apiKey        `yaml:"apikey"`

    // Konfigurasi HTTP client
	Http       http          `yaml:"http"`
}

// CONFIG CONTEXT --------------------------------
type aplikasi struct {
	Mode      string `yaml:"mode"`
	TmpFolder string `yaml:"tmp_folder"`
}

type token struct {
	ExpTimeRegist int `yaml:"exp_time_regist"`
	ExpTimeAccess int `yaml:"exp_time_access"`
}

type apiKey struct {
	Key string `yaml:"key"`
}

type http struct {
	Timeout   int `yaml:"timeout"`
	MaxClient int `yaml:"max_client"`
}
