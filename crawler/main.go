package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

func main() {
	resp, err := http.Get("http://www.yahoo.com.tw")
	if err != nil {
		fmt.Printf("http get error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http read error:", err)
		return
	}

	src := string(body)

	// Precompiled regex patterns
	tagPattern := regexp.MustCompile(`\<[^\>]+\>`)
	stylePattern := regexp.MustCompile(`\<style[^\<]*\</style\>`)
	scriptPattern := regexp.MustCompile(`\<script[^\<]*\</script\>`)

	// Convert HTML tags to lowercase
	src = tagPattern.ReplaceAllStringFunc(src, strings.ToLower)

	// Remove <style> and <script> sections
	src = stylePattern.ReplaceAllString(src, "")
	src = scriptPattern.ReplaceAllString(src, "")

	// Remove all remaining HTML tags and replace with newline
	src = tagPattern.ReplaceAllString(src, "\n")

	// Normalize whitespace
	src = regexp.MustCompile(`\s+`).ReplaceAllString(src, "\n")

	fmt.Println(strings.TrimSpace(src))
}
