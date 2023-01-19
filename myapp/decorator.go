package myapp

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type DecoratorFunc func(http.ResponseWriter, *http.Request, http.Handler)

type DecoHandler struct {
	fn DecoratorFunc
	h  http.Handler
}

func indexxxHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World") // Writer에 print 하라
}

func (self *DecoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	self.fn(w, r, self.h)
}

func logger(w http.ResponseWriter, r *http.Request, h http.Handler) {
	start := time.Now()
	log.Print("[LOGGER1] Started")
	h.ServeHTTP(w, r)
	log.Println("[LOGGER1] Completed", time.Since(start).Milliseconds())
}

func logger2(w http.ResponseWriter, r *http.Request, h http.Handler) {
	start := time.Now()
	log.Print("[LOGGER2] Started")
	h.ServeHTTP(w, r)
	log.Println("[LOGGER2] Completed", time.Since(start).Milliseconds())
}

func NewDecoHandler(h http.Handler, fn DecoratorFunc) http.Handler {
	return &DecoHandler{
		fn: fn,
		h:  h,
	}
}

func NewDecoServer() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexxxHandler)

	// "/" 에 logger handler
	// decorator 패턴으로 mux - logger - looger2 로 logger2 start - logger1 start - mux - logger1 end - logger2 end
	h := NewDecoHandler(mux, logger)
	h = NewDecoHandler(mux, logger2)

	return h
}
