package voice

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Data struct {
  Mp3Url string `json:"mp3DownloadUrl"`
}

func voiceAPI(contents string) (string, error) {
  apiUrl := "https://api.tts.quest/v1/voicevox/?text="+contents+"&speaker=2"

  resp, err := http.Get(apiUrl)
  if resp == nil {
    fmt.Println("responce is nil")
  }
	if err != nil {
		fmt.Printf("NewRequest err=%s", err.Error())
	}
  var data Data
  err = json.NewDecoder(resp.Body).Decode(&data)
  if err != nil {
      panic(err)
  }

  mp3Url := data.Mp3Url
  

 return mp3Url, nil 
}

// Download mp3 file from URL
func MP3FromURL(url string) error {
  res, err := http.Get(url)
	if err != nil {
    fmt.Println(err)
	}

	file, err := os.Create("./audio/audio.mp3")
	if err != nil {
		// Handle error
    fmt.Println(err)
	}

  _, err = io.Copy(file, res.Body)
  if err != nil {
    fmt.Println(err)
    return err
  }

  return nil
}

