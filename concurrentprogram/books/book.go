package books

import "fmt"

type Book struct {
	ID            int
	Title         string
	Author        string
	YearPublished int
}

func (b *Book) String() string {
	return fmt.Sprintf(
		"Title:\t\t%q\n"+
			"Author:\t\t%q\n"+
			"Published:\t%v\n", b.Title, b.Author, b.YearPublished)
}

var MyBook = []Book{
	Book{
		ID:            1,
		Title:         "The Hitchhiker's Guide to the Galaxy",
		Author:        "Douglas Adams",
		YearPublished: 1979,
	},
	Book{
		ID:            2,
		Title:         "The Hobbit",
		Author:        "J.R.R. Tolkien",
		YearPublished: 1937,
	},
	Book{
		ID:            3,
		Title:         "A Tale of Two Cities",
		Author:        "Charles Dickens",
		YearPublished: 1859,
	},
	Book{
		ID:            4,
		Title:         "Les Misérables",
		Author:        "Victor Hugo",
		YearPublished: 1862,
	},
	Book{
		ID:            5,
		Title:         "Harry Potter and the Philosopher's Stone",
		Author:        "J.K. Rowling",
		YearPublished: 1997,
	},
	Book{
		ID:            6,
		Title:         "I, Robot",
		Author:        "Isaac Asimov",
		YearPublished: 1950,
	},
	Book{
		ID:            7,
		Title:         "The Gods Themselves",
		Author:        "Isaac Asimov",
		YearPublished: 1973,
	},
	Book{
		ID:            8,
		Title:         "The Moon is a Harsh Mistress",
		Author:        "Robert A. Heinlein",
		YearPublished: 1966,
	},
	Book{
		ID:            9,
		Title:         "On Basilisk Station",
		Author:        "David Weber",
		YearPublished: 1993,
	},
	Book{
		ID:            10,
		Title:         "The Android's Dream",
		Author:        "John Scalzi",
		YearPublished: 2006,
	},
}
