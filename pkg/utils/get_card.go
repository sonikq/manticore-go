package utils

import "bytes"

// GetCard gets important data from body and places into special frame
func GetCard(body []byte) []byte {

	var buf bytes.Buffer

	buf.WriteRune(openCurlyBracket)
	buf.WriteString(headerCards)

	start, finish, count := -1, -1, 0
	for i, s := range body {
		if s == '{' {
			count++
			if count == 3 {
				start, finish = i, -1
			}
		} else if s == '}' {
			if count == 3 {
				finish = i
			}
			count--
		}

		if start > 0 && finish > 0 {
			buf.Write(body[start : finish+1])
			start, finish = -1, -1
		}
	}

	buf.WriteRune(closeCurlyBracket)

	return buf.Bytes()
}
