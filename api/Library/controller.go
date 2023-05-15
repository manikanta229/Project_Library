package Library

import (
	"Project_Library/model"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

var Books []model.Books

func Postby(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Book model.Books
	err := json.NewDecoder(r.Body).Decode(&Book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	Books = append(Books, Book)

	my_db := Connect()
	// fmt.Println("uyyhs", Book.Book_id, Book.Author_name)
	Adding(my_db, Book)
	json.NewEncoder(w).Encode(Book)
}

func Deleteby(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range Books {
		if (item.Book_id) == params["book_id"] {
			Books = append(Books[:index], Books[index+1:]...)
			my_db := Connect()
			Deleting(my_db, item.Book_id)
			break
		}
	}
	json.NewEncoder(w).Encode(Books)
}

func Getby(w http.ResponseWriter, r *http.Request) {
	fmt.Println("You are in GET books ")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range Books {
		if params["Book_id"] == item.Book_id {
			json.NewEncoder(w).Encode(item)
			my_db := Connect()
			Getting(my_db, item)
			return
		}
	}
}

func Putby(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Book model.Books
	params := mux.Vars(r)

	for index, items := range Books {
		if items.Book_id == params["Book_id"] {
			Books = append(Books[:index], Books[index+1:]...)
			_ = json.NewDecoder(r.Body).Decode(&Book)
			Book.Book_id = params["Book_id"]
			my_db := Connect()
			Updating(my_db, Book)

			Books = append(Books, Book)
			json.NewEncoder(w).Encode(Book)
			return
		}
	}
}
