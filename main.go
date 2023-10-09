package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	controller "com.gedalias/infra/controller"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Seja bem vindo")
}

func configRouter(router *mux.Router) {
	router.HandleFunc("/", homePage)
	router.HandleFunc("/livros", controller.ListarTodosLivros).Methods("GET")
	router.HandleFunc("/livros/{id}", controller.ConsultarLivro).Methods("GET")
	router.HandleFunc("/livros", controller.CadastrarLivros).Methods("POST")
	router.HandleFunc("/livros/{id}", controller.ModificarLivro).Methods("PUT")
	router.HandleFunc("/livros/{id}", controller.ExcluirLivro).Methods("DELETE")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	fmt.Println("carregando as rotas...")
	configRouter(router)

	fmt.Println("iniciando o servidor...")
	log.Fatal(http.ListenAndServe(":1337", router))
}
