package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	domain "com.gedalias/domain"
	dto "com.gedalias/infra/controller/dto"
	livrorepository "com.gedalias/infra/livrorepository"
	"github.com/gorilla/mux"
)

func ListarTodosLivros(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")
	livros, err := livrorepository.ListaLivros()

	if err != nil {
		log.Println("Houve um erro interno na listagem dos livros. Error: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(livros)
}

func CadastrarLivros(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	body, error := io.ReadAll(r.Body)

	if error != nil {
		fmt.Printf("Erro ao deserializar o corpo da request de crição do livro. Erro: %s\n", error.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var novoLivro domain.Livro
	json.Unmarshal(body, &novoLivro)

	livroCriado, err := livrorepository.CreateNewLivro(novoLivro)

	if err != nil {
		log.Println("Houve um erro interno na criação do novo livro. Error: " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(livroCriado)
}

func ExcluirLivro(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	error := livrorepository.ExcluirLivro(id)

	if error != nil {
		fmt.Printf("Erro ao excluir o livro. Erro: %s\n", error.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func ModificarLivro(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Printf("Erro ao atualizar o livro. Erro: %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !livrorepository.ExistLivro(id) {
		log.Println("O liro não foi encontrado para atualização.")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	corpo, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var livroModificado dto.UpdateLivroDTO

	errorJson := json.Unmarshal(corpo, &livroModificado)

	if errorJson != nil {
		fmt.Printf("Erro ao deserializar o json em objeto na atualização do livro. Erro: %s\n", errorJson.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	livroAlterado, err := livrorepository.AtualizaLivro(id, livroModificado.Domain())

	if err != nil {
		fmt.Printf("Erro ao atualizar o livro. Erro: %s\n", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(livroAlterado)
}

func ConsultarLivro(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if params["id"] == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	id, _ := strconv.Atoi(params["id"])

	if !livrorepository.ExistLivro(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	livroSelecionado := livrorepository.BuscarLivro(id)

	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(livroSelecionado)
}

func LivroRouter(router *mux.Router) {
	router.HandleFunc("/livros", ListarTodosLivros).Methods("GET")
	router.HandleFunc("/livros/{id}", ConsultarLivro).Methods("GET")
	router.HandleFunc("/livros", CadastrarLivros).Methods("POST")
	router.HandleFunc("/livros/{id}", ModificarLivro).Methods("PUT")
	router.HandleFunc("/livros/{id}", ExcluirLivro).Methods("DELETE")
}
