package main

import (
	"fmt"
	"net/http"
	"testing"
)

type MiddlewareHandler struct{}

func (mh *MiddlewareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}

func TestNoSurf(t *testing.T) {
	var mh MiddlewareHandler
	h := NoSurf(&mh)

	switch v := h.(type) {
	case http.Handler:
		//
	default:
		t.Error(fmt.Printf("type is %T and not http.Handler", v))
	}
}

func TestSessionLoad(t *testing.T) {
	var mh MiddlewareHandler
	h := NoSurf(&mh)

	switch v := h.(type) {
	case http.Handler:
		//
	default:
		t.Error(fmt.Printf("type is %T and not http.Handler", v))
	}
}
