package main

import (
	"github.com/MaxAtslega/live-spotify-readme/template"
	"github.com/MaxAtslega/live-spotify-readme/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

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

		data = template.PageData{
			IsPlaying: true,
			Title:     track.Item.Name,
			Cover:     track.Item.Album.Images[0].URL,
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
