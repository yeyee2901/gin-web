# Percobaan Gin Framework
Percobaan membuat API dengan bahasa Go-lang. 
![image](https://user-images.githubusercontent.com/55247343/153563216-1f62d9e2-169e-49cf-a95c-bdb5e1daa76b.png)


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
mkdir deps
```
3. Set Environment Variable GOPATH & install dependency project secara lokal terhadap project
```bash
export GOPATH=$(pwd)/deps
go get
```
4. Download and run `swaggo`
```bash
go install github.com/swaggo/swag/cmd/swag@latest

# di direktori root project:
swag init
```
5. Run Project
```bash
go run .   # OR
go run main.go
```

## Notes
- Setiap kali edit dokumentasi swagger, run `swag init` lagi untuk re-generate file docs yaml
