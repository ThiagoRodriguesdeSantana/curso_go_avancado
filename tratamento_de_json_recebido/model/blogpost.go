package model

//BlogPost armazena dados do json
type BlogPost struct {
	UsuarioID int    `json:"id"`
	Titulo    string `json:"title"`
}
