# Plano de Workshop: Verificador de URLs Concorrente em Go

## Visão Geral

Este workshop tem como objetivo ensinar os fundamentos de concorrência em Go através da construção passo a passo de um verificador de URLs. Os participantes aprenderão desde a sintaxe básica de Go até conceitos avançados como goroutines e canais, aplicando esses conhecimentos em um projeto prático.

*Duração:* 2 horas
*Público-alvo:* Estudantes de engenharia do primeiro ao terceiro ano
*Abordagem:* Construção progressiva em pacotes, cada um expandindo o anterior

## Estrutura do Workshop

### Pacote 1: ⁠ 01-intro-go-basics ⁠ (25 minutos)

*Objetivo:* Introduzir a sintaxe básica de Go e tratamento de erros
•⁠  ⁠Sintaxe básica: declarações, funções, pacotes
•⁠  ⁠Estruturas de dados relevantes (slices, maps)
•⁠  ⁠Padrão de tratamento de erros em Go
•⁠  ⁠Manipulação de erros com verificação explícita e propagação

*Resultados esperados:*
•⁠  ⁠Compreensão da estrutura básica de um programa Go
•⁠  ⁠Familiaridade com o sistema de tipos de Go
•⁠  ⁠Entendimento do padrão de tratamento de erros em Go

### Pacote 2: ⁠ 02-url-checker-sync ⁠ (20 minutos)

*Objetivo:* Implementar uma solução síncrona básica
•⁠  ⁠Introduzir o problema da verificação de URLs
•⁠  ⁠Implementar uma solução sequencial simples
•⁠  ⁠Aplicar tratamento de erros adequado
•⁠  ⁠Medir o desempenho com várias URLs

*Resultados esperados:*
•⁠  ⁠Implementação funcional de um verificador de URLs sequencial
•⁠  ⁠Compreensão das limitações de uma abordagem síncrona
•⁠  ⁠Capacidade de medir desempenho de operações de I/O

### Pacote 3: ⁠ 03-goroutines-intro ⁠ (25 minutos)
*Objetivo:* Apresentar o conceito de goroutines
•⁠  ⁠Explicar o modelo de concorrência em Go
•⁠  ⁠Implementar o verificador usando goroutines simples
•⁠  ⁠Demonstrar o problema de término prematuro
•⁠  ⁠Discutir casos de uso para goroutines

*Resultados esperados:*
•⁠  ⁠Entendimento do conceito de goroutines
•⁠  ⁠Capacidade de criar goroutines para tarefas concorrentes
•⁠  ⁠Identificação de problemas de sincronização

### Pacote 4: ⁠ 04-waitgroup-sync ⁠ (25 minutos)
*Objetivo:* Resolver o problema de sincronização
•⁠  ⁠Introduzir ⁠ sync.WaitGroup ⁠
•⁠  ⁠Implementar verificador com sincronização adequada
•⁠  ⁠Discutir o desafio da coleta de resultados
•⁠  ⁠Comparar desempenho com a versão síncrona

*Resultados esperados:*
•⁠  ⁠Compreensão do funcionamento do ⁠ sync.WaitGroup ⁠
•⁠  ⁠Implementação de um verificador concorrente com sincronização adequada
•⁠  ⁠Análise de desempenho comparativa

### Pacote 5: ⁠ 05-channels-basic ⁠ (25 minutos)
*Objetivo:* Introduzir canais para comunicação
•⁠  ⁠Explicar canais em Go (conceito, criação, operações)
•⁠  ⁠Implementar verificador com canais para retornar resultados
•⁠  ⁠Demonstrar coleta e exibição de resultados
•⁠  ⁠Implementar timeout para requisições usando context

*Resultados esperados:*
•⁠  ⁠Entendimento do conceito e operações com canais
•⁠  ⁠Implementação de comunicação entre goroutines
•⁠  ⁠Capacidade de implementar timeouts em operações concorrentes

## Detalhamento Técnico

### Pacote 1: Introdução à Sintaxe de Go e Tratamento de Erros
•⁠  ⁠Estrutura básica de um programa Go:
  ⁠ go
  package main
  
  import (
      "fmt"
      "errors"
  )
  
  func main() {
      // Código principal
  }
   ⁠
•⁠  ⁠Tipos de dados fundamentais:
  - Slices: ⁠ urls := []string{"https://exemplo1.com", "https://exemplo2.com"} ⁠
  - Maps: ⁠ resultados := make(map[string]bool) ⁠
•⁠  ⁠Padrão de tratamento de erros:
  - Verificação explícita: ⁠ if err != nil { return err } ⁠
  - Propagação de erros: ⁠ return nil, fmt.Errorf("falha ao conectar: %w", err) ⁠

### Pacote 2: URL Checker Síncrono
•⁠  ⁠Função principal que recebe lista de URLs
•⁠  ⁠Verificação HTTP usando a biblioteca padrão
•⁠  ⁠Loop sequencial para verificar cada URL
•⁠  ⁠Medição de tempo com ⁠ time.Now() ⁠ e ⁠ time.Since() ⁠
•⁠  ⁠Estrutura de dados para armazenar resultados

### Pacote 3: Introdução às Goroutines
•⁠  ⁠Conceito de concorrência vs. paralelismo
•⁠  ⁠Sintaxe de goroutines: ⁠ go funcaoVerificarURL(url) ⁠
•⁠  ⁠Problemas comuns: término prematuro, falta de sincronização
•⁠  ⁠Comparação com abordagens de concorrência em outras linguagens

### Pacote 4: Sincronização com WaitGroup
•⁠  ⁠Funcionamento do ⁠ sync.WaitGroup ⁠:
  ⁠ go
  var wg sync.WaitGroup
  wg.Add(1)
  go func() {
      defer wg.Done()
      // Tarefa concorrente
  }()
  wg.Wait()
   ⁠
•⁠  ⁠Desafios de compartilhamento de dados entre goroutines
•⁠  ⁠Análise de ganhos de desempenho

### Pacote 5: Canais Básicos
•⁠  ⁠Criação e operações com canais:
  ⁠ go
  resultados := make(chan Resultado)
  // Envio para o canal
  resultados <- resultado
  // Recebimento do canal
  res := <-resultados
  // Fechamento do canal
  close(resultados)
   ⁠
•⁠  ⁠Estrutura de resultado completa:
  ⁠ go
  type Resultado struct {
      URL     string
      Status  int
      Erro    error
      Tempo   time.Duration
  }
   ⁠
•⁠  ⁠Implementação de timeout com context:
  ⁠ go
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()
  req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
   ⁠

## Estrutura do Repositório GitHub


github.com/seu-usuario/url-checker-workshop/
├── README.md
├── go.mod
├── 01-intro-go-basics/
│   ├── main.go
│   └── README.md
├── 02-url-checker-sync/
│   ├── main.go
│   └── README.md
├── 03-goroutines-intro/
│   ├── main.go
│   └── README.md
├── 04-waitgroup-sync/
│   ├── main.go
│   └── README.md
└── 05-channels-basic/
    ├── main.go
    └── README.md


## Materiais de Suporte

### Pré-requisitos para Participantes
•⁠  ⁠Go instalado (versão 1.18 ou superior)
•⁠  ⁠Editor de código (VSCode, GoLand, etc.)
•⁠  ⁠Git para clonar o repositório do workshop
•⁠  ⁠Conhecimentos básicos de programação

### Recursos Visuais Recomendados
•⁠  ⁠Diagrama de execução síncrona vs. concorrente
•⁠  ⁠Ilustração do funcionamento de goroutines
•⁠  ⁠Representação visual de canais
•⁠  ⁠Gráficos de comparação de desempenho

### Lista de URLs para Teste
•⁠  ⁠Sites funcionais: google.com, github.com, golang.org
•⁠  ⁠Sites com delay: httpbin.org/delay/2
•⁠  ⁠Sites não-existentes: naoexiste123456789.com

## Cronograma Detalhado

•⁠  ⁠*00:00-00:05* - Boas-vindas e visão geral do workshop
•⁠  ⁠*00:05-00:30* - Pacote 1: Sintaxe de Go e tratamento de erros
•⁠  ⁠*00:30-00:50* - Pacote 2: Implementação do verificador síncrono
•⁠  ⁠*00:50-01:15* - Pacote 3: Introdução às goroutines
•⁠  ⁠*01:15-01:40* - Pacote 4: Sincronização com WaitGroup
•⁠  ⁠*01:40-02:05* - Pacote 5: Comunicação com canais
•⁠  ⁠*02:05-02:15* - Recapitulação, perguntas e próximos passos

## Boas Práticas a Enfatizar

1.⁠ ⁠*Tratamento de Erros*
   - Sempre verifique valores de erro retornados
   - Use mensagens de erro descritivas
   - Considere centralizar a lógica de tratamento de erros

2.⁠ ⁠*Concorrência*
   - Evite compartilhamento de dados mutáveis entre goroutines
   - Prefira comunicação por canais ao invés de sincronização com mutex
   - Sempre feche canais após o uso
   - Evite vazamentos de goroutines

3.⁠ ⁠*Performance*
   - Limite o número de goroutines em sistemas de produção
   - Considere o uso de buffered channels para melhor desempenho
   - Implemente timeouts para evitar bloqueios indefinidos

## Extensões Futuras

Para participantes interessados em expandir o projeto após o workshop:

•⁠  ⁠Implementação de worker pools para limitar concorrência
•⁠  ⁠Persistência de resultados em banco de dados
•⁠  ⁠Interface web para visualização de resultados
•⁠  ⁠Verificação periódica e alertas
•⁠  ⁠Testes automatizados para código concorrente

---

Este plano de workshop foi desenvolvido para introduzir conceitos de concorrência em Go de forma progressiva e prática, adequado para estudantes de engenharia nos primeiros anos de formação.