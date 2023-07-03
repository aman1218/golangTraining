package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "testdb"
)

func main() {

	type student struct {
		id    int
		fname string
		lname string
		age   int
	}

	postgresqlDbInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", postgresqlDbInfo)
	checkErr(err)
	err = db.Ping()
	checkErr(err)
	fmt.Println("connected to database!")

	insertStmt := `insert into "students"(id, fname, lname, age) values(1, 'Aman', 'Prasad', 23),(2, 'Bijay', 'Sharma', 24),(3, 'Satyam', 'Gupta', 22)`
	_, err = db.Query(insertStmt)
	checkErr(err)
	fmt.Println("Inserted data successfully")

	viewStmt := "select * from students"
	data, err := db.Query(viewStmt)
	checkErr(err)

	for data.Next() {
		var s student
		err = data.Scan(&s.id, &s.fname, &s.lname, &s.age)
		checkErr(err)
		fmt.Println(s)
	}
	db.Query("delete from students where 1=1")
	defer db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
