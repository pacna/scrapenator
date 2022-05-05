package models

import (
	"archive/zip"
	"os"
)

// ZipInfo -- holds necessary info to successfully create a zip file
type ZipInfo struct {
	ZipFile   *os.File
	ZipWriter *zip.Writer
	FileName  string
	ImgURL    string
}
