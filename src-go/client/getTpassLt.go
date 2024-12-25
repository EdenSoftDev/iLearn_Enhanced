package client

import (
	"fmt"
	"iLearn_Enhanced/model"
	"io"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func GetTpassLt() (string, string, error) {
	resp, err := http.Get("https://cas.jlu.edu.cn/tpass/login")
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Reading response failed:", err)
	}

	doc, err := html.Parse(strings.NewReader(string(body)))
	if err != nil {
		log.Fatal("Encoding HTML failed:", err)
	}

	tokens := findTokens(doc)

	if tokens.Lt == "" {
		return "", "", fmt.Errorf("lt value not found")
	}

	if tokens.Execution == "" {
		return "", "", fmt.Errorf("execution value not found")
	}

	return tokens.Lt, tokens.Execution, nil
}

func findTokens(n *html.Node) model.LtResponse {
	var tokens model.LtResponse

	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "input" {
			var id, name, value string
			for _, attr := range n.Attr {
				switch attr.Key {
				case "id":
					id = attr.Val
				case "name":
					name = attr.Val
				case "value":
					value = attr.Val
				}
			}

			if id == "lt" && name == "lt" && value != "" {
				tokens.Lt = value
			}

			if name == "execution" && value != "" {
				tokens.Execution = value
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}

	traverse(n)
	return tokens
}
