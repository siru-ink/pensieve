package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	_ "modernc.org/sqlite"
)

type GuestbookEntry struct {
	Name    string
	Content string
	SiteUrl string
}

var db *sql.DB

func main() {
	err := initDatabase()
	if err != nil {
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			fmt.Fprintf(w, "You have ended up somewhere that only computers are supposed to be. Feel free to close this tab.")
			return
		}

		// fmt.Fprintf(w, "N: %s, Content: %s, SiteUrl: %s",
		// 	r.FormValue("name"), r.FormValue("content"), r.FormValue("siteurl"))

		entry := GuestbookEntry{
			Name:    r.FormValue("name"),
			Content: r.FormValue("content"),
			SiteUrl: r.FormValue("siteurl"),
		}

		err = addentry(entry)
		if err != nil {
			fmt.Println("Something went wrong.")
		}

		http.Redirect(w, r, r.FormValue("siteurl"), http.StatusSeeOther)
	})

	http.ListenAndServe(":8010", nil)
}

func initDatabase() error {
	var err error
	db, err = sql.Open("sqlite", "guestbook.db")
	if err != nil {
		return err
	}
	_, err = db.ExecContext(
		context.Background(),
		`create table if not exists 'guestbookentries' (
			id integer not null unique,
			name text not null,
			content text not null,
			siteurl text not null,
			primary key('id' autoincrement)
		);`,
	)
	if err != nil {
		return err
	}
	return nil
}

func addentry(entry GuestbookEntry) error {
	_, err := db.ExecContext(
		context.Background(),
		`insert into guestbookentries (name, content, siteurl) VALUES (?,?,?);`, entry.Name, entry.Content, entry.SiteUrl,
	)
	if err != nil {
		return err
	}
	return nil
}
