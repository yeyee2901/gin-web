package main

import "github.com/yeyee2901/gin-web/src/app"

// @Title Coba Gin Golang
// @Version 0.0.1
// @Schemes http https
// @Contact.email gabriel777sh@gmail.com
//
// @Description Percobaan Swagger & Gin pada Golang untuk pembuatan API
// @Description 
// @Description <h4>Note: HTTP status dari response API secara umum:</h4>
// @Description - <strong>200</strong> : Sukses
// @Description - <strong>206</strong> : Error lain
// @Description - <strong>209</strong> : Data input salah / tidak valid
// @Description - <strong>210</strong> : Error middleware internal
// @Description - <strong>214</strong> : Data tidak ditemukan
// @Description - <strong>401</strong> : Token invalid
// @Description - <strong>500</strong> : Internal server error (fatal)
// @Description 
// @Description Jika response code tidak <strong>200</strong>, maka yang dikembalikan adalah JSON dengan 1 elemen "msg"
//
// @tag.name login
// @tag.description API yang digunakan untuk login
//
// @tag.name test
// @tag.description API yang digunakan untuk coba-coba

func main() {
   app.RunApp() 
}
