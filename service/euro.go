package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type EuroCambiosResponse struct {
	Id     string `json:"id"`
	Fecha  string `json:"fecha"`
	Compra string `json:"compra"`
	Venta  string `json:"venta"`
}

var EuroCambiosCodes = map[string]int{
	"US": 1,
	"PA": 1,
	"RS": 1,
}

func Euro() {
	data := url.Values{}

	data.Set("param", "getCotizacionesbySucursal")
	data.Set("sucursal", "1")

	resp, err := http.PostForm("https://eurocambios.com.py/v2/sgi/utilsDto.php", data)

	if err != nil {
		fmt.Printf(err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf(err.Error())
	}

	var (
		euroResp []EuroCambiosResponse
		item     EuroCambiosResponse
	)

	json.Unmarshal(body, &euroResp)

	fmt.Println("--- EURO ---")

	for _, item = range euroResp {
		_, exists := EuroCambiosCodes[item.Id]

		if exists {
			fmt.Println(item.Id, " : ", item.Compra, " - ", item.Venta)
		}
	}
}
