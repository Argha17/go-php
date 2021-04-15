package handler

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	// "os"
	"time"
)

func DbConnect() *sql.DB {
	// os.Remove("./movies.db")
	db, err := sql.Open("sqlite3", "./movies.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func DbInit() {
	db := DbConnect()
	sqlStmt := `
	create table if not exists movies 
		(id integer not null primary key, 
			title text not null,
			year text not null,
			imdbid text not null,
			type text not null,
			poster text not null,
			created_at text not null,
			update_at text not null
			);
	`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}
	defer db.Close()

	log.Println("success to create table movies")
}

func DbStoreData() {
	db := DbConnect()
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into movies(title, year, imdbid, type, poster, created_at, update_at) values(?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	
	_, err = stmt.Exec("batman", 
					"2010", 
					"182112", 
					"action", 
					"https://m.media-amazon.com/images/M/MV5BMTdjZTliODYtNWExMi00NjQ1LWIzN2MtN2Q5NTg5NTk3NzliL2ltYWdlXkEyXkFqcGdeQXVyNTAyODkwOQ@@._V1_SX300.jpg",
					time.Now(),
					time.Now())
	if err != nil {
		log.Fatal(err)
	
	}
	tx.Commit()
	defer db.Close()
}

func DbQuery(pagination int, keyword string) ([]Movie, error) {
	db := DbConnect()

	rows, err := db.Query("select id, title, year, imdbid, type, poster from movies limit ?", pagination)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	defer db.Close()

	var movie Movie
	var movies []Movie

	for rows.Next() {
		err = rows.Scan(&movie.Id, 
			&movie.Title,
			&movie.Year,
			&movie.Imdbid,
			&movie.Type,
			&movie.Poster,
		)
		if err != nil {
			log.Fatal(err)
		} else {
			movies = append(movies, movie)
		}
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return movies, err
	
}