package service

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

var VisionPricePaths = [3]PricePath{
	{
		IsoCode:  "USD",
		BuyPath:  "tr:nth-child(1) > td:nth-child(2) > p:nth-child(1)",
		SellPath: "tr:nth-child(1) > td:nth-child(3) > p:nth-child(1)",
	},
	{
		IsoCode:  "BRS",
		BuyPath:  "tr:nth-child(2) > td:nth-child(2) > p:nth-child(1)",
		SellPath: "tr:nth-child(2) > td:nth-child(3) > p:nth-child(1)",
	},
	{
		IsoCode:  "ARG",
		BuyPath:  "tr:nth-child(3) > td:nth-child(2) > p:nth-child(1)",
		SellPath: "tr:nth-child(3) > td:nth-child(3) > p:nth-child(1)",
	},
}

func Vision() {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0")
		r.Ctx.Put("timeout", 10*time.Second)
	})

	c.OnHTML("#efectivo table tbody", func(x *colly.HTMLElement) {
		fmt.Println("--- VISION ---")

		for _, item := range VisionPricePaths {

			buy := GetValue(x.DOM.Find(item.BuyPath).Text())
			sell := GetValue(x.DOM.Find(item.SellPath).Text())

			fmt.Println(item.IsoCode, " : ", buy, " - ", sell)
		}
	})

	c.Visit("https://www.visionbanco.com")
}
