package media

import "strings"

/*  This exportable.
type Movie struct {
	Title     string
	Rating    Rating
	BoxOffice float32
} */

type Catalogable interface {
	NewMovie(title string, rating Rating, boxOffice float32)
	GetTitle() string
	GetRating() string
	GetBoxOffice() float32
	SetTitle(newTitle string)
	SetRating(newRating Rating)
	SetBoxOffice(newBoxOffice float32)
}

type Movie struct {
	title     string
	rating    Rating
	boxOffice float32
}

type Rating string

const (
	R    = "R (Restricted)"
	G    = "G (General audiences)"
	PG   = "PG (Parental Guidance)"
	PG13 = "PG-13 (Parental Caution)"
	NC17 = "NC-17 (No children under 17)"
)

/*
func NewMovie(title string, rating Rating, boxOffice float32) Movie {
	return Movie{
		title:     title,
		rating:    rating,
		boxOffice: boxOffice,
	}
} */
// Dont want to return a copy of moview instead use same copy.
func (m *Movie) NewMovie(title string, rating Rating, boxOffice float32) {
	m.title = title
	m.rating = rating
	m.boxOffice = boxOffice
}

func (m *Movie) GetTitle() string {
	return strings.ToTitle(m.title)
}

func (m *Movie) GetRating() string {
	return string(m.rating)
}

func (m *Movie) GetBoxOffice() float32 {
	return m.boxOffice
}

func (m *Movie) SetTitle(newTitle string) {
	m.title = newTitle
}

func (m *Movie) SetRating(newRating Rating) {
	m.rating = newRating
}

func (m *Movie) SetBoxOffice(newBoxOffice float32) {
	m.boxOffice = newBoxOffice
}
