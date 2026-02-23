package main

import (
	"fmt"
	"math/rand"
	"time"
)

func jogador(nome string, numeroScreto int, resultado chan string) {
	for tentativa := 1; tentativa <= 5; tentativa++ {
		guess := rand.Intn(100) + 1
		fmt.Printf("%s tentou adivinhar o número: %d\n", nome, guess)
		if guess == numeroScreto {
			resultado <- fmt.Sprintf("%s acertou o número secreto %d em %d tentativas!", nome, numeroScreto, tentativa)
			return
		}
		time.Sleep(500 * time.Millisecond) // Simula o tempo de pensamento
	}
	resultado <- fmt.Sprintf("%s não conseguiu adivinhar o número secreto %d em 5 tentativas.", nome, numeroScreto)
}

func main() {
	rand.Seed(time.Now().UnixNano()) //garante que os números aleatórios sejam diferentes a cada execução
	numeroScreto := rand.Intn(100) + 1

	fmt.Println("Número secreto sorteado! Que vença o melhor jogador.")
	resultado := make(chan string)

	nomesJogadores := []string{"Alice", "Bob", "Charlie", "Diana"}
	for _, nome := range nomesJogadores {
		go jogador(nome, numeroScreto, resultado)
	}

	for i := 0; i < len(nomesJogadores); i++ {
		fmt.Println(<-resultado)
	}

	fmt.Println("Fim do jogo! Número secreto era:", numeroScreto)
}
