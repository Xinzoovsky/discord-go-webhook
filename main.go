package discordGoWebhook

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type Thumbnail struct {
	URL string `json:"url"`
}

type Fields struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

type Footer struct {
	Text    string `json:"text"`
	IconURL string `json:"icon_url"`
}

type Embed struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	Color       int       `json:"color"`
	Timestamp   string    `json:"timestamp"`
	Thumbnail   Thumbnail `json:"thumbnail"`
	Fields      []Fields  `json:"fields"`
	Footer      Footer    `json:"footer"`
}

type Webhook struct {
	Username  string  `json:"username"`
	AvatarURL string  `json:"avatar_url"`
	Embeds    []Embed `json:"embeds"`
}

func CreateWebhook() Webhook {
	w := Webhook{
		Username:  "",
		AvatarURL: "",
		Embeds: []Embed{
			{
				Title:     "",
				URL:       "",
				Color:     000000,
				Thumbnail: Thumbnail{URL: ""},
				Fields:    []Fields{},
			},
		},
	}

	return w
}

func (w *Webhook) SetWebhookUsername(username string) {
	w.Username = username
}

func (w *Webhook) SetURL(URL string) {
	w.Embeds[0].URL = URL
}

func (w *Webhook) SetWebhookAvatarURL(avatarURL string) {
	w.AvatarURL = avatarURL
}

func (w *Webhook) SetColor(color int) {
	w.Embeds[0].Color = color
}

func (w *Webhook) SetTitle(title, description string) {
	w.Embeds[0].Title = title
	if description != "" {
		w.Embeds[0].Description = description
	}
}

func (w *Webhook) SetThumbnailURL(thumbnailURL string) {
	w.Embeds[0].Thumbnail.URL = thumbnailURL
}

func (w *Webhook) AddField(title string, value string, inline bool) {

	newField := Fields{
		Name:   title,
		Value:  value,
		Inline: inline,
	}

	w.Embeds[0].Fields = append(w.Embeds[0].Fields, newField)

}

func (w *Webhook) AddFooter(text string, iconURL string) {
	w.Embeds[0].Footer = Footer{
		Text:    text,
		IconURL: iconURL,
	}
}

func (w Webhook) SendWebhook(url string) (*http.Response, error) {
	client := &http.Client{}

	webhookData, err := json.Marshal(w)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(webhookData))

	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}

	webhookPost, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	switch webhookPost.StatusCode {
	case 204:
		return webhookPost, nil
	case 429:
		return webhookPost, errors.New("too many requests")
	default:
		return webhookPost, errors.New("invalid status code")
	}
}
