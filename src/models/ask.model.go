package models

type Ask struct {
    Model  string `json:"model"`
    Prompt string `json:"prompt"`
    Stream bool   `json:"stream"`
}
