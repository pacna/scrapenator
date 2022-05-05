package server

import (
	"archive/zip"
	"os"
)

// ImgURLResponse -- response model for returning a list of images from a webpage
type ImgURLResponse struct {
	Imgs []string `json:"imgs"`
}

// ZipInfo -- holds necessary info to successfully create a zip file
type ZipInfo struct {
	ZipFile   *os.File
	ZipWriter *zip.Writer
	FileName  string
	ImgURL    string
}
