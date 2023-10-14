package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	controller "com.gedalias/infra/controller"
	"com.gedalias/infra/database"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Seja bem vindo")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	fmt.Println("iniciando o banco de dados...")

	database.GetInstance()

	fmt.Println("carregando as rotas...")

	router.HandleFunc("/", homePage)

	controller.LivroRouter(router)

	fmt.Println("iniciando o servidor...")
	log.Fatal(http.ListenAndServe(":1337", router))
}
