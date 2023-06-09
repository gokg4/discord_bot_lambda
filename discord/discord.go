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

type NewsData struct {
	XMLName xml.Name `xml:"rss"`
	Text    string   `xml:",chardata"`
	Media   string   `xml:"media,attr"`
	Version string   `xml:"version,attr"`
	Channel struct {
		Text          string `xml:",chardata"`
		Generator     string `xml:"generator"`
		Title         string `xml:"title"`
		Link          string `xml:"link"`
		Language      string `xml:"language"`
		WebMaster     string `xml:"webMaster"`
		Copyright     string `xml:"copyright"`
		LastBuildDate string `xml:"lastBuildDate"`
		Description   string `xml:"description"`
		Item          []struct {
			Text  string `xml:",chardata"`
			Title string `xml:"title"`
			Link  string `xml:"link"`
			Guid  struct {
				Text        string `xml:",chardata"`
				IsPermaLink string `xml:"isPermaLink,attr"`
			} `xml:"guid"`
			PubDate     string `xml:"pubDate"`
			Description struct {
				Text string `xml:",chardata"`
				A    struct {
					Text   string `xml:",chardata"`
					Href   string `xml:"href,attr"`
					Target string `xml:"target,attr"`
				} `xml:"a"`
				Font struct {
					Text  string `xml:",chardata"`
					Color string `xml:"color,attr"`
				} `xml:"font"`
			} `xml:"description"`
			Source struct {
				Text string `xml:",chardata"`
				URL  string `xml:"url,attr"`
			} `xml:"source"`
		} `xml:"item"`
	} `xml:"channel"`
}

const GoogleNewsUrl1 = "https://news.google.com/rss/search?q=when:3h+intitle:" // Searches for the latest articles published in the last 3 hours.
const GoogleNewsUrl2 = "&ceid=IN:en&hl=en-IN&gl=IN"                            // Additional parameters for Language and Country.
const RequestURL = "https://discord.com/api/webhooks/#########"                // Get your discord webhook URL. Refer the Readme.

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

		log.Println(len(newsData.Channel.Item))

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
		log.Printf("%v\n", res.StatusCode)
	}
}
