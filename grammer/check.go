package grammer

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const langToolURL = "https://api.languagetool.org/v2/check"

func CheckGrammer(text string) (*LanguageToolErrors, error) {
	data := url.Values{}
	data.Set("text", text)
	data.Set("language", "en-US")

	client := &http.Client{}

	req, err := http.NewRequest("POST", langToolURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("Error creating request: %w ", err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error sending request: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response: %w", err)
	}

	var errors LanguageToolErrors
	err = json.Unmarshal(body, &errors)
	if err != nil {
		fmt.Println(err)
	}

	return &errors, nil
}
