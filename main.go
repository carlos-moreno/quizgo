package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Question struct {
	Text    string
	Options []string
	Answer  string
}

type GameState struct {
	Name      string
	Points    int
	Questions []Question
}

func (g *GameState) Init() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Seja bem vindo(a) ao quiz")
	fmt.Print("Escreva o seu nome: ")
	name, _ := reader.ReadString('\n')
	g.Name = strings.TrimSpace(name)

	fmt.Printf("Olá %s, o arquivo contém %d perguntas.\n", g.Name, len(g.Questions))

	for {
		fmt.Printf("Quantas perguntas você deseja responder? (1 a %d): ", len(g.Questions))
		input, _ := reader.ReadString('\n')
		limit, err := strconv.Atoi(strings.TrimSpace(input))

		if err != nil || limit < 1 || limit > len(g.Questions) {
			fmt.Println("\033[31mQuantidade inválida!\033[0m")
			continue
		}

		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(g.Questions), func(i, j int) {
			g.Questions[i], g.Questions[j] = g.Questions[j], g.Questions[i]
		})
		g.Questions = g.Questions[:limit]
		break
	}

	fmt.Printf("Vamos ao jogo!\n\n")
}

func (g *GameState) ProccessCSV() {
	f, err := os.Open("assets/quiz-go.csv")
	if err != nil {
		panic("Erro ao abrir o arquivo. Verifique se a pasta 'assets' existe.")
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		panic("Erro ao ler o arquivo")
	}

	for idx, record := range records {
		if idx > 0 {
			options := record[1:5]

			rand.Seed(time.Now().UnixNano())
			rand.Shuffle(len(options), func(i, j int) {
				options[i], options[j] = options[j], options[i]
			})

			question := Question{
				Text:    record[0],
				Options: options,
				Answer:  record[5],
			}
			g.Questions = append(g.Questions, question)
		}
	}
}

func (g *GameState) Run() {
	reader := bufio.NewReader(os.Stdin)

	for idx, question := range g.Questions {
		fmt.Printf("\033[33m %d. %s\033[0m\n", idx+1, question.Text)

		for i, option := range question.Options {
			fmt.Printf("[%d] %s\n", i+1, option)
		}

		var answer int
		erroAtivo := false

		for {
			fmt.Print("Digite uma alternativa: ")
			read, _ := reader.ReadString('\n')
			read = strings.TrimSpace(read)

			var err error
			answer, err = strconv.Atoi(read)

			if err != nil || answer < 1 || answer > len(question.Options) {
				fmt.Print("\033[1A\033[K")
				if erroAtivo {
					fmt.Print("\033[1A\033[K")
				}
				fmt.Printf("\033[31mOpção inválida! Digite de 1 a %d\033[0m\n", len(question.Options))
				erroAtivo = true
				continue
			}

			if erroAtivo {
				fmt.Print("\033[1A\033[K")
				fmt.Print("\033[1A\033[K")
				fmt.Printf("Digite uma alternativa: %d\n", answer)
			}
			break
		}

		if question.Options[answer-1] == question.Answer {
			fmt.Println("\033[32mParabéns você acertou!!!\033[0m")
			g.Points += 10
		} else {
			fmt.Printf("\033[31mResposta incorreta!!!\033[0m\n")
		}
		fmt.Println(strings.Repeat("-", 60))
	}
}

func main() {
	game := &GameState{Points: 0}

	game.ProccessCSV()
	game.Init()
	game.Run()

	fmt.Printf("Fim de jogo %s, você fez %d pontos.\n", game.Name, game.Points)
}
