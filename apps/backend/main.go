package main

import (
	"fmt"
	"net/http"
	"xpends/webapi/configs"
	"xpends/webapi/router"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}
	r := router.Router()
	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
