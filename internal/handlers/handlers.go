package handlers

import (
	"context"
	"net/http"
	"path/filepath"

	"github.com/benleem/benmarshall/internal/handlers/routes"
)

func Init(key string) *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./static")})
	home := routes.NewHomeHandler()

	mux.Handle("GET /static", http.NotFoundHandler())
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.Handle("GET /{$}", htmxMiddleware(http.HandlerFunc(home.Get)))

	return mux
}

type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if err != nil {
		return nil, err
	}

	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}

	return f, nil
}

func htmxMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hxReq := r.Header.Get("Hx-Request")
		if hxReq != "" {
			r.Header.Set("Vary", "Hx-Request")
			// return page.Render(context.Background(), c.Response().Writer)
			next.ServeHTTP(w, r.WithContext(context.WithValue(context.Background(), "htmx", true)))
			return
		}
		// return templates.Layout(page, "benmarshall").Render(context.Background(), c.Response().Writer)
		next.ServeHTTP(w, r.WithContext(context.WithValue(context.Background(), "htmx", false)))
	})
}
