package main

import "fmt"

type Musicas struct {
	Nome  string
	Plays int
}

func main() {
	musicas := []Musicas{
		{Nome: "Controllah - Gorillaz", Plays: 15},
		{Nome: "I KNOW? - Travis Scott", Plays: 26},
		{Nome: "Guardians of Earth - Sepultura", Plays: 54},
	}

	var ordenadas []Musicas

	for range musicas {
		maior := buscaMaior(musicas)
		ordenadas = append(ordenadas, musicas[maior])
		musicas = append(musicas[:maior], musicas[maior+1:]...)
	}

	fmt.Println(ordenadas)
}

func buscaMaior(musicas []Musicas) int {
	maior := 0
	for k, musica := range musicas {
		if musica.Plays > musicas[maior].Plays {
			maior = k
		}
	}

	return maior
}
