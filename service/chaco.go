package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ChacoResponse struct {
	UpdatedTs      string      `json:"updatedTs"`
	BranchOfficeId string      `json:"branchOfficeId"`
	Items          []ChacoItem `json:"items"`
}

type ChacoItem struct {
	IsoCode        string `json:"isoCode"`
	PurchasePrice  int32  `json:"purchasePrice"`
	SalePrice      int32  `json:"salePrice"`
	FractionDigits int8   `json:"fractionDigits"`
	UpdatedAt      string `json:"updatedAt"`
}

var ChacoCodes = map[string]int{
	"USD": 1,
	"BRL": 1,
	"ARS": 1,
}

func Chaco() {
	resp, err := http.Get("http://www.cambioschaco.com.py/api/branch_office/1/exchange")

	if err != nil {
		fmt.Printf(err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf(err.Error())
	}

	var (
		chachoResp ChacoResponse
		item       ChacoItem
	)

	json.Unmarshal(body, &chachoResp)

	fmt.Println("--- CHACO ---")

	for _, item = range chachoResp.Items {
		_, exists := ChacoCodes[item.IsoCode]

		if exists {
			fmt.Println(item.IsoCode, " : ", item.PurchasePrice, " - ", item.SalePrice)
		}
	}
}
