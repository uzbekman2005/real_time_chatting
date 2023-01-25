package models

type GetYoutubeLinksResponse struct {
	Id         int `json:"id"`
	VideoTitle string `json:"video_title"`
	VideoURL   string `json:"video_url"`
	VideoType  string `json:"video_type"`
}
