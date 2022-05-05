package models

// ImgURLResponse -- response model for returning a list of images from a webpage
type ImgURLResponse struct {
	Imgs []string `json:"imgs"`
}
