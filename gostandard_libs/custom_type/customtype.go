package main

import (
	"fmt"

	"github.com/erankitcs/golang_learning/gostandard_libs/custom_type/media"
)

func main() {
	fmt.Println("My Fav movies")
	//myMov := media.Movie{}
	//myMov.Title = "Top Gun"
	//myMov.Rating = media.R
	//myMov.BoxOffice = 43.2
	//fmt.Printf("My fav movie is : %s \n", myMov.Title)
	//fmt.Printf("its rating is : %s \n", myMov.Rating)
	//fmt.Printf("It made %f in the box office. \n", myMov.BoxOffice)

	//myMov := media.NewMovie("Top Gun", media.R, 43.2)
	//fmt.Printf("My fav movie is : %s \n", myMov.GetTitle())
	//fmt.Printf("its rating is : %s \n", myMov.GetRating())
	//fmt.Printf("It made %f in the box office. \n", myMov.GetBoxOffice())
	//myMov.SetTitle("New Top Gun")
	//fmt.Printf("My fav movie is : %s \n", myMov.GetTitle())

	var myMov media.Catalogable = &media.Movie{}
	myMov.NewMovie("Top Gun", media.PG, 32.8)
	fmt.Printf("My fav movie is : %s \n", myMov.GetTitle())
	fmt.Printf("its rating is : %s \n", myMov.GetRating())
	fmt.Printf("It made %f in the box office. \n", myMov.GetBoxOffice())
	myMov.SetTitle("New Top Gun")
	fmt.Printf("My fav movie is : %s \n", myMov.GetTitle())
}
