package utils

import (
	"fmt"
	"io"
	"strings"
)

// readRequestBody godoc
// Fungsi helper untuk logger
func ReadRequestBody(reqBody io.ReadCloser) string {
	body, err := io.ReadAll(reqBody)

	if err != nil {
		return fmt.Sprintf("%s", err.Error())
	}

	strip := strings.ReplaceAll(string(body), "\n", "")

	return fmt.Sprintf("%s", strip)
}

// ParseHttpError godoc
// Fungsi helper untuk logger
func ParseHttpError(error string) string {

	if len(error) < 1 {
		return ""
	}

	return fmt.Sprintf(" - [REQ ERROR] %s", error)
}
