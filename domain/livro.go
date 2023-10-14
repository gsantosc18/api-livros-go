package domain

type Livro struct {
	Id     int    `json:"id"`
	Titulo string `json:"titulo"`
	Autor  string `json:"autor"`
}

func NewLivro(id int, titulo, autor string) *Livro {
	livro := Livro{
		Id:     id,
		Titulo: titulo,
		Autor:  autor,
	}
	return &livro
}
