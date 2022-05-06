package scraper

import "archive/zip"

type ZipInfoGetter interface {
	getFileName() string
	getImgUrl() string
	getZipWriter() *zip.Writer
}