package dto

import "com.gedalias/domain"

type UpdateLivroDTO struct {
	Titulo string
	Autor  string
}

func (u *UpdateLivroDTO) Domain() domain.Livro {
	return domain.Livro{
		Titulo: u.Titulo,
		Autor:  u.Autor,
	}
}
