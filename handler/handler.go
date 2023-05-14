package handler

import (
	_ "fmt"
	"golang-web/entity"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// w.Write([]byte("Welcome"))
	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Print(err)
		http.Error(w, "Error is happening keep calm", http.StatusInternalServerError)
		return
	}

	// data := map[string]interface{}{
	// 	"title": "Learning Golang Web",
	// 	"content": "I'm learning golang web.",
	// }

	data := []entity.Product{
		{ID: 1,
			Name:  "Laptop",
			Price: 6000000,
			Stock: 1},
		{ID: 2,
			Name:  "ZX25R",
			Price: 120000000,
			Stock: 5},
		{ID: 3,
			Name:  "Iphone 14",
			Price: 7000000,
			Stock: 11},
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Print(err)
		http.Error(w, "Error is happening keep calm", http.StatusInternalServerError)
		return
	}

}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world, this is golang"))
}

func NameHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("My name is Gera Anggara Putra"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNumber, err := strconv.Atoi(id)

	if err != nil || idNumber < 1 {
		http.NotFound(w, r)
		return
	}
	// w.Write([]byte("Product Page"))
	// fmt.Fprintf(w, "Product page : %d", idNumber)
	data := map[string]interface{}{
		"id": idNumber,
	}

	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Print(err)
		http.Error(w, "Error is happening keep calm", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Print(err)
		http.Error(w, "Error is happening keep calm", http.StatusInternalServerError)
		return
	}
}

func PostGet(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	switch method {
	case "GET":
		w.Write([]byte("Ini adalah GET"))
	case "POST":
		w.Write([]byte("Ini adalah POST"))
	default:
		http.Error(w, "Error is happening, keep calm", http.StatusBadRequest)
	}
}

func Form(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles(path.Join("views", "form.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Print(err)
			http.Error(w, "Error is happening keep calm", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Print(err)
			http.Error(w, "Error is happening keep calm", http.StatusInternalServerError)
			return
		}
		
		return
	}
	
	http.Error(w, "Error is happening keep calm", http.StatusInternalServerError)
}

func Process(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()	
		if err != nil {
			log.Print(err)
			http.Error(w, "Error is happening keep calm", http.StatusInternalServerError)
			return
		}

		name := r.Form.Get("name")
		message := r.Form.Get("message")

		data := map[string]interface{}{
			"name": name,
			"message": message,
		}
		tmpl, err := template.ParseFiles(path.Join("views", "result.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Print(err)
			http.Error(w, "Error is happening keep calm", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			log.Print(err)
			http.Error(w, "Error is happening keep calm", http.StatusInternalServerError)
			return
		}
		return
	}

	http.Error(w, "Error is happening keep calm", http.StatusInternalServerError)
}
