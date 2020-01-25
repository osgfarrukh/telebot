package telebot

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Bot struct {
	Token       string
	WebhookUrl  string
	ApiEndpoint string
	Client      *http.Client
}

type UpdateChan chan *WebhookUpdate

func NewBot(token string) *Bot {
	return &Bot{
		Token:       token,
		ApiEndpoint: fmt.Sprintf("https://api.telegram.org/bot%s/", token),
		Client:      &http.Client{},
	}
}

func (b Bot) SetWebhook(domen string) error {
	v := url.Values{}
	v.Add("url", domen)
	_, err := b.makeRequest("setWebhook", v)
	if err != nil {
		return err
	}
	return nil
}

func (b *Bot) WebhookUpdates(mux *mux.Router) UpdateChan {
	updateChan := make(chan *WebhookUpdate, 0)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		if cT := r.Header.Get("Content-Type"); cT != "application/json" {
			return
		}
		p, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("%+v", err)
		}
		var update WebhookUpdate
		err = json.Unmarshal(p, &update)
		if err != nil {
			log.Printf("%+v", err)
		}
		updateChan <- &update
	})
	return updateChan
}

func (b *Bot) makeRequest(method string, values url.Values) (*APIResponse, error) {
	response, err := b.Client.PostForm(b.ApiEndpoint+method, values)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var apiRes APIResponse
	err = json.Unmarshal(data, &apiRes)
	if err != nil {
		return nil, err
	}
	if !apiRes.Ok {
		return nil, fmt.Errorf(apiRes.Description)
	}
	return &apiRes, nil
}

func (b *Bot) SendMessage(values url.Values) error {
	_, err := b.makeRequest("sendMessage", values)
	if err != nil {
		return err
	}
	return nil
}

func (b *Bot) DeleteMessage(values url.Values) error {
	_, err := b.makeRequest("deleteMessage", values)
	if err != nil {
		return err
	}
	return nil
}

func (b *Bot) EditMessageReplyMarkup(values url.Values) error {
	_, err := b.makeRequest("editMessageReplyMarkup", values)
	if err != nil {
		return err
	}
	return nil
}

func (b *Bot) SendPhoto(values url.Values) error {
	_, err := b.makeRequest("sendPhoto", values)
	if err != nil {
		return err
	}
	return nil
}

func (b *Bot) FileUrl(fileId string) (string, error) {
	v := url.Values{}
	v.Add("file_id", fileId)
	res, err := b.makeRequest("getFile", v)
	if err != nil {
		return "", err
	}
	var f File
	err = json.Unmarshal(res.Result, &f)
	if err != nil {
		return "", nil
	}
	return fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", b.Token, f.FilePath), nil
}
