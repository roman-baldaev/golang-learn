package strings

import (
	"bytes"
	"fmt"
)

func RecursiveSeparator(s string, sep string, qnt int) string {
	n := len(s)
	if n <= qnt {
		return s
	}
	return RecursiveSeparator(s[:n-qnt], sep, qnt) + sep + s[n-qnt:]
}

func BufferSeparator(s string, sep string, qnt int) string {
	var buf bytes.Buffer
	n := len(s)
	if n <= qnt {
		return s
	}
	buf.WriteString(s[:qnt])
	i := qnt
	for ;i + qnt < n; i += qnt {
		buf.WriteString(sep)
		fmt.Fprintf(&buf, "%s", s[i:i+qnt])
	}
	buf.WriteString(sep)
	buf.WriteString(s[i:])

	return buf.String()
}