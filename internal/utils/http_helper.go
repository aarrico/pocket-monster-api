package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetBodyFromUrl(url string, exitOnFail bool) []byte {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error accessing url:", err)
		if exitOnFail {
			os.Exit(1)
		}
		return []byte{}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading response body:", err)
		if exitOnFail {
			os.Exit(1)
		}
		return []byte{}
	}

	return body
}
