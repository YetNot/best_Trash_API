package models

var DB []Book

type Book struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Author        Author `json:"author"`
	YearPublished int    `json:"year_published"`
}

type Author struct {
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	BornYear int    `json:"born_year"`
}

func init() {
	book1 := Book{
		ID: 1,
		Title: "Lord of the Rings",
		YearPublished: 1978,
		Author: Author{
			Name: "J.J",
			LastName: "Tolkin",
			BornYear: 1892,
		},
	}
	DB = append(DB, book1)
}

func FindBookId(id int) (Book, bool) {
	var book Book
	var found bool
	for _, b := range DB {
		if b.ID == id {
			book = b
			found = true
			break
		}
	}
	return book, found
}