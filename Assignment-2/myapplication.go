package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sellerapp/base/utils"
	"strings"

	"github.com/gocolly/colly"
)

type AmazonProductDetails struct {
	Id        string   `param:"id" json:"id"`
	Name      string   `param:"name" json:"name"`
	ImagesURL []string `param:"images_url" json:"images_url"`
	Desc      string   `param:"desc" json:"desc"`
	Price     string   `param:"price" json:"price"`
	Reviews   string   `param:"reviews" json:"reviews"`
}

type pageInfo struct {
	StatusCode int                  `param:"status_code" json:"status_code"`
	Message    string               `param:"message" json:"message"`
	Data       AmazonProductDetails `param:"data" json:"data"`
}

func GetAmazonProductDetails(url, id string) (p pageInfo) {
	c := colly.NewCollector()

	// title
	c.OnHTML("title", func(e *colly.HTMLElement) {
		//log.Println("Name:", e.Text, e.Attr("title"))
		p.Data.Name = e.Text
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.HasSuffix(link, "jpg") || strings.HasSuffix(link, "png") {
			p.Data.ImagesURL = append(p.Data.ImagesURL, link)
		}
	})

	c.OnHTML(`div[id=feature-bullets]`, func(e *colly.HTMLElement) {
		//log.Println("Desc:", strings.Trim(e.Text, "'\t', '\n', '\v', '\f', '\r', ' '"))
		e.ForEach("ul", func(_ int, el *colly.HTMLElement) {
			//log.Println("Desc:", el.ChildText("li"))
			p.Data.Desc = el.ChildText("li")
		})
	})

	c.OnHTML(`span[id=priceblock_ourprice]`, func(e *colly.HTMLElement) {
		//log.Println("Price:", e.Text)
		p.Data.Price = e.Text
	})

	// reviews
	c.OnHTML(`a[id=acrCustomerReviewLink]`, func(e *colly.HTMLElement) {
		//log.Println("Review:", strings.TrimSpace(e.Text))
		p.Data.Reviews = strings.TrimSpace(e.Text)
	})

	p.Data.Id = id

	// extract status code
	c.OnResponse(func(r *colly.Response) {
		log.Println("response received", r.StatusCode)
		p.StatusCode = r.StatusCode
	})
	c.OnError(func(r *colly.Response, err error) {
		log.Println("error:", r.StatusCode, err)
		p.StatusCode = r.StatusCode
	})

	c.Visit(url)

	return
}

// init
func Init() {
	setupRouter()
}

func main() {
	Init()
}

func setupRouter() {
	http.HandleFunc("/product/amazon", GetProductDetails)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func GetProductDetails(w http.ResponseWriter, r *http.Request) {
	pageDetails := GetAmazonProductDetails("https://www.amazon.com/dp/B00D3JA5V4/", "B00D3JA5V4")
	jsonReq, _ := json.Marshal(pageDetails.Data)

	_, err := utils.ApiCall("POST", "http://:50051/sellerapp/v1/product/amazon", jsonReq)
	if err == nil {
		utils.Json("{\"success\":true, \"data\":{\"message\":\"success.\"}}")(w, r)
	} else {
		utils.Json("{\"success\":false, \"data\":{\"message\":\"Something went wrong.\"}}")(w, r)
	}
}
