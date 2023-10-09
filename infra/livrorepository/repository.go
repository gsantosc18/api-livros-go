package infra

import (
	"errors"

	"com.gedalias/domain"
)

var livros []domain.Livro = []domain.Livro{}

var nextId int = len(livros)

func buscaLivroPeloId(id int, callback func(index int)) {
	for index, livro := range livros {
		if livro.Id == id {
			callback(index)
			break
		}
	}
}

func CreateNewLivro(livro domain.Livro) domain.Livro {
	livro.Id = GetNextId()
	livros = append(livros, livro)
	return livro
}

func ListaLivros() *[]domain.Livro {
	return &livros
}

func ExcluirLivro(id int) error {

	indexLivro := -1

	buscaLivroPeloId(id, func(index int) { indexLivro = index })

	if indexLivro == -1 {
		return errors.New("não foi encontrado o livro")
	}

	livros = append(livros[:indexLivro], livros[indexLivro+1:]...)

	return nil
}

func AtualizaLivro(id int, livro domain.Livro) (domain.Livro, error) {
	indexLivro := -1

	buscaLivroPeloId(id, func(index int) { indexLivro = index })

	if indexLivro == -1 {
		return domain.Livro{}, errors.New("não foi encontrado o livro")
	}

	livroAntigo := livros[indexLivro]
	livro.Id = livroAntigo.Id
	livros[indexLivro] = livro

	return livro, nil
}

func BuscarLivro(id int) (domain.Livro, error) {
	indexLivro := -1
	buscaLivroPeloId(id, func(index int) { indexLivro = index })

	if indexLivro > -1 {
		return livros[indexLivro], nil
	}

	return domain.Livro{}, errors.New("não foi encontrada o livro")
}

func GetNextId() int {
	nextId++
	return nextId
}
