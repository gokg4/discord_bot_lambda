package discord

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
)

const GoogleNewsUrl1 = "https://news.google.com/rss/search?q=when:3h+intitle:" // Searches for the latest articles published in the last 3 hours.
const GoogleNewsUrl2 = "&ceid=IN:en&hl=en-IN&gl=IN"                            // Additional parameters for Language and Country.
const RequestURL = "https://discord.com/api/webhooks/#########"                // Get your discord webhook URL. Refer the Readme.

type NewsData struct {
	Channel struct {
		Item []struct {
			Text    string `xml:",chardata"`
			Title   string `xml:"title"`
			Link    string `xml:"link"`
			PubDate string `xml:"pubDate"`
		} `xml:"item"`
	} `xml:"channel"`
}

func NewsCheck(cT string) {
	cloudTechnology := cT
	log.Println(cloudTechnology)

	res, err := http.Get(GoogleNewsUrl1 + cloudTechnology + GoogleNewsUrl2)
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
		err2 := xml.Unmarshal(body, &newsData)

		if err2 != nil {
			log.Fatal(err2)
		}

		topicLength := len(newsData.Channel.Item)

		log.Println(topicLength)

		if topicLength <= 20 {
			for k := range newsData.Channel.Item {
				content := map[string]string{"content": fmt.Sprintf("%s, %s", newsData.Channel.Item[k].Title, newsData.Channel.Item[k].Link)}
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
			for i := 0; i < 20; i++ {
				content := map[string]string{"content": fmt.Sprintf("%s, %s", newsData.Channel.Item[i].Title, newsData.Channel.Item[i].Link)}
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
		}
	} else {
		log.Printf("%v\n", res.StatusCode)
	}
}
