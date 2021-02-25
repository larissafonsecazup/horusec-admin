package main

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/thedevsaddam/renderer"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var rnd = renderer.New(renderer.Options{ParseGlobPattern: "web/template/*.gohtml"})

const addr = ":3001"

func main() {
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		err := rnd.HTML(w, http.StatusOK, "index", nil)
		checkErr(err)
	})
	srv := &http.Server{Addr: addr, Handler: r}

	go func() {
		log.Println("Listening on ", addr)
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	<-stopChan
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	srv.Shutdown(ctx)
	defer cancel()
	log.Println("Server gracefully stopped!")
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err) //respond with error page or message
	}
}
