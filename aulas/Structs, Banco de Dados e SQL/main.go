package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Variável global com a conexão do banco
var db *sql.DB

// Struct que representa um usuário da tabela
type Usuario struct {
	ID     int
	Nome   string
	Email  string
	Idade  int
	Cidade string
}

// Conecta ao banco uma vez
func conectar() {
	host := "localhost"
	porta := "5433"
	usuario := "postgres"
	senha := "postgres"
	banco := "postgres"
	conexao := "postgres://" + usuario + ":" + senha + "@" + host + ":" + porta + "/" + banco + "?sslmode=disable"

	var err error
	db, err = sql.Open("postgres", conexao)
	if err != nil {
		panic(err)
	}

	fmt.Println("Conectado!")
}

// Cria a tabela se ela ainda não existir
func criarTabela() {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS usuarios (id SERIAL PRIMARY KEY, nome TEXT, email TEXT, idade INT, cidade TEXT)")
	if err != nil {
		panic(err)
	}
}

// Insere alguns usuários de exemplo na tabela
func popularBanco() {
	usuarios := []Usuario{
		{Nome: "karla", Email: "karla@email.com", Idade: 21, Cidade: "Manaus"},
		{Nome: "Jordana", Email: "jordana@email.com", Idade: 19, Cidade: "Itacoatiara"},
		{Nome: "Marina", Email: "marina@email.com", Idade: 24, Cidade: "Parintins"},
		{Nome: "Pedro", Email: "pedro@email.com", Idade: 30, Cidade: "Manacapuru"},
		{Nome: "Luana", Email: "luana@email.com", Idade: 22, Cidade: "Tefé"},
	}

	for _, u := range usuarios {
		db.Exec(
			"INSERT INTO usuarios (nome, email, idade, cidade) VALUES ($1, $2, $3, $4)",
			u.Nome, u.Email, u.Idade, u.Cidade,
		)
	}
}

// Troca o nome do usuário que tem o id informado
func atualizarUsuarioNome(id int, novoNome string) {
	_, err := db.Exec("UPDATE usuarios SET nome = $1 WHERE id = $2", novoNome, id)
	if err != nil {
		panic(err)
	}

	fmt.Println("Usuário atualizado!")
}

// Busca todos os usuários do banco e mostra na tela
func consumirDados() {
	linhas, err := db.Query("SELECT id, nome, email, idade, cidade FROM usuarios ORDER BY id")
	if err != nil {
		panic(err)
	}

	var usuarios []Usuario

	for linhas.Next() {
		var u Usuario
		linhas.Scan(&u.ID, &u.Nome, &u.Email, &u.Idade, &u.Cidade)
		usuarios = append(usuarios, u)
	}

	linhas.Close()

	for _, u := range usuarios {
		fmt.Println(u.ID, "-", u.Nome, "-", u.Email, "-", u.Idade, "anos -", u.Cidade)
	}
}

// Ponto de partida do programa: chama as outras funções em ordem
func main() {
	conectar()
	criarTabela()
	popularBanco()
	//atualizarUsuarioNome()
	consumirDados()
}
