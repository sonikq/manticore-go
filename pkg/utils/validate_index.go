package utils

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
)

var layoutValidateIndex = `{
    "index": "%s",
    "limit":0
}`

func ValidateIndex(urlManticore, index string) error {
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf(layoutValidateIndex, index))

	resp, err := http.Post(urlManticore, "application/json", &buf)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body := make([]byte, 1)
	_, err = io.ReadAtLeast(resp.Body, body, 1)
	if err != nil {
		return err
	}

	if body[0] == '[' {
		return errors.New("there is no such index")
	}

	return nil
}
