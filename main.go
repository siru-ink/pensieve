package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "modernc.org/sqlite"
)

type CommentEntry struct {
	Name    string
	Comment string
	SiteUrl string
}

var db *sql.DB

func init() {
	// Open a database (internally stores a database pool for concurrent use)
	var err error
	db, err = sql.Open("sqlite", "comments.db")
	if err != nil {
		log.Fatalf("Database [comments.db] could not be opened: %v\n", err)
	}

	// Create the table to store entries
	_, err = db.ExecContext(
		context.Background(),
		"CREATE TABLE IF NOT EXISTS comments (id INTEGER NOT NULL UNIQUE PRIMARY KEY"+
			"AUTOINCREMENT, name TEXT NOT NULL, comment TEXT NOT NULL, siteurl TEXT NOT NULL);",
	)
	if err != nil {
		log.Fatalf("Database table [comments] could not be created: %v\n", err)
	}

	// Log a startup message to indicate db is redaable/editable
	log.Println("Database connection to [comments.db] was constructed.")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.Header().Set("Allow", "POST")
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		entry := CommentEntry{
			Name:    r.FormValue("name"),
			Comment: r.FormValue("comment"),
			SiteUrl: r.FormValue("siteurl"),
		}

		err := addentry(entry)
		if err != nil {
			fmt.Fprintf(w, "Something went wrong in saving your comment. Please be so kind to contact me via email at email [at] siru [dot] ink, so that I can try to resolve the problem.")
			return
		}

		http.Redirect(w, r, r.FormValue("siteurl"), http.StatusSeeOther)
	})

	http.ListenAndServe(":8020", nil)
}

func addentry(entry CommentEntry) error {
	_, err := db.ExecContext(
		context.Background(),
		`INSERT INTO comments (name, comment, siteurl) VALUES (?,?,?);`, entry.Name, entry.Comment, entry.SiteUrl,
	)
	if err != nil {
		log.Printf("Comment could not be added to database [%+v]: %v\n", entry, err)
		return err
	}
	return nil
}
