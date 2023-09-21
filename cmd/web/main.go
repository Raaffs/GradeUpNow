package main

/*
Function/varibales in different packages(e.g in theory folder) must start with capital letter
if you want to access them in other package(e.g main). And access them with package_name.FunctionName().


Setup:
go get github.com/texttheater/golang-levenshtein/levenshtein	//for pattern matching
go get github.com/gorilla/mux									//to handle http request
go get github.com/go-sql-driver/mysql							//to manage database
go mod init github.com/Suy56/GradeUpNow
go mod tidy

Run program:
go ./cmd/web

Format to enter keywords in the database: 
"word1,word2,word3..."
ENTER ALL WORDS IN LOWER CASE, SEPARATE THEM WITH ',' AND LEAVE NO WHITE SPACE IN BETWEEN.
*/

import (
	"flag"
	"database/sql"
	//"html/template"
	"log"
	"net/http"
	"os"
	_"github.com/go-sql-driver/mysql"
	"github.com/Suy56/GradeUpNow/internal/models"
)

type application struct{
	errorLog *log.Logger
	infoLog *log.Logger
	user *models.Model	// wrapper around *sql.DB
}

//connects to the database pool
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
		return db, nil
	}


func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")								//default http address set to 4000
	dsn := flag.String("dsn", "root:root@/GradeUpNow?parseTime=true", "MySQL data source name")
	
	//we use the flag.Parse() function to parse the command-line flag.
	// This reads in the command-line flag value and assigns it to the addr
	// variable. You need to call this *before* you use the addr variable
	// otherwise it will always contain the default value of ":4000". If any errors are
	// encountered during parsing the application will be terminated.
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)
	db,err:=openDB(*dsn)
	if err!=nil{
		errorLog.Fatal(err)
	}
	defer db.Close()
	// Hold the application-wide dependencies for the webapp. 

	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
		user:	&models.Model{DB:db},
	}
	
	// Initialize a new http.Server struct. We set the Addr and Handler fields so
	// that the server uses the same network address and routes as before, and set
	// the ErrorLog field so that the server now uses the custom errorLog logger in
	// the event of any problems. 
	srv:=&http.Server{
		Addr: 		*addr,
		ErrorLog:	errorLog,
		Handler: 	app.routes(),	// Calls app.routes() method to get the servemux containing our routes.
	}
	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}


