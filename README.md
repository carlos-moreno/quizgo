# 🧠 Go Tech Quiz CLI

O **Go Tech Quiz CLI** é uma aplicação de terminal robusta e interativa desenvolvida em **Go**. Ele oferece um motor de quiz dinâmico que consome bases de dados em CSV, focado em avaliar conhecimentos técnicos em Engenharia de Software, Linguagens de Programação, DevOps e Arquitetura.

Este projeto foi construído seguindo as melhores práticas da linguagem Go, visando performance, legibilidade e uma experiência de usuário (UX) fluida via linha de comando.

## ✨ Funcionalidades

*   **Engine de Perguntas Dinâmicas:** Carrega perguntas de forma escalonável via arquivo CSV.
*   **Embaralhamento Inteligente:** Utiliza algoritmos de *shuffle* para garantir que tanto a ordem das perguntas quanto a posição das alternativas sejam únicas a cada rodada.
*   **UX de Terminal Avançada:** Implementa sequências de escape ANSI para validação de entrada, permitindo apagar e sobrescrever linhas de erro sem "poluir" o histórico do terminal.
*   **Métricas de Performance:** Ao final de cada sessão, o jogador recebe um relatório com o total de acertos e o percentual de aproveitamento.
*   **Fluxo Contínuo:** Suporte a reinicialização rápida de partidas com preservação de estado e validação de comandos estilo Linux (`Y/n`).
*   **Documentação Nativa:** Código 100% documentado seguindo os padrões do `Go Doc`.

## 🛠️ Tecnologias e Conceitos

*   **Linguagem:** Go (v1.20+)
*   **I/O:** `bufio` para leitura otimizada e `encoding/csv` para parsing de dados.
*   **UX:** Sequências de escape ANSI para manipulação de cursor no terminal.
*   **Random:** Implementação moderna utilizando o gerador global do Go (v1.20+).

## 🚀 Como Começar

### Pré-requisitos
*   [Go](https://go.dev/doc/install) instalado no sistema.

### Instalação e Execução
1. Clone este repositório:
   ```bash
   git clone https://github.com/seu-usuario/go-tech-quiz.git
   cd go-tech-quiz
   ```

2. Certifique-se de que a estrutura de diretórios contém o arquivo de perguntas:
   ```text
   .
   ├── main.go
   └── assets/
       └── quiz-go.csv
   ```

3. Execute o programa:
   ```bash
   go run main.go
   ```

## 📂 Estrutura do Arquivo CSV
O jogo espera um arquivo em `assets/quiz-go.csv` com a seguinte estrutura:
`Pergunta, Opção1, Opção2, Opção3, Opção4, RespostaCorreta`

## 📖 Documentação Técnica
Para explorar a documentação técnica das structs e métodos através do terminal, utilize:
```bash
go doc -all
```

## ⚖️ Licença

Este programa é um software livre; você pode redistribuí-lo e/ou modificá-lo sob os termos da **Licença Pública Geral GNU (GPL)** conforme publicada pela *Free Software Foundation*; tanto a **versão 3** da Licença, como (a seu critério) qualquer versão posterior.

Consulte o arquivo [LICENSE](LICENSE) para obter mais detalhes.

---
*Desenvolvido como projeto de estudo em Go para aprimoramento de lógica e interfaces CLI.*
