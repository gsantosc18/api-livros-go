package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"com.gedalias/domain"
	livrorepository "com.gedalias/infra/livrorepository"
)

func listarLivros(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(livrorepository.ListaLivros())
}

func cadastrarLivros(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

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

func excluirLivro(w http.ResponseWriter, r *http.Request) {
	partes := strings.Split(r.URL.Path, "/")

	id, _ := strconv.Atoi(partes[2])

	error := livrorepository.ExcluirLivro(id)

	if error != nil {
		fmt.Printf("Erro ao excluir o livro. Erro: %s\n", error.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func modificarLivro(w http.ResponseWriter, r *http.Request) {
	partes := strings.Split(r.URL.Path, "/")

	id, error := strconv.Atoi(partes[2])

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

func consultarLivro(w http.ResponseWriter, r *http.Request) {
	partes := strings.Split(r.URL.Path, "/")

	if len(partes) > 3 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	id, _ := strconv.Atoi(partes[2])

	livroSelecionado, error := livrorepository.BuscarLivro(id)

	if error != nil {
		fmt.Printf("Erro ao consultar o livro. Erro: %s\n", error.Error())
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(livroSelecionado)
}

func LivroController(w http.ResponseWriter, r *http.Request) {
	partes := strings.Split(r.URL.Path, "/")
	switch r.Method {
	case "GET":
		if len(partes) >= 3 {
			consultarLivro(w, r)
		} else {
			listarLivros(w, r)
		}
	case "POST":
		cadastrarLivros(w, r)
	case "DELETE":
		excluirLivro(w, r)
	case "PUT":
		modificarLivro(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Método não reconhecido")
	}
}
