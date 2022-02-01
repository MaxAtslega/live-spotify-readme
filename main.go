package main

import (
	"encoding/base64"
	"github.com/MaxAtslega/live-spotify-readme/template"
	"github.com/MaxAtslega/live-spotify-readme/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func getSpotify(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	w.Header().Set("Cache-Control", "s-maxage=1, stale-while-revalidate")

	track := utils.GetCurrentlyPlaying()

	data := template.PageData{
		IsPlaying: false,
		Title:     "Nothing playing...",
		Cover:     "",
		Artist:    "",
		Progress:  template.CalcProgressBar(0, 0),
	}

	if track.IsPlaying && track.Item.URI != "" {
		artist := ""

		for index, element := range track.Item.Artists {
			if index != (len(track.Item.Artists) - 1) {
				artist += element.Name + ", "
			} else {
				artist += element.Name
			}

		}

		res, err := http.Get(track.Item.Album.Images[0].URL)

		if err != nil {
			log.Printf("http.Get -> %v", err)
			return
		}

		// We read all the bytes of the image
		// Types: data []byte
		bytes, err := ioutil.ReadAll(res.Body)

		if err != nil {
			log.Println("ioutil.ReadAll -> %v", err)
			return
		}
		// Append the base64 encoded output

		data = template.PageData{
			IsPlaying: true,
			Title:     track.Item.Name,
			Cover:     toBase64(bytes),
			Artist:    artist,
			Progress:  template.CalcProgressBar(float64(track.ProgressMs), float64(track.Item.DurationMs)),
		}
	}

	templ := template.GenerateSpotifyCard()
	err := templ.Execute(w, data)
	if err != nil {
		return
	}

}

func handleRequests(port string) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/spotify", getSpotify)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	port := os.Getenv("PORT")

	handleRequests(port)
}
