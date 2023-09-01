package colly

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type item struct {
	Name   string `json."name"`
	Price  string `json."price"`
	ImgUrl string `json:"imgurl"`
}

var snowboard = "Orca"

func ScrapeAndBake() {
	c := colly.NewCollector()
	c.SetRequestTimeout(120 * time.Second)
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Lets take a look at", r.URL, "       =D")
	})

	var items []item

	// This code will run when an element is found. It will allow us to access the text from the web elements.
	c.OnHTML("div.results-products div[class=product-thumb-details]", func(h *colly.HTMLElement) {
		if strings.Contains(h.ChildText("span.product-thumb-title"), snowboard) {
			item := item{
				Name:   h.ChildText("span.product-thumb-title"),
				Price:  h.ChildText("span.product-thumb-price"),
				ImgUrl: h.ChildAttr("img", "src"),
			}
			items = append(items, item)
		}
	})

	// This function will output the URL name each time its hit. This was used to test next buttons and multiple web pages.
	c.OnRequest(func(r *colly.Request) {
		fmt.Println(r.URL.String())
	})

	// Outputs error message if encountered.
	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("ERROR:", e)
	})

	c.Visit("https://www.evo.com/shop/snowboard/snowboards/rpp_400")
	c.Visit("https://www.evo.com/shop/snowboard/snowboards/p_2/rpp_400")
	fmt.Println(items)

	// This generates a json output file that contains our data.
	content, err := json.Marshal(items)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Outputs the json file and defines permissions to the file.
	os.WriteFile("cool-products.json", content, 0644)
}
