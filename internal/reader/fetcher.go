package reader

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	htmltomd "github.com/JohannesKaufmann/html-to-markdown/v2"
)

func FetchArticle(url string) (title string, markdown string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", "", fmt.Errorf("fetching %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("unexpected status %d for %s", resp.StatusCode, url)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", fmt.Errorf("reading response body: %w", err)
	}

	html := string(body)
	title = extractTitle(html)

	md, err := htmltomd.ConvertString(html)
	if err != nil {
		return title, html, fmt.Errorf("converting to markdown: %w", err)
	}

	return title, md, nil
}

func extractTitle(html string) string {
	lower := strings.ToLower(html)
	start := strings.Index(lower, "<title>")
	if start == -1 {
		return "Untitled"
	}
	start += len("<title>")
	end := strings.Index(lower[start:], "</title>")
	if end == -1 {
		return "Untitled"
	}
	title := strings.TrimSpace(html[start : start+end])
	if title == "" {
		return "Untitled"
	}
	return title
}
