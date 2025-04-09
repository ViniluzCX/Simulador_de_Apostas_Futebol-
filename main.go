package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // Garante que o número aleatório mude

	var time1, time2 string
	var palpite1, palpite2 int

	fmt.Println("🏆 Simulador de Apostas de Futebol 🏆")
	fmt.Print("Digite o nome do primeiro time: ")
	fmt.Scanln(&time1) // grava oque eu escrevi no time 1

	fmt.Print("Digite o nome do segundo time: ")
	fmt.Scanln(&time2) // grava oque eu escrevi no time 2

	fmt.Printf("Qual seu palpite para o jogo %s x %s?\n", time1, time2)
	fmt.Print("Gols do ", time1, ": ")
	fmt.Scanln(&palpite1) // grava meu palpite do time 1

	fmt.Print("Gols do ", time2, ": ")
	fmt.Scanln(&palpite2) // grava meu palpite do time 2

	// Sorteia os gols aleatórios de cada time (0 a 8)
	gols1 := rand.Intn(8)
	gols2 := rand.Intn(8)

	fmt.Printf("\nResultado real: %s %d x %d %s\n", time1, gols1, gols2, time2)

	if palpite1 == gols1 && palpite2 == gols2 {
		fmt.Println("✅ Você acertou o placar! Parabéns!")
	} else {
		fmt.Println("❌ Você errou o placar. Tente novamente!")
	}
}
