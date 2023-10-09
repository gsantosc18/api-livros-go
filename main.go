package main

import (
	"fmt"
	"log"
	"net/http"

	controller "com.gedalias/infra/controller"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Seja bem vindo")
}

func routes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/livros", controller.LivroController)
	http.HandleFunc("/livros/", controller.LivroController)
}

func main() {
	fmt.Println("carregando as rotas...")
	routes()
	fmt.Println("Iniciando o servidor...")
	log.Fatal(http.ListenAndServe(":1337", nil))
}
