package discord

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/hde-oliv/ayumu/utils"
)

type LatestResponse struct {
	Meta     interface{}        `json:"meta"`
	Response interface{}        `json:"response"`
	Date     interface{}        `json:"date"`
	Base     interface{}        `json:"base"`
	Rates    map[string]float64 `json:"rates"`
}

func GetDollarRate() float64 {
	base_addr := "https://api.currencybeacon.com/v1/latest"
	api_key := "?api_key=" + os.Getenv("CURRENCY_TOKEN")

	resp, err := http.Get(base_addr + api_key)
	utils.Check(err)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	utils.Check(err)

	var cResp LatestResponse

	err = json.Unmarshal(body, &cResp)
	utils.Check(err)

	return cResp.Rates["BRL"]
}
