package application

import "com.gedalias/domain"

type LivroRepository interface {
	CreateNewLivro(livro domain.Livro)
}
