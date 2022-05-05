package scraper

import "archive/zip"

type ZipInfoReceiver struct {
	fileName string
	imgURL string
	zipWriter *zip.Writer
}

func (zipInfoReceiver ZipInfoReceiver) getFileName() string {
	return zipInfoReceiver.fileName
}

func (zipInfoReceiver ZipInfoReceiver) getImgUrl() string {
	return zipInfoReceiver.imgURL
}

func (zipInfoReceiver ZipInfoReceiver) getZipWriter() *zip.Writer {
	return zipInfoReceiver.zipWriter
}