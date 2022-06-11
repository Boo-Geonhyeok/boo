package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"rest/models"
)

var song models.Song
var songList []models.Song

var songData models.SongJSON
var songDataList []models.SongJSON

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := (template.ParseGlob("templates/*.html"))
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(w, "index.html", nil)
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Header.Get("Content-Type") {

	case "application/json":
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		decoder.Decode(&songData)
		songData.VideoID = ""
		songDataList = append(songDataList, songData)

		for index, song := range songDataList {
			query := song.Title + song.Artist
			searchYoutube(query[:len(query)-1])
			videoID := SearchedYoutubeStruct.Items[0].ID.VideoID
			songDataList[index].VideoID = videoID
		}
		songData, err := json.Marshal(songDataList)
		if err != nil {
			fmt.Println(err)
		}

		_, err = http.Post("http://127.0.0.1:8080/1/", "application/json", bytes.NewBuffer(songData))
		if err != nil {
			fmt.Println(err)
		}

	default:
		songTitle := r.FormValue("songTitle")

		searchSpotify(songTitle)

		for _, data := range SearchedSpotifyStruct.Tracks.Items {
			song.ID = data.ID
			song.Title = data.Name
			song.Artist = ""
			song.SpotifyURL = "https://open.spotify.com/embed/track/" + song.ID + "?utm_source=generator"
			for _, artist := range data.Artists {
				if len(song.Artist) != 0 {
					song.Artist += ", "
				}
				song.Artist += artist.Name
			}
			songList = append(songList, song)
		}

		t, err := template.ParseGlob("templates/*.html")

		if err != nil {
			panic(err)
		}

		t.ExecuteTemplate(w, "search.html", models.SongStruct{Searched: songTitle, SongData: songList})
		songList = nil
	}
}

func PlaylistHandler(w http.ResponseWriter, r *http.Request) {

	// for index, song := range songDataList {
	// 	query := song.Title + song.Artist
	// 	searchYoutube(query[:len(query)-1])
	// 	videoID := SearchedYoutubeStruct.Items[0].ID.VideoID
	// 	songDataList[index].VideoID = videoID
	// }
	// songData, err := json.Marshal(songDataList)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// _, err = http.Post("http://127.0.0.1:8080/1/", "application/json", bytes.NewBuffer(songData))
	// if err != nil {
	// 	fmt.Println(err)

	t, err := template.ParseGlob("templates/*.html")
	if err != nil {
		panic(err)
	}

	t.ExecuteTemplate(w, "playlist.html", nil)
	//models.SongJSONStruct{SongJSONSlice: songDataList}
}
