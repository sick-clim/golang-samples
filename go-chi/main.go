package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func handlePostDeail(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "postID")

	response := fmt.Sprintf("Detail of post id %s", postID)
	fmt.Fprint(w, response)

}

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	r.Get("/posts/{postID}/detail", handlePostDeail)

	http.ListenAndServe(":3000", r)
}
