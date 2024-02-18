package utils

import "net/http"

type middleware func(http.HandlerFunc) http.HandlerFunc

func CreateChain(f http.HandlerFunc, m ...middleware) http.HandlerFunc {
	if len(m) == 0 {
		return f
	}

	return m[0](CreateChain(f, m[1:]...))
}
