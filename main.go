package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var tpl *template.Template
var db *sql.DB

func init() {
	tpl = template.Must(template.ParseGlob("tmpl/*.html"))
}

func main() {

	// Database first
	var err error
	db, err = sql.Open("mysql", "root:Chris95&1@/forgedfindsdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Println(err)
	}

	// HTTP Handlers
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/sell", sellHandler)
	http.HandleFunc("/shop", shopHandler)
	http.HandleFunc("/temp", tempHandler)

	http.HandleFunc("/interest", addInterested)

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server is shutting down...")
	}

}

func checkFatal(e error) {
	if e != nil {
		log.Fatal(e)
		os.Exit(0)
	}
}

func addInterested(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()

	firstname := req.Form["firstname"][0]
	lastname := req.Form["lastname"][0]
	email := req.Form["email"][0]
	customerType := req.Form["customer-type"][0]

	values := "\"" + firstname + "\", \"" + lastname + "\", \"" + email + "\", \"" + customerType + "\""

	statement, err := db.Prepare("INSERT IGNORE INTO interested (firstname, lastname, email, customerType) VALUES (" + values + ");")
	checkFatal(err)

	_, err = statement.Exec()
	checkFatal(err)

	err = tpl.ExecuteTemplate(w, "thanks.html", nil)
	if err != nil {
		log.Println(err)
	}
}

/*
	---------- Page handlers ----------
*/

func indexHandler(w http.ResponseWriter, req *http.Request) {
	tempHandler(w, req)
}

func aboutHandler(w http.ResponseWriter, req *http.Request) {
	tempHandler(w, req)
}

func sellHandler(w http.ResponseWriter, req *http.Request) {
	tempHandler(w, req)

}

func shopHandler(w http.ResponseWriter, req *http.Request) {
	tempHandler(w, req)
}

func tempHandler(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "temp.html", nil)
	if err != nil {
		log.Println(err)
	}
}
