package lyrics

import (
	"fmt"
)

type SongLyrics struct {
	Lyrics      string `json:"lyrics,omitempty"`
	SearchError string `json:"error,omitempty"`
}

func GetLyrics(b, s string) string {
	url := fmt.Sprintf("https://api.lyrics.ovh/v1/%s/%s", b, s)
	// ADD CODE
}
