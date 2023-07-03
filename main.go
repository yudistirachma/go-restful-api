package main

import (
	"net/http"
	"yudistirachma/go-restful-api/app"
	"yudistirachma/go-restful-api/helper"
	"yudistirachma/go-restful-api/middleware"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := app.NewRouter()

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
