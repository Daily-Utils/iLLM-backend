package utils

import (
	"encoding/json"
	"fmt"

	"github.com/gocolly/colly"
)

func GetTextFromLink(url string, domain []string) (string, error) {
	textArr := []string{}
	text := ""

	c := colly.NewCollector(
		colly.AllowedDomains(domain...),
	)

	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36"

	c.OnHTML("p", func(e *colly.HTMLElement) {
		text = e.Text
		textArr = append(textArr, text)
		text = ""
	})

	err := c.Visit(url)
	if err != nil {
		return "", fmt.Errorf("failed to visit url: %v", err)
	}

	jsonArr, err := json.Marshal(textArr)
	if err != nil {
		return "", fmt.Errorf("failed to marshal text: %v", err)
	}

	return string(jsonArr), nil
}
