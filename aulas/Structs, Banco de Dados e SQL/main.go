package main

import "fmt"

// soma recebe dois números e devolve a soma deles.
func soma(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("Olá, Estruturas de Dados!")

	// Variáveis
	nome := "Maria"
	idade := 20
	fmt.Println(nome, "tem", idade, "anos")

	// Função
	fmt.Println("2 + 3 =", soma(2, 3))

	// Slice e for
	notas := []int{7, 8, 10}
	for i := 0; i < len(notas); i++ {
		fmt.Println("Nota", i, "=", notas[i])
	}
}
