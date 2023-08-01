package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	// side effects
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "cesa_backend"
)

func main() {
	// connStr := "postgres://postgres:postgres@localhost/cesa_backend?sslmode=disable"

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	checkError(err)
	defer db.Close()

	wg := sync.WaitGroup{}
	wg.Add(1)

	// liveness
	go func(db *sql.DB) {
		for i := 0; i < 3; i++ {
			err := db.Ping()
			if err != nil {
				log.Println("failed db ping:", err)
			} else {
				log.Println("pinged successfully!")
			}

			time.Sleep(1 * time.Second)
		}

		wg.Done()
	}(db)

	err = db.Ping()
	checkError(err)

	fmt.Println("Successfully connected!")

	// drop table
	dropTable(db)

	// create table
	createTable(db)

	// insert statis
	insertStatic(db)

	// query
	querySpecific(db)

	// insert dynamic
	insertDynamic(db)

	// query
	query(db)

	// update
	update(db)

	// query
	querySpecific(db)

	// delete
	delete(db)

	// query
	querySpecific(db)

	// query
	query(db)

	wg.Wait()
}

func dropTable(db *sql.DB) {
	dropStmt := `drop table if exists "students"`
	_, e := db.Exec(dropStmt)
	checkError(e)
}

func createTable(db *sql.DB) {
	createStmt := `create table "students" (
		id serial primary key,
		name varchar(20) not null,
		phone varchar(15)
	)`
	_, e := db.Exec(createStmt)
	checkError(e)
}

func insertStatic(db *sql.DB) {
	insertStmt := `insert into "students"("name", "phone") values ('Ali', '09120000001')`
	_, e := db.Exec(insertStmt)
	checkError(e)
}

func insertDynamic(db *sql.DB) {
	userName, phone := getCredentials()
	insertDynStmt := `insert into "students"("name", "phone") values($1, $2)`
	_, e := db.Exec(insertDynStmt, userName, phone)
	checkError(e)
}

func getCredentials() (string, string) {
	return "mohammad", "09120000003"
}

func update(db *sql.DB) {
	updateStmt := `update "students" set "phone"=$1 where "name"=$2`
	_, e := db.Exec(updateStmt, "09120000008", "Ali")
	checkError(e)
}

func delete(db *sql.DB) {
	deleteStmt := `delete from "students" where "name"=$1`
	_, e := db.Exec(deleteStmt, "Mohammad")
	checkError(e)
}

func querySpecific(db *sql.DB) {
	rows, err := db.Query(`SELECT "name", "phone" FROM "students"`)
	checkError(err)
	defer rows.Close()

	for rows.Next() {
		var (
			name, phone string
		)
		err = rows.Scan(&name, &phone)
		checkError(err)

		fmt.Println("student rec (specific):", name, phone)
	}

	checkError(err)
}

type Student struct {
	ID    int
	Name  string
	Phone string
}

func query(db *sql.DB) {

	rows, err := db.Query(`SELECT * FROM "students"`)
	checkError(err)

	defer rows.Close()
	for rows.Next() {
		std := Student{}
		// not safe at all!
		err = rows.Scan(&std.ID, &std.Name, &std.Phone)
		checkError(err)

		fmt.Printf("student rec: %+v\n", std)
	}

	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal("failed:", err)
	}
}
