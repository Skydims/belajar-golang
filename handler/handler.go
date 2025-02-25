package handler

import (
	"log"
	"net/http"
	"path"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf(r.URL.Path)

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("html", "index.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "lagi ada error beroo ", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "lagi ada nnuuu", http.StatusInternalServerError)
		return
	}

}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok masbre"))
}

func DimasHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok wir"))
}
