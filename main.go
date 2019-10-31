package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/kudejwiktor/go-api-example/app/servid"
)

func main() {
	server := servid.NewApp()
	server.Start()

	//fmt.Println("Go MySQL Tutorial")
	//
	//// Open up our database connection.
	//// I've set up a database on my local machine using phpmyadmin.
	//// The database is called testDb
	//db, err := sql.Open("mysql", "admin:Admin.123@tcp(go-api-example-db:3306)/users_db")
	//
	//// if there is an error opening the connection, handle it
	//if err != nil {
	//	panic(err.Error())
	//}
	//
	//defer db.Close()
	//
	//// Execute the query
	//results, err := db.Query("SELECT id, name FROM users")
	//if err != nil {
	//	panic(err.Error()) // proper error handling instead of panic in your app
	//}
	//
	//for results.Next() {
	//	var tag User
	//	// for each row, scan the result into our tag composite object
	//	err = results.Scan(&tag.ID, &tag.Name)
	//	if err != nil {
	//		panic(err.Error()) // proper error handling instead of panic in your app
	//	}
	//	// and then print out the tag's Name attribute
	//	log.Printf(tag.Name)
	//}

}
