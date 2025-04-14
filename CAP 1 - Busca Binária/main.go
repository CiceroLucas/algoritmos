package main

import (
	"fmt"
	"sort"
)


func ordenarNomes(nomes []string) []string {
	sort.Strings(nomes)
	return nomes
}

func main() {
	nomes := []string{"Carlos", "Lucas", "Bruno", "Daniel", "Beatriz", "Laryssa", "Marcos", "Arthur"}
	
	nomesOrdenados := ordenarNomes(nomes)
	fmt.Println("Lista ordenada:", nomesOrdenados)

	nomePesquisado := "Lucas"

	inicio := 0
	fim := len(nomesOrdenados) - 1
	operacoes := 0

	for inicio <= fim {
		meio := (inicio + fim) / 2
		nomeAtual := nomesOrdenados[meio]
		
		if nomeAtual == nomePesquisado {
			println("Nome encontrado", nomeAtual) 
			break
		} 
		if nomeAtual > nomePesquisado  {
			fim = meio - 1
		}
		if nomeAtual < nomePesquisado{
			inicio = meio + 1
		}
		operacoes++
	}
	println("Quantidade de operações", operacoes)
}
