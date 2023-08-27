# Anotações de Estudos - Go Lang

## Conceitos Básicos
- Keyword package
    - A palavra-chave package é utilizada para declarar a qual pacote um arquivo pertence. Serve para organizar o código em diferentes arquivos.
    
- Variáveis
    - Uma variável não pode ser declarada sem um propósito. Caso contrário o código não roda.
    - Da mesma forma, não pode haver dois valores em apenas uma variável;
    - Caso haja este cenário, uma solução para resolver, é utilizar o underscore (_) para que o código não dê problema.
    - Exemplo:
        _, erros := fmt.Println("Hello, World!", "Mensagem de erro")
        fmt.Println(erros)
    - No exemplo acima, o código simplesmente pega o valor "Hello, World!" e joga fora, pois está sendo armazenado em _.
    - Uma variável não pode ser utilizada antes de ser declarada.

- Operadores
    - O operador de declaração em um escopo é o :=. Ele não pode ser utilizado fora de um codeblock.
        - Exemplo: n1 := 10
    - O operador de atribuição é o =.
        - Exemplo: n1 = 20
    - No primeiro exemplo, n1 é declarado contendo o valor 10, enquanto no segundo, é atribuído outro valor a variável n1.
    - A declaração de uma variável fora de um codeblock é feita da seguinte maneira:
        - var n1 = 10
    - O operador de comparação é o ==.
        - Exemplo: n1 := 10 == 10
        - Saída: true



- Tipos de Dados
- Estruturas de Controle (if, for, switch)
- Funções e Escopo
- Arrays e Slices

## Ponteiros e Estruturas
- Ponteiros e Referências
- Structs e Tipos Embutidos
- Métodos em Structs

## Interfaces e Goroutines
- Interfaces e Composição
- Concorrência com Goroutines
- Canais (Channels) e Goroutines

## Pacotes e Testes
- Organização em Pacotes
- Testes e Benchmarking
- Documentação em Go (godoc)

## Exemplos de Aplicações
- Aplicação de Linha de Comando
- Web Services com o Pacote "net/http"

## Recursos Adicionais
- Livro: "The Go Programming Language" - Alan A. A. Donovan, Brian W. Kernighan
- Site Oficial: https://golang.org/
