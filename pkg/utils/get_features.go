package utils

import (
	"bytes"
)

func GetFeatures(body []byte) []byte {

	var buf bytes.Buffer

	buf.WriteRune(openCurlyBracket)
	buf.WriteString(headerFeatures)
	buf.WriteRune(openBracket)

	start, finish, count := -1, -1, 0
	for i, s := range body {
		if s == '{' {
			count++
			if count == 5 {
				start, finish = i, -1
			}
		} else if s == '}' {
			if count == 5 {
				finish = i
			}
			count--
		}

		if start > 0 && finish > 0 {
			buf.Write(body[start : finish+1])
			buf.WriteRune(delimiter)
			start, finish = -1, -1
		}
	}

	buf.Truncate(buf.Len() - 1)

	buf.WriteRune(closeBracket)
	buf.WriteRune(closeCurlyBracket)

	return buf.Bytes()
}
