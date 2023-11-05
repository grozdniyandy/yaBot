package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"io"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	customTextDone := make(chan bool)

	go displayCustomText(customTextDone)

	<-customTextDone

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/cap" {
			http.NotFound(w, r)
			return
		}

		anchorURL := r.URL.Query().Get("url")

		key, token, err := extractKeyAndToken(anchorURL)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		solution, err := getSolution(key, token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result := struct {
			Key      string `json:"key"`
			Token    string `json:"token"`
			Solution string `json:"solution"`
		}{key, token, solution}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(result); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	fmt.Println("Listening on :9090...")
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func displayCustomText(customTextDone chan<- bool) {
	customText := "Made with   ˗ˋˏ ♡ ˎˊ˗  by GrozdniyAndy of XSS.is"

	for i := 0; i <= len(customText); i++ {
		fmt.Print("\r" + customText[:i] + "_")
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Print("\r" + customText)
	time.Sleep(1000 * time.Millisecond)

	customTextDone <- true

	fmt.Println("")
}

func extractKeyAndToken(anchorURL string) (string, string, error) {
	parsedURL, err := url.Parse(anchorURL)
	if err != nil {
		return "", "", err
	}

	queryValues := parsedURL.Query()
	key := queryValues.Get("k")

	resp, err := http.Get(anchorURL)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("Failed to retrieve URL: %v", resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", "", err
	}

	// Find the token value in the HTML content
	token := ""
	doc.Find("#recaptcha-token").Each(func(index int, element *goquery.Selection) {
		token, _ = element.Attr("value")
	})

	return key, token, nil
}

func getSolution(key, token string) (string, error) {
	url := fmt.Sprintf("https://www.google.com/recaptcha/api2/reload?k=%s", key)
	payload := fmt.Sprintf("reason=q&c=%s&size=invisible", token)

	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(payload))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Failed to retrieve solution: %v", resp.Status)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	re := regexp.MustCompile(`"rresp","(.*?)",`)
	match := re.FindStringSubmatch(string(body))
	if len(match) < 2 {
		return "", fmt.Errorf("Failed to extract 'rresp' value from the response body")
	}

	solution := match[1]

	return solution, nil
}
