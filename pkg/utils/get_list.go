package utils

import "bytes"

// GetList gets important data from body and places into special frame
func GetList(body []byte) []byte {

	var buf bytes.Buffer

	buf.WriteRune(openCurlyBracket)
	buf.WriteString(headerLists)
	buf.WriteRune(openBracket)

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
			buf.WriteRune(delimiter)
			start, finish = -1, -1
		}
	}

	buf.Truncate(buf.Len() - 1)

	buf.WriteRune(closeBracket)
	buf.WriteRune(closeCurlyBracket)

	return buf.Bytes()
}
