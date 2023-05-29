package service

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

var MaxiPricePaths = [3]PricePath{
	{
		IsoCode:  "USD",
		BuyPath:  "tr:nth-child(1) > td:nth-child(2)",
		SellPath: "tr:nth-child(1) > td:nth-child(3)",
	},
	{
		IsoCode:  "ARG",
		BuyPath:  "tr:nth-child(2) > td:nth-child(2)",
		SellPath: "tr:nth-child(2) > td:nth-child(3)",
	},
	{
		IsoCode:  "BRS",
		BuyPath:  "tr:nth-child(3) > td:nth-child(2)",
		SellPath: "tr:nth-child(3) > td:nth-child(3)",
	},
}

func Maxi() {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0")
		r.Ctx.Put("timeout", 10*time.Second)
	})

	c.OnHTML(".fixed-plugin table tbody", func(x *colly.HTMLElement) {
		fmt.Println("--- MAXI ---")

		for _, item := range MaxiPricePaths {

			buy := x.DOM.Find(item.BuyPath).Text()
			sell := x.DOM.Find(item.SellPath).Text()

			fmt.Println(item.IsoCode, " : ", buy, " - ", sell)
		}
	})

	c.Visit("https://www.maxicambios.com.py/")
}
