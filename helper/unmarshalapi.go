package helper

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

func HttpGetJson(url string, target interface{}) error {
	// Pastikan URL ada protokolnya
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Jika status bukan 200, anggap error
	if resp.StatusCode != http.StatusOK {
		return errors.New("failed request with status: " + resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, target)
	if err != nil {
		return err
	}

	return nil
}
