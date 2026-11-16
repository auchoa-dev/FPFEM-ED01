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
	ID   int
	Nome string
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
	// Testa se a conexão realmente funciona
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Conectado!")
}

// Cria a tabela se ela ainda não existir
func criarTabela() {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS usuarios (id SERIAL PRIMARY KEY, nome TEXT)")
	if err != nil {
		panic(err)
	}
}

func popularBanco() {
	usuarios := []Usuario{
		{Nome: "Caleta"},
		{Nome: "Jords"},
	}

	for _, u := range usuarios {
		db.Exec("INSERT INTO usuarios (nome) VALUES ($1)", u.Nome)
	}
}

func consumirDados() {
	rows, err := db.Query("SELECT id, nome FROM usuarios ORDER BY id")
	if err != nil {
		panic(err)
	}

	rows.Close()

	var usuarios []Usuario

	for rows.Next() {
		var u Usuario
		rows.Scan(&u.ID, &u.Nome)
		usuarios = append(usuarios, u)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}

	fmt.Print(usuarios)
}

func atualizarUsuario(id int, novoNome string) {
	resultado, err := db.Exec(
		"UPDATE usuarios SET nome = $1 WHERE id = $2",
		novoNome,
		id,
	)

	if err != nil {
		panic(err)
	}

	linhas, err := resultado.RowsAffected()
	if err != nil {
		panic(err)
	}

	if linhas == 0 {
		fmt.Println("Usuário não encontrado!")
		return
	}

	fmt.Println("Usuário atualizado com sucesso!")
}

func main() {
	conectar()
	criarTabela()
	popularBanco()
	//atualizarUsuario()
	consumirDados()
}
