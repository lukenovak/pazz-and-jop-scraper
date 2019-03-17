// simple scraper that works with CSS selectors
package main

import (
	"database/sql"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"time"
)

func main() {
	// constant start and end years
	const start_year int = 71
	const end_year int = 100

	// for each year
	for i := start_year; i < end_year; i++ {
		pageURL := fmt.Sprintf("https://www.robertchristgau.com/xg/pnj/pjres%d.php", i)

		fmt.Printf("Pazz and Jop Poll, '%d\n", i)
		// get the page
		response, err := http.Get(pageURL)
		if err != nil {
			log.Fatal(err)
		}

		// make the response into a goquery doc
		doc, err := goquery.NewDocumentFromReader(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		albumElements := doc.Find("p:nth-child(4) table td:nth-child(2) b")
		albumElements.Each(printAlbumInfo)
		albumElements.Each(putAlbumIntoDatabase)

		err = response.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}
}

// on that postgresql grind
func putAlbumIntoDatabase(i int, selection *goquery.Selection) {
	db, err := sql.Open("postgres", /*put db info here*/)
	if err != nil {
		log.Fatal(err)
	}

	album, artist := getAlbumAndArtistFromSelection(selection)
	stmt, err:= db.Prepare("INSERT INTO albums VALUES ($1, $2);")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(album, artist)
	err = db.Close()
	if err != nil {
		log.Fatal(err)
	}
}

// prints album info given the element containing album and artist
func printAlbumInfo(i int, selection *goquery.Selection) {
	album, artist := getAlbumAndArtistFromSelection(selection)
	fmt.Printf("%s: %s\n", artist, album)
	time.Sleep(10 * time.Millisecond)
}

// helper
func getAlbumAndArtistFromSelection(selection *goquery.Selection) (string, string) {
	artist := selection.Contents().Not("i").Text()
	album := selection.Find("i").Text()
	if len(artist) > 2 {
		artist = artist[:len(artist) - 2]
	} else {
		artist = album
	}
	return album, artist
}

