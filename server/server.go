package server

import (
	"net/http"
	"time"
)

func FileServe(port string, path string) error {
	fs := http.FileServer(http.Dir(path))
	return http.ListenAndServe(port, fs)
}
func Serve(port string, handler http.Handler) error {
	s := &http.Server{
		Addr:           port,
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return s.ListenAndServe()

}
