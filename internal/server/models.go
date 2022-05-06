package server

// ImgURLCollection -- response model for returning a list of images from a webpage
type ImgURLCollection struct {
	Imgs []string `json:"imgs"`
}