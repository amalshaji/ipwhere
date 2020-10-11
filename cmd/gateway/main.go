package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
)

// func main() {
// 	port := flag.Int("port", -1, "specify a port to use http rather than AWS Lambda")
// 	flag.Parse()
// 	listener := gateway.ListenAndServe
// 	portStr := ""
// 	if *port != -1 {
// 		portStr = fmt.Sprintf(":%d", *port)
// 		listener = http.ListenAndServe
// 		http.Handle("/", http.FileServer(http.Dir("./public")))
// 	}

// 	http.Handle("/api/feed", feed2json.Handler(
// 		feed2json.StaticURLInjector("https://news.ycombinator.com/rss"),
// 		nil, nil, nil, cacheControlMiddleware))
// 	log.Fatal(listener(portStr, nil))
// }
type response struct {
	Status    string  `json:"status"`
	Country   string  `json:"country"`
	Region    string  `json:"regionName"`
	Latitude  float32 `json:"lat"`
	Longitude float32 `json:"lon"`
	ISP       string  `json:"isp"`
}

func cacheRequest(exp time.Duration) fiber.Handler {
	return cache.New(cache.Config{
		Expiration:   exp,
		CacheControl: true,
	})
}

func getGeoLocation(c *fiber.Ctx) error {
	ip := c.Params("ip")
	res, _ := http.Get("http://ip-api.com/json/" + ip)
	body, _ := ioutil.ReadAll(res.Body)

	var resp response
	json.Unmarshal(body, &resp)
	if resp.Status == "fail" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "enter an ip",
		})
	}
	return c.Status(fiber.StatusOK).JSON(resp)
}

func main() {
	app := fiber.New()
	app.Static("/", "./public", fiber.Static{
		Compress: true,
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("index")
	})
	app.Get("/api/:ip", cacheRequest(10*time.Minute), getGeoLocation)
	log.Fatal(app.Listen(":3000"))
}
