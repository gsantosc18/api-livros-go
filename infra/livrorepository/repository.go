package infra

import (
	"errors"
	"log"

	"com.gedalias/domain"
	"com.gedalias/infra/database"
)

func CreateNewLivro(livro domain.Livro) (*domain.Livro, error) {
	stmt := "INSERT INTO livro (titulo, autor) VALUES (?, ?)"

	result, err := database.GetInstance().Exec(stmt, livro.Titulo, livro.Autor)

	if err != nil {
		return &domain.Livro{}, err
	}

	id, _ := result.LastInsertId()
	livro.Id = int(id)
	return &livro, nil
}

func ListaLivros() ([]domain.Livro, error) {
	stmt := "SELECT id, titulo, autor FROM livro"
	rows, err := database.GetInstance().Query(stmt)

	if err != nil {
		return []domain.Livro{}, err
	}

	defer rows.Close()

	var livros []domain.Livro = make([]domain.Livro, 0)

	for rows.Next() {
		var livro domain.Livro

		err := rows.Scan(&livro.Id, &livro.Titulo, &livro.Autor)

		if err != nil {
			return []domain.Livro{}, err
		}

		livros = append(livros, livro)
	}

	return livros, nil
}

func ExcluirLivro(id int) error {

	if !ExistLivro(id) {
		return errors.New("não foi encontrado o livro")
	}

	stmt := "DELETE FROM livro WHERE id = ?"

	_, err := database.GetInstance().Exec(stmt, id)

	return err
}

func AtualizaLivro(livroId int, livro domain.Livro) (domain.Livro, error) {

	if !ExistLivro(livroId) {
		return domain.Livro{}, errors.New("não foi encontrado o livro")
	}

	stmt := "UPDATE livro SET titulo = ?, autor = ? WHERE id = ?"

	_, err := database.GetInstance().Exec(stmt, livro.Titulo, livro.Autor, livroId)

	if err != nil {
		return domain.Livro{}, err
	}

	livro.Id = int(livroId)

	return livro, nil
}

func BuscarLivro(livroId int) domain.Livro {
	stmt := "SELECT id, titulo, autor FROM livro WHERE id = ?"
	row := database.GetInstance().QueryRow(stmt, livroId)

	var livro domain.Livro

	err := row.Scan(&livro.Id, &livro.Titulo, &livro.Autor)

	if err != nil {
		log.Fatal(err.Error())
	}

	return livro
}

func ExistLivro(livroId int) bool {
	stmt := "SELECT 1 FROM livro WHERE id = ?"
	row := database.GetInstance().QueryRow(stmt, livroId)

	var exist int
	err := row.Scan(&exist)

	if err != nil {
		log.Print(err.Error())
		return false
	}

	return exist > 0
}
