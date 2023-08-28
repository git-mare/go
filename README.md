# Anotações de Estudos - Go Lang

## Conceitos Básicos
- Sobre a palavra-chave *package* e projetos GO que possuem arquivos a serem importados no arquivo principal
    - A palavra-chave package é utilizada para declarar a qual pacote um arquivo pertence. Serve para organizar um projeto em diferentes arquivos.
    - [Exemplo da estrutura de um projeto](https://github.com/git-mare/go/blob/main/exemplo-package/)
    - [Arquivo principal da estrutura de um projeto](https://github.com/git-mare/go/blob/main/exemplo-package/main.go)
    - Para criar um projeto na própria máquina, é necessário inicia-los no diretório em que o Go foi instalado.
    - Para descobrir onde foi instalado, digite no terminal: go env GOPATH. Navegue até o diretório retornado pelo terminal e crie uma pasta chamada src.
    - Após estar no diretório do projeto, digite no terminal: go mod init NomeDaPasta.
    - Depois de todos estes passos, será possível executar o arquivo main, que irá puxar os arquivos importados.
    
- Variáveis
    - Uma variável não pode ser declarada sem um propósito. Caso contrário o código não roda.
    - Não pode haver dois valores em apenas uma variável;
    - Caso haja este cenário, uma solução para resolver, é utilizar o underscore (_) para que o código não dê problema.
    - [Exemplo](https://github.com/git-mare/go/blob/main/variaveis/main.go)
    - No exemplo acima, o código simplesmente pega o valor "Hello, World!" e joga fora, pois está sendo armazenado em _.
    - Uma variável não pode ser utilizada antes de ser declarada.

- Operadores
    - O operador de declaração em um escopo é o :=. Ele não pode ser utilizado fora de um codeblock.
        - Exemplo: n1 := 10
    - O operador de atribuição é o =.
        - Exemplo: n1 = 15
    - No primeiro exemplo, n1 é declarado contendo o valor 10, enquanto no segundo, é atribuído outro valor a variável n1.
    - A declaração de uma variável fora de um codeblock é feita da seguinte maneira:
        - var n1 = 5
    - O operador de comparação é o ==.
        - Exemplo: n2 := 10 == 10
        - Saída: true
    - [Exemplo completo em código](https://github.com/git-mare/go/blob/main/operadores/main.go)



## Tópicos ainda não vistos
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
