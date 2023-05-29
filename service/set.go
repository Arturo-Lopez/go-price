package service

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

var SetPricePaths = [3]PricePath{
	{
		IsoCode:  "USD",
		BuyPath:  "/div[1]/div[2]",
		SellPath: "/div[1]/div[3]",
	},
	{
		IsoCode:  "ARG",
		BuyPath:  "/div[5]/div[2]",
		SellPath: "/div[5]/div[3]",
	},
	{
		IsoCode:  "BRS",
		BuyPath:  "/div[9]/div[2]",
		SellPath: "/div[9]/div[3]",
	},
}

func Set() {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0")
		r.Ctx.Put("timeout", 10*time.Second)
	})

	c.OnXML("//*[@id='main']", func(x *colly.XMLElement) {
		fmt.Println("--- SET ---")

		for _, item := range SetPricePaths {

			buy := GetValue(x.ChildText(item.BuyPath))
			sell := GetValue(x.ChildText(item.SellPath))

			fmt.Println(item.IsoCode, " : ", buy, " - ", sell)
		}
	})

	c.Visit("https://www.set.gov.py/portal/PARAGUAY-SET")
}
