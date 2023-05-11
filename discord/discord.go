package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type NewsData struct {
	Status       string `json:"status"`
	TotalResults int    `json:"totalResults"`
	Results      []struct {
		Title       string      `json:"title"`
		Link        string      `json:"link"`
		Keywords    []string    `json:"keywords"`
		Creator     interface{} `json:"creator"`
		VideoURL    interface{} `json:"video_url"`
		Description string      `json:"description"`
		Content     string      `json:"content"`
		PubDate     string      `json:"pubDate"`
		ImageURL    string      `json:"image_url"`
		SourceID    string      `json:"source_id"`
		Category    []string    `json:"category"`
		Country     []string    `json:"country"`
		Language    string      `json:"language"`
	} `json:"results"`
	NextPage string `json:"nextPage"`
}

const GoogleNewsUrl = "https://newsdata.io/api/1/news?apikey={your-api-key}&country=in,us&language=en&q=" // use your api key
const RequestURL = "{your-webhook-url}"                                                                   // eg: https://discord.com/api/webhooks/########

func NewsCheck(cT string) {
	cloudTechnologie := cT
	log.Println(cloudTechnologie)
	res, err := http.Get(GoogleNewsUrl + cloudTechnologie)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode == 200 {
		var newsData NewsData
		err2 := json.Unmarshal(body, &newsData)

		if err2 != nil {
			log.Fatal(err2)
		}

		log.Println(len(newsData.Results))

		for k := range newsData.Results {
			content := map[string]string{"content": fmt.Sprintf("%s, %s", newsData.Results[k].Title, newsData.Results[k].Link)}
			json_data, err := json.Marshal(content)

			if err != nil {
				log.Fatal(err)
			}

			res, err := http.Post(RequestURL, "application/json", bytes.NewBuffer(json_data))

			if err != nil {
				log.Fatal(err)
			}

			defer res.Body.Close()

			if res.StatusCode == 204 {
				log.Printf("%v\n", res.StatusCode)
			} else {
				log.Printf("%v\n", res.StatusCode)
			}
		}
	} else {
		log.Printf("%v\n", res.StatusCode)
	}
}
