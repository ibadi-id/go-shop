package main

import (
	"fmt"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/ibadi-id/gostart/pkg/config"
)

func TestRouter(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)
	switch v := mux.(type){
	case *chi.Mux:
		// pass
	default:
		t.Error(fmt.Sprintf("type is not http.Handle, but %T", v))
	}
}