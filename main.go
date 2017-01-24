package main

import (
	"log"
	"net/http"

	"github.com/colombia9503/RESTful-SqlServer/common"
	"github.com/colombia9503/RESTful-SqlServer/routers"
	"github.com/urfave/negroni"
)

func main() {
	common.StartUp()
	r := routers.InitRouters()

	n := negroni.Classic()
	n.UseHandler(r)

	log.Println("Listening.. 8081")
	http.ListenAndServe(":8081", n)
}
