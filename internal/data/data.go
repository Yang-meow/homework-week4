package data

import (
	"database/sql"
	"homework-week4/internal/conf"
	"os"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	_ "github.com/mattn/go-sqlite3"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewLibraryRepo)

// Data .
type Data struct {
	// TODO warpped database client
	db *sql.DB
}

type Members struct {
	Id   int64
	Name string
}

type Book struct {
	Id     int64
	Name   string
	Isbn   string
	Author string
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		logger.Log("msg", "closing the data resources")
	}
	os.Remove("./library.db")
	file, err := os.Create("library.db")
	if err != nil {
		panic("create library failed")
	}
	file.Close()
	db, err := sql.Open("sqlite3", "./library.db")
	if err != nil {
		panic("open db failed")
	}
	//defer db.Close()
	createTable(db)
	insertData(db)
	data := &Data{db: db}
	return data, cleanup, nil
}

func createTable(db *sql.DB) {
	createMembersSQL := `CREATE TABLE members (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"name" TEXT	
	  );` // SQL Statement for Create Table

	statement, err := db.Prepare(createMembersSQL) // Prepare SQL Statement
	if err != nil {
		panic("create members failed")
	}
	statement.Exec() // Execute SQL Statements

	createBookSQL := `CREATE TABLE book (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"name" TEXT,
		"isbn" TEXT,
		"author" TEXT		
	  );` // SQL Statement for Create Table

	statement2, err := db.Prepare(createBookSQL) // Prepare SQL Statement
	if err != nil {
		panic("create book failed")
	}
	statement2.Exec() // Execute SQL Statements

	createOpsSQL := `CREATE TABLE ops (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"member_id" integer,
		"book_id" integer,
		"borrow_date" integer,
		"return_date" integer
      );`
	st, err := db.Prepare(createOpsSQL)
	if err != nil {
		panic("create ops failed")
	}
	st.Exec()
}

func insertData(db *sql.DB) {
	insertSQL := `INSERT INTO members (name) VALUES (?)`
	statement, err := db.Prepare(insertSQL)
	if err != nil {
		panic("insert failed")
	}
	_, err = statement.Exec("test_name")
	if err != nil {
		panic("insert failed")
	}

	insertSQL = `INSERT INTO book (name, isbn, author) VALUES (?, ?, ?)`
	statement, err = db.Prepare(insertSQL)
	if err != nil {
		panic("insert failed")
	}
	_, err = statement.Exec("test_book_name", "test_isbn_123456", "test_author")
	if err != nil {
		panic("insert failed")
	}
}
