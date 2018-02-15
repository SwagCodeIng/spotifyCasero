package main

import(
	"net/http"
	"io"
	"fmt"
)

func main() {

	http.HandleFunc("/otraRuta", otraRuta)

	http.HandleFunc("/",handler)

	http.ListenAndServe(":9000", nil)
}

func otraRuta(w http.ResponseWriter,r *http.Request){
	fmt.Println("Otra ruta")
	io.WriteString(w, "Este path es para otra ruta")
}

func handler(w http.ResponseWriter,r *http.Request){
	fmt.Println("Hay una nueva petición")
	io.WriteString(w, "hola mundo")
}