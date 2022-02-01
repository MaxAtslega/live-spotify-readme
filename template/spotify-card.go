package template

import (
	"html/template"
	"math"
)

type PageData struct {
	IsPlaying bool
	Title     string
	Artist    string
	Cover     string
	Progress  float64
}

// GenerateSpotifyCard generate Spotify card
func GenerateSpotifyCard() *template.Template {
	templ := template.Must(template.ParseFiles("template/card.svg"))

	return templ
}

func CalcProgressBar(progressMs float64, durationMs float64) float64 {
	return math.Ceil((progressMs / durationMs * 100 * (600 - 140)) / 100)
}
