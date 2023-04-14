package goimgbb

import (
	"encoding/json"
	"io"

	"net/http"
	"net/url"
)

type ImgbbImage struct {
	Filename  string `json:"filename"`
	Name      string `json:"name"`
	Mime      string `json:"mime"`
	Extension string `json:"extension"`
	Url       string `json:"url"`
}

type ImgbbResponse struct {
	Data    ImgbbResponseData `json:"data"`
	Success bool              `json:"success"`
	Status  int               `json:"status"`
}

type ImgbbResponseData struct {
	ID string `json:"id"`

	DisplayURL string `json:"display_url"`
	DeleteURL  string `json:"delete_url"`

	Expiration int `json:"expiration"`

	Height int `json:"height"`
	Width  int `json:"width"`

	Image ImgbbImage `json:"image"`
	Thumb ImgbbImage `json:"thumb"`
}

/*
Uploads image to imgbb.com

@param imageBBkey string: It's your imgBB API key

@param base64Image string: It's the stringified as base64 image (without "data:**;base64,")
*/
func Upload(imageBBkey string, base64Image string) (*ImgbbResponse, error) {
	// Make request to ImgBB with Key and a Base64 image
	resp, err := http.PostForm("https://api.imgbb.com/1/upload?key="+imageBBkey,
		url.Values(map[string][]string{
			"image": {base64Image},
		}),
	)
	if err != nil {
		return nil, err
	}

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Transform response body into type
	imgbb_response := &ImgbbResponse{}

	err = json.Unmarshal(body, imgbb_response)
	if err != nil {
		return nil, err
	}

	// Close body
	defer resp.Body.Close()

	return imgbb_response, nil
}
