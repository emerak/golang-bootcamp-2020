package main

import (
	"fmt"
	"log"

	"github.com/emerak/golang-bootcamp-2020/config"
	//"github.com/emerak/golang-bootcamp-2020/infrastructure/datastore"
	"github.com/emerak/golang-bootcamp-2020/infrastructure/router"
	"github.com/emerak/golang-bootcamp-2020/registry"
	"github.com/labstack/echo"
)

func main() {
	config.ReadConfig()

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	fmt.Println("Server listen at http://localhost" + ":" + config.C.Server.Address)
	if err := e.Start(":" + config.C.Server.Address); err != nil {
		log.Fatalln(err)
	}
}
