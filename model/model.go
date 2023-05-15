package model

type Books struct {
	Book_id        string `json:"book_id"`
	Year_published string `json:"year_published"`
	Book_title     string `json:"book_title"`
	Author_name    string `json:"author_name"`
	Book_category  string `json:"book_category"`
}

type Item struct {
	Book_id    int    `json:"book_id",bun:"book_id"`
	Start_date string `json:"start_date",bun:"start_date"`
	End_date   string `json:"end_date",bun:"end_date"`
	Person_id  string `json:"person_id",bun:"person_id"`
}

type Member struct {
	Person_id   int    `json:"person_id",bun:"person_id"`
	Name        string `json:"name",bun:"name"`
	PhoneNumber int    `json:"phoneNumber",bun:"phoneNumber"`
	Email       string `json:"email",bun:"email"`
	Department  string `json:"department",bun:"department"`
}

type Location struct {
	Book_id  int `json:"book_id",bun:"book_id"`
	Floor_no int `json:"floor_no",bun:"floor_no"`
	Shelf_no int `json:"shelf_no",bun:"shelf_no"`
}
