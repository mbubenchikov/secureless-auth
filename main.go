package main

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"log"
	"net/http"

	"github.com/luikyv/go-oidc/pkg/goidc"
	"github.com/luikyv/go-oidc/pkg/provider"
)

func main() {
	key, _ := rsa.GenerateKey(rand.Reader, 2048)
	jwks := goidc.JSONWebKeySet{
		Keys: []goidc.JSONWebKey{{
			KeyID:     "key_id",
			Key:       key,
			Algorithm: "RS256",
		}},
	}

	op, _ := provider.New(
		goidc.ProfileOpenID,
		"http://localhost",
		func(_ context.Context) (goidc.JSONWebKeySet, error) {
			return jwks, nil
		},
	)

	mux := http.NewServeMux()
	mux.Handle("/", op.Handler())

	server := &http.Server{
		Addr:    ":80",
		Handler: mux,
	}
	log.Fatal(server.ListenAndServe())
}
