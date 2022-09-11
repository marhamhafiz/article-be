package main

import (
	p "article-be/posts"
)

// func connectToDB() *sql.DB {
// 	fmt.Println("Connected to Database")

// 	// Open up our database connection.
// 	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/articles")

// 	// if there is an error opening the connection, handle it
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	// // executing
// 	// defer db.Close()

// 	return db
// }

// func createArticle(db *sql.DB) {
// 	// perform a db.Query insert
// 	insert, err := db.Query("INSERT INTO posts (title,content,category,status) VALUES ('Apa lagi cok', 'mengapa jadi begini', 'Paper', 'Trash')")

// 	// if there is an error inserting, handle it
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	fmt.Println("Insert to Database success")
// 	// be careful deferring Queries if you are using transactions
// 	defer insert.Close()
// }

func main() {
	p.HandleRequests()
}
