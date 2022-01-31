package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type CurrentlyPlaying struct {
	Timestamp int64 `json:"timestamp"`
	Context   struct {
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href string `json:"href"`
		Type string `json:"type"`
		URI  string `json:"uri"`
	} `json:"context"`
	ProgressMs int `json:"progress_ms"`
	Item       struct {
		Album struct {
			AlbumType string `json:"album_type"`
			Artists   []struct {
				ExternalUrls struct {
					Spotify string `json:"spotify"`
				} `json:"external_urls"`
				Href string `json:"href"`
				ID   string `json:"id"`
				Name string `json:"name"`
				Type string `json:"type"`
				URI  string `json:"uri"`
			} `json:"artists"`
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href   string `json:"href"`
			ID     string `json:"id"`
			Images []struct {
				Height int    `json:"height"`
				URL    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"images"`
			Name                 string `json:"name"`
			ReleaseDate          string `json:"release_date"`
			ReleaseDatePrecision string `json:"release_date_precision"`
			TotalTracks          int    `json:"total_tracks"`
			Type                 string `json:"type"`
			URI                  string `json:"uri"`
		} `json:"album"`
		Artists []struct {
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href string `json:"href"`
			ID   string `json:"id"`
			Name string `json:"name"`
			Type string `json:"type"`
			URI  string `json:"uri"`
		} `json:"artists"`
		AvailableMarkets []string `json:"available_markets"`
		DiscNumber       int      `json:"disc_number"`
		DurationMs       int      `json:"duration_ms"`
		Explicit         bool     `json:"explicit"`
		ExternalIds      struct {
			Isrc string `json:"isrc"`
		} `json:"external_ids"`
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href        string `json:"href"`
		ID          string `json:"id"`
		IsLocal     bool   `json:"is_local"`
		Name        string `json:"name"`
		Popularity  int    `json:"popularity"`
		PreviewURL  string `json:"preview_url"`
		TrackNumber int    `json:"track_number"`
		Type        string `json:"type"`
		URI         string `json:"uri"`
	} `json:"item"`
	CurrentlyPlayingType string `json:"currently_playing_type"`
	Actions              struct {
		Disallows struct {
			Resuming bool `json:"resuming"`
		} `json:"disallows"`
	} `json:"actions"`
	IsPlaying bool `json:"is_playing"`
}

func GetAuthorizationToken() string {
	spotifyClientID := os.Getenv("SPOTIFY_CLIENT_ID")
	spotifyClientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")
	spotifyRefreshToken := os.Getenv("SPOTIFY_REFRESH_TOKEN")

	basic := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", spotifyClientID, spotifyClientSecret)))

	url := "https://accounts.spotify.com/api/token"
	method := "POST"

	payload := strings.NewReader("grant_type=refresh_token&refresh_token=" + spotifyRefreshToken)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	req.Header.Add("Authorization", "Basic "+basic)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	var result map[string]string

	json.NewDecoder(res.Body).Decode(&result)

	return "Bearer " + result["access_token"]
}

//GetCurrentlyPlaying get playing track from spotify
func GetCurrentlyPlaying() CurrentlyPlaying {
	var accessToken = GetAuthorizationToken()

	url := "https://api.spotify.com/v1/me/player/currently-playing"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return CurrentlyPlaying{}
	}
	req.Header.Add("Authorization", accessToken)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return CurrentlyPlaying{}
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return CurrentlyPlaying{}
	}
	
	result := CurrentlyPlaying{}

	err2 := json.Unmarshal([]byte(string(body)), &result)
	if err2 != nil {
		fmt.Println(err2)
		return CurrentlyPlaying{}
	}

	return result
}
