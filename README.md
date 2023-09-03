# Anotações de Estudos - Go Lang

## Conceitos Básicos
- Sobre a palavra-chave *package* e projetos GO que possuam *imports* a serem realizados no arquivo principal
    - A palavra-chave package é utilizada para declarar a qual pacote um arquivo pertence. Serve para organizar um projeto em diferentes arquivos.
    - [Exemplo da estrutura de um projeto](https://github.com/git-mare/go/blob/main/exemplos/exemplo-package/)
    - [Arquivo principal da estrutura de um projeto](https://github.com/git-mare/go/blob/main/exemplos/exemplo-package/main.go)
    - Para criar um projeto na própria máquina, é necessário inicia-los no diretório em que o Go foi instalado.
    - Para descobrir onde foi instalado, digite no terminal `go env GOPATH`. Navegue até o diretório retornado pelo terminal e crie uma pasta chamada `src`. Nela, crie a pasta em que seu projeto será desenvolvido.
    - Exemplo: `C:\Users\[USERNAME]\go\src\exemplo-package`
    - Após estar no diretório do projeto, digite no terminal `go mod init exemplo-package`.
    - Este comando irá criar um arquivo que vai gerenciar a lista de pacotes necessárias para que o projeto possa ser executado.
    - Depois de todos estes passos, será possível executar o arquivo main sem problemas, utilizando no terminal o comando `go run main.go` (ou o nome do arquivo principal que você criou para seu projeto).
    
- Variáveis
    - Uma variável não pode ser declarada sem um propósito. Caso contrário o código não roda.
    - Não pode haver dois valores em apenas uma variável;
    - Caso haja este cenário, uma solução para resolver, é utilizar o underscore (`_`) para que o código não dê problema.
    - [Exemplo](https://github.com/git-mare/go/blob/main/exemplos/variaveis/main.go)
    - No exemplo acima, o código simplesmente pega o valor "Hello, World!" e joga fora, pois está sendo armazenado em `_`.
    - Uma variável não pode ser utilizada antes de ser declarada.

- Operadores
    - O operador de declaração em um escopo é o `:=`. Ele não pode ser utilizado fora de um codeblock.
        - Exemplo: *n1 := 10*
    - O operador de atribuição é o `=`.
        - Exemplo: *n1 = 15*
    - No primeiro exemplo, `n1` é declarado contendo o valor 10, enquanto no segundo, é atribuído outro valor a variável `n1`.
    - A declaração de uma variável fora de um *codeblock* é feita da seguinte maneira:
        - *var n1 = 5*
    - O operador de comparação é o `==`.
        - Exemplo: *n2 := 10 == 10*
        - Saída: `true`.
    - [Exemplo completo em código](https://github.com/git-mare/go/blob/main/exemplos/operadores/main.go)
- Tipos de dados
  - Os tipos de dados primitivos em GO são *int*, *string* e *bool*.
  - Os tipos de dados compostos são *slice*, *array*, *struct* e *map*.
  - Quando o tipo de uma variável é declarada, ele não pode ser modificado. Se uma variável for declarada *int*, ela receberá valores *int* e nada mais.
    
- Pacote fmt
  - `fmt.Print()`: somente imprime o que foi pedido.
  - `fmt.Println()`: imprime o que foi pedido e insere uma linha nova após a impressão.
  - `fmt.Printf()`: utilizado para imprimir valores que possuam variaveis junto de string. [Exemplo](https://github.com/git-mare/go/blob/main/exemplos/fmt/main.go?plain=1#L8&L9)

  - `fmt.Sprint()`: O valor vai ser salvo em uma string (não será impresso na tela). Poderá ser armazenado dentro de uma variável e utilizado num `fmt.Println()`. [Exemplo](https://github.com/git-mare/go/blob/main/exemplos/fmt/main.go?plain=1#L12-L15)
    
  - `fmt.Sprintf()`:
  - `fmt.Fprint()`:
## Recursos Adicionais
- [Layout padrão de projetos em Go](https://github.com/golang-standards/project-layout/blob/master/README_ptBR.md)
- [Guia para GO](https://github.com/caioreix/go4noobs#go4noobs)
- Livro: "The Go Programming Language" - Alan A. A. Donovan, Brian W. Kernighan
- [Site Oficial](https://golang.org/)

# Tópicos ainda não vistos
- Estrutura de Controle `for`
- Arrays e Slices

### Ponteiros e Estruturas
- Ponteiros e Referências
- Structs e Tipos Embutidos
- Métodos em Structs

### Interfaces e Goroutines
- Interfaces e Composição
- Concorrência com Goroutines
- Canais (Channels) e Goroutines

### Pacotes e Testes
- Organização em Pacotes
- Testes e Benchmarking
- Documentação em Go (godoc)

### Exemplos de Aplicações
- Aplicação de Linha de Comando
- Web Services com o Pacote "net/http"
