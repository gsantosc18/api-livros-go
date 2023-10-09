package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"com.gedalias/domain"
	livrorepository "com.gedalias/infra/livrorepository"
	"github.com/gorilla/mux"
)

func ListarTodosLivros(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(livrorepository.ListaLivros())
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

	livroCriado := livrorepository.CreateNewLivro(novoLivro)

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

	id, error := strconv.Atoi(params["id"])

	if error != nil {
		fmt.Printf("Erro ao atualizar o livro. Erro: %s\n", error.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	corpo, error := io.ReadAll(r.Body)

	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var livroModificado domain.Livro

	errorJson := json.Unmarshal(corpo, &livroModificado)

	if errorJson != nil {
		fmt.Printf("Erro ao deserializar o json em objeto na atualização do livro. Erro: %s\n", errorJson.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	livroAlterado, error := livrorepository.AtualizaLivro(id, livroModificado)

	if error != nil {
		fmt.Printf("Erro ao atualizar o livro. Erro: %s\n", error.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-type", "application/json")
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

	livroSelecionado, error := livrorepository.BuscarLivro(id)

	if error != nil {
		fmt.Printf("Erro ao consultar o livro. Erro: %s\n", error.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(livroSelecionado)
}
