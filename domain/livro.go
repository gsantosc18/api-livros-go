package domain

type Livro struct {
	Id     int
	Titulo string
	Autor  string
}

func NewLivro(id int, titulo, autor string) *Livro {
	livro := Livro{
		Id:     id,
		Titulo: titulo,
		Autor:  autor,
	}
	return &livro
}
