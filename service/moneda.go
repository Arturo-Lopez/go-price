package service

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

var MonedaPricePaths = [3]PricePath{
	{
		IsoCode:  "USD",
		BuyPath:  "tr:nth-child(1) > td:nth-child(3)",
		SellPath: "tr:nth-child(1) > td:nth-child(4)",
	},
	{
		IsoCode:  "BRS",
		BuyPath:  "tr:nth-child(2) > td:nth-child(3)",
		SellPath: "tr:nth-child(2) > td:nth-child(4)",
	},
	{
		IsoCode:  "ARG",
		BuyPath:  "tr:nth-child(3) > td:nth-child(3)",
		SellPath: "tr:nth-child(3) > td:nth-child(4)",
	},
}

func Moneda() {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0")
		r.Ctx.Put("timeout", 10*time.Second)
	})

	c.OnHTML("#about > div:nth-child(2) > div > div:nth-child(1) > div > table > tbody", func(x *colly.HTMLElement) {
		fmt.Println("--- MONEDA ---")

		for _, item := range MonedaPricePaths {

			buy := GetValue(x.DOM.Find(item.BuyPath).Text())
			sell := GetValue(x.DOM.Find(item.SellPath).Text())

			fmt.Println(item.IsoCode, " : ", buy, " - ", sell)
		}
	})

	c.Visit("http://www.lamoneda.com.py/")
}
