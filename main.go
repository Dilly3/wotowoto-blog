package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 8089
	user     = "dilly"
	password = 0000
	dbname   = "mac"
)

var Postgresdb *sql.DB
var BlogsQueue H.Blogger
var err error

//var Blgs H.Blogger

func main() {
	//Blgs = H.Blogger{}

}

func ViewBlog(writer http.ResponseWriter, req *http.Request) {
	RenderPage(writer, "./tmpl/basetemplate.gohtml", BlogsQueue.Blogs)

}

func RenderPage(writer http.ResponseWriter, addr string, data interface{}) {
	parsedTemplate, err := template.ParseFiles(addr)

	if err != nil {
		H.Println(err)
	}

	parsedTemplate.Execute(writer, addr)

}

func RenderHomePage(writer http.ResponseWriter, addr string) {
	parsedTemplate, err := template.ParseFiles(addr)

	if err != nil {
		H.Println(err)
	}

	parsedTemplate.Execute(writer, BlogsQueue.Blogs)

}

func Home(writer http.ResponseWriter, req *http.Request) {
	RenderHomePage(writer, "./tmpl/frontpage.gohtml")
}
func BlogPage(writer http.ResponseWriter, req *http.Request) {
	RenderHomePage(writer, "./tmpl/basetemplate.gohtml")
}

func About(writer http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(writer, "This is my About page")
}

//	http.Redirect(writer, req, "/", 302)
