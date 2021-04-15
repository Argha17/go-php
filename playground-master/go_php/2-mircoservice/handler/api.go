package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"encoding/json"
	"strconv"
	"time"
)

type Movie struct {
	Id			int 		`json:"id"`
	Title		string 		`json:"title"`
	Year		string 		`json:"year"`
	Imdbid		string 		`json:"imdbid"`
	Type		string 		`json:"type"`
	Poster		string 		`json:"poster"`
	CreatedAt	time.Time	`json:created_at`
	UpdateAt	time.Time	`json:update_at`
}

// type MovieUrl struct {
// 	Title		string 		`json:"Title"`
// 	Year		string 		`json:"Year"`
// 	Imdbid		string 		`json:"ImdbID"`
// 	Type		string 		`json:"Type"`
// 	Poster		string 		`json:"Poster"`
// }

// type Movies struct {
// 	Search 			[]MovieUrl		`json:Search`
// 	TotalResults	string			`json:totalResults`
// 	Response		string 			`json:Response`
// }	

type Response struct {
	Ok 		bool 		`json:ok`
	Data    []Movie
}

var result map[string]interface{}

func getJson(url string, target interface{}) error {
    r, err := http.Get(url)
    if err != nil {
        return err
    }
    defer r.Body.Close()

    return json.NewDecoder(r.Body).Decode(target)
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func SearchMoviews(c *gin.Context) {
	pagination := c.Query("pagination")
	searchWord := c.Query("searchword")

	// foo1 := new(Movies)
    // err := getJson("http://www.omdbapi.com/?apikey=faf7e5bb&s=Batman&page=2", foo1)
	// log.Println(foo1.TotalResults)
	// log.Println(foo1.Response)

	// search := foo1.Search
	// for i:=0; i<len(search); i++ {
	// 	log.Print(search[i].([]interface{})["Title"])
	// }
	
	// if err != nil {
	// 	log.Fatal(err)
	// 	c.JSON(404, gin.H{
	// 		"error": 404,
	// 		"message": "not found",
	// 	})
	// }

	page, err := strconv.Atoi(pagination)
	if err != nil {
		log.Fatal(err)
	}

	movies, err := DbQuery(page, searchWord)

	c.JSON(200, gin.H{
		"message": "hai",
		"pagination": pagination,
		"search_word" : searchWord,
		"movies": movies,
	})
}