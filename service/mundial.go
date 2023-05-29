package service

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

var MundialPricePaths = [3]PricePath{
	{
		IsoCode:  "USD",
		BuyPath:  "div:nth-child(2) > div:nth-child(2) > h3",
		SellPath: "div:nth-child(2) > div:nth-child(3) > h3",
	},
	{
		IsoCode:  "BRS",
		BuyPath:  "div:nth-child(3) > div:nth-child(2) > h3",
		SellPath: "div:nth-child(3) > div:nth-child(3) > h3",
	},
	{
		IsoCode:  "ARG",
		BuyPath:  "div:nth-child(4) > div:nth-child(2) > h3",
		SellPath: "div:nth-child(4) > div:nth-child(3) > h3",
	},
}

func Mundial() {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36")
		r.Ctx.Put("timeout", 20*time.Second)
	})

	c.OnHTML("#w-node-1a5c1e9ec89d-d4665980", func(x *colly.HTMLElement) {
		fmt.Println("--- MUNDIAL ---")

		for _, item := range MundialPricePaths {

			buy := GetValue(x.DOM.Find(item.BuyPath).Text())
			sell := GetValue(x.DOM.Find(item.SellPath).Text())

			fmt.Println(item.IsoCode, " : ", buy, " - ", sell)
		}
	})

	c.Visit("http://www.mundialcambios.com.py/?branch=6")
}
