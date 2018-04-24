package lyrics

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type SongLyrics struct {
	Lyrics      string `json:"lyrics,omitempty"`
	SearchError string `json:"error,omitempty"`
	parseError  error
	getError	error
}

func (sl *SongLyrics) Populate(b, s string) {
	url := fmt.Sprintf("https://api.lyrics.ovh/v1/%s/%s", b, s)
	resp, err := http.Get(url)
	if err != nil {
		sl.getError = err
	}
	defer resp.Body.Close()	

	sl.parseError = json.NewDecoder(resp.Body).Decode(sl)	
}

func (sl SongLyrics) String() string {
	if sl.SearchError != "" {
		return sl.SearchError
	}

	return sl.Lyrics
}

func GetLyrics(b, s string) string {
	var song SongLyrics

	song.Populate(b, s)
	return fmt.Sprint(song)

}

func StandAlone_GetLyrics(b, s string) string {
	noLyrics := "No lyrics found"

	url := fmt.Sprintf("https://api.lyrics.ovh/v1/%s/%s", b, s)
	resp, err := http.Get(url)
	if err != nil {
		return noLyrics
	}
	defer resp.Body.Close()

	var song SongLyrics
	err = json.NewDecoder(resp.Body).Decode(&song)

	if err != nil {
		return noLyrics
	}

	if song.SearchError != "" {
		return song.SearchError
	}

	return song.Lyrics
}
