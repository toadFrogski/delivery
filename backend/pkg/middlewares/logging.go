package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s: Before\n", time.Now().Local())
		next(w, r)
		fmt.Printf("%s: After\n", time.Now().Local())
	}
}
