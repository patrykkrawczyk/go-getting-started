package main

import (
	"github.com/patrykkrawczyk/go-getting-started/controllers"
	"net/http"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
