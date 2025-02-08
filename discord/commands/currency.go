package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/hde-oliv/ayumu/utils"
)

type LatestResponse struct {
	Meta     interface{}        `json:"meta"`
	Response interface{}        `json:"response"`
	Date     interface{}        `json:"date"`
	Base     interface{}        `json:"base"`
	Rates    map[string]float64 `json:"rates"`
}

func Dollar(s *discordgo.Session, i *discordgo.InteractionCreate) {
	rate := getDollarRate()
	text := fmt.Sprintf("ドルは%.2fです！", rate)

	s.InteractionRespond(i.Interaction, textResponse(text))
}

func getDollarRate() float64 {
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
