package controllers

import "net/http"

//CriarUsuário adiciona um novo usuário ao banco de dados
func CriarUsuario(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Criando um usuário"))
}

//BuscarUsuarios busca todos os usuários no banco de dados
func BuscarUsuarios(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Buscando todos os usuários"))
}

//BuscaUsuario serve apenas para buscar um usuário
func BuscarUsuario(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Buscando um usuário"))
}

//AtualizarUsuario modifica alteracoes de um usuario no banco
func AtualizarUsuario(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Atualizando um usuário"))
}

//DeletarUsuario exclui as informacoes de um usuario.
func DeletarUsuario(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Criando um usuário"))
}
