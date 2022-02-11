# Percobaan Gin Framework
Percobaan membuat API dengan bahasa Go-lang. 

## Tech Stack
- [Gin](https://github.com/gin-gonic/gin) - Web framework
- [Gin Swaggo](https://github.com/swaggo/gin-swagger) - Swagger Docs integration for Gin web framework
- [Golang JWT v4](https://github.com/golang-jwt/jwt/v4) - JWT library for Go
- [Go Yaml v2](https://github.com/go-yaml/yaml/tree/v2.4.0) - Go library for processing yaml files

## Brief
Pada percobaan API ini hanya menyertakan 2 routing & 1 middleware.
- GET: `/cek_middleware`
- POST: `/login`
- Middleware: `AuthMiddleware`

## Getting Started
1. Clone project nya
```bash
git clone https://github.com/yeyee2901/gin-web.git
```
2. Buat direktori `/logs` & `/deps`
```bash
cd gin-web
mkdir logs
mkidr deps
```
3. Set Environment Variable GOPATH & install dependency project secara lokal terhadap project
```bash
export GOPATH=$(pwd)/deps
go get
```
4. Run Project
```bash
go run .   # OR
go run main.go
```
