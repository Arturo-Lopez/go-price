package service

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

var BonanzaPricePaths = [3]PricePath{
	{
		IsoCode:  "USD",
		BuyPath:  "tr:nth-child(1) > td:nth-child(3)",
		SellPath: "tr:nth-child(1) > td:nth-child(5)",
	},
	{
		IsoCode:  "BRS",
		BuyPath:  "tr:nth-child(4) > td:nth-child(3)",
		SellPath: "tr:nth-child(4) > td:nth-child(5)",
	},
	{
		IsoCode:  "ARG",
		BuyPath:  "tr:nth-child(3) > td:nth-child(3)",
		SellPath: "tr:nth-child(3) > td:nth-child(5)",
	},
}

func Bonanza() {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0")
		r.Ctx.Put("timeout", 10*time.Second)
	})

	c.OnHTML("section.flat-tab-services div.content-tab > div:nth-child(1) > div:nth-child(1) table tbody", func(x *colly.HTMLElement) {
		fmt.Println("--- BONANZA ---")

		for _, item := range BonanzaPricePaths {

			buy := GetValue(x.DOM.Find(item.BuyPath).Text())
			sell := GetValue(x.DOM.Find(item.SellPath).Text())

			fmt.Println(item.IsoCode, " : ", buy, " - ", sell)
		}
	})

	c.Visit("https://www.bonanzacambios.com.py/")
}
