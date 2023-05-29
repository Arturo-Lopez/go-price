package service

import "regexp"

type PricePath struct {
	IsoCode  string
	BuyPath  string
	SellPath string
}

func GetValue(text string) string {
	pattern := `\d[\d.,]*`

	re := regexp.MustCompile(pattern)
	match := re.FindString(text)

	match = regexp.MustCompile(`\.`).ReplaceAllString(match, "")
	match = regexp.MustCompile(`,`).ReplaceAllString(match, ".")

	return match
}
