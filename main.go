package main

import (
	"gin-sample-app/server"
	"net/http"
	"time"
	"golang.org/x/sync/errgroup"
	"log"
)

var (
	g errgroup.Group
)

func main() {
	server01 := &http.Server{
		Addr: ":8080",
		Handler: server.NewsServer(),
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server02 := &http.Server{
		Addr: ":8081",
		Handler: server.WeatherServer(),
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server03 := &http.Server{
		Addr: ":8082",
		Handler: server.NewBasicAuthServer(),
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return server01.ListenAndServe()
	})

	g.Go(func() error {
		return server02.ListenAndServe()
	})

	g.Go(func() error {
		return server03.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
