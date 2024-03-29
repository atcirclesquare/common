package htmlrender

import (
	"context"
	"errors"
	"net/http"

	"github.com/atcirclesquare/common/http/middleware"
	"github.com/atcirclesquare/common/template/html"
)

type ctxKey struct{}

var kCtxKey = ctxKey{}

func Middleware(r html.Renderer) middleware.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			next.ServeHTTP(w, req.WithContext(context.WithValue(req.Context(), kCtxKey, &render{r})))
		})
	}
}

// Render renders an HTML template with the given name and variables.
func Render(w http.ResponseWriter, req *http.Request, name string, vars ...any) error {
	render, ok := req.Context().Value(kCtxKey).(*render)
	if !ok {
		return errors.New("http/render: unable to retrieve render from request context. Make sure to use corresponding middleware.")
	}

	if err := render.Render(w, name, vars); err != nil {
		return err
	}

	return nil
}
