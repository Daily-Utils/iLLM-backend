package models

type BodyForPDF struct {
    Data string `json:"data"`
}

type RequestBodyForPDF struct {
    Body BodyForPDF `json:"body"`
	Model string `json:"model"`
}