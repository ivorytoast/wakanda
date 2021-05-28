package config

import "time"

type SetupConfig struct {
	LiveBaseUrl string
	LiveDataUrl string
	LiveApiVersion string
	LiveClientRequestTimeout time.Duration

	PaperBaseUrl string
	PaperDataUrl string
	PaperApiVersion string
	PaperClientRequestTimeout time.Duration
}