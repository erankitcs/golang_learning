package main

import (
	"net/http"

	"github.com/erankitcs/golang_learning/webservice/controllers"
)

func main() {
	// u := models.User{
	// 	ID:        2,
	// 	FirstName: "Ankit",
	// 	LastName:  "Singh",
	// }
	// fmt.Println(u)
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
