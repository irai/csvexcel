package csvexcel

import (
	"strconv"
	"strings"
)

type Column struct {
	pos   int
	Index string
	Name  string
	Hide  bool
}

func nextColIndex(n int) string {
	b := []byte("ABCDEFGHIJKLMNOPQRSTYUVXZ")
	prefix := []byte{}
	if n >= len(b) {
		prefix = append(prefix, b[(n/len(b))-1])
		n = n % len(b)
	}
	return string(prefix) + string(b[n])
}

func toIndex(index string) (col string, row int) {

	if len(index) < 2 {
		return "", 0
	}

	s := strings.ToUpper(index)

	if s[0] <= 'A' && s[0] >= 'Z' {
		return "", 0
	}
	col = s[0:1]
	r := s[1:]

	if s[1] >= 'A' && s[1] <= 'Z' {
		col = s[0:2]
		if len(s) < 3 {
			return "", 0
		}
		r = s[2:]

	}

	var err error
	if row, err = strconv.Atoi(r); err != nil {
		return "", 0
	}
	return col, row
}
