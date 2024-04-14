package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"golang.org/x/sync/errgroup"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Course struct {
	Name   string
	Rating string
	Link   string
	Diff   string
	Dur    string
	Skills string
	Price  string
	Num    string
}

var DB *sql.DB

// handle default page
func defaultPage(w http.ResponseWriter) {
	var fileName = "main-page.html"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		log.Printf("Error while parsing the files: %s", err)
		return
	}
	err = t.ExecuteTemplate(w, fileName, nil)
	if err != nil {
		log.Printf("Error while executing the files: %s", err)
		return
	}
}

// handle page with expressions
func listCourses(w http.ResponseWriter, r *http.Request) {
	var courses []Course

	query, err := DB.Query("SELECT * FROM public.ways")

	for query.Next() {
		var course Course
		err = query.Scan(&course.Name, &course.Rating, &course.Link, &course.Diff, &course.Skills, &course.Price, &course.Rating)
		if err != nil {
			log.Fatalln(err)
		}
		courses = append(courses, course)
	}

	var fileName = "personal-page.html"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		log.Printf("Error while parsing the files: %s", err)
		return
	}
	err = t.ExecuteTemplate(w, fileName, courses)
	if err != nil {
		log.Printf("Error while executing the files: %s", err)
		return
	}
}

func login(w http.ResponseWriter) {
	var fileName = "login.html"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		log.Printf("Error while parsing the files: %s", err)
		return
	}
	err = t.ExecuteTemplate(w, fileName, nil) // executing page with data from Cache
	if err != nil {
		log.Printf("Error while executing the files: %s", err)
		return
	}
}

// page handlers
func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		defaultPage(w)
	case "/login":
		login(w)
	case "/listCourses":
		listCourses(w, r)
	default:
		http.Error(w, "404 PAGE NOT FOUND", http.StatusNotFound)
	}
}

// deploying server
func main() {
	// deploying server and graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	var err error

	go func() {
		c := make(chan os.Signal, 1) // we need to reserve to buffer size 1, so the notifier are not blocked
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		<-c
		cancel()
	}()

	connStr := "user=postgres password=s dbname=lila"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(err)
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handler)

	httpServer := &http.Server{
		Addr: ":8000",
	}

	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return httpServer.ListenAndServe()
	})

	g.Go(func() error {
		<-gCtx.Done()

		err = DB.Close()
		if err != nil {
			log.Fatalln(err)
		}

		return httpServer.Shutdown(context.Background())
	})

	if err = g.Wait(); err != nil {
		fmt.Printf("exit reason: %s \n", err)
	}
}
