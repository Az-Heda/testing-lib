package tg

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	uri "net/url"
)

type Bot struct {
	Config   *BotConfig
	Endpoint string
}

type BotConfig struct {
	Token string `json:"-"`
}

func NewBot(config *BotConfig) Bot {
	var b Bot = Bot{
		Config:   config,
		Endpoint: "https://api.telegram.org/bot" + config.Token,
	}
	return b
}

func (b *Bot) sendRequest(method string, url string) ([]byte, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return []byte{}, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()

	return io.ReadAll(res.Body)
}

func (b *Bot) SendMessage(msg MessagePayload) MessageResponse {
	var url = b.Endpoint + "/sendMessage"

	var query uri.Values = uri.Values{}

	query.Add("chat_id", fmt.Sprint(msg.ChatId))
	query.Add("text", msg.Text)
	query.Add("parse_mode", msg.ParseMode)

	url += "?" + query.Encode()

	fmt.Printf("url: %v\n", url)

	var body, err = b.sendRequest(http.MethodGet, url)
	if err != nil {
		panic(err)
	}
	var out MessageResponse
	json.Unmarshal(body, &out)
	// os.WriteFile("tmp.json", body, os.ModePerm)
	return out
}
