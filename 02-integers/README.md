Bom inteiros são o que você já pode imaginar, números kkkkkkk

Aqui vamos criar uma função que vai somar dois números pra testar este tipo primitivo. Então nós vamos criar um arquivo de teste chamado adder_test.go.

Obs: os arquivos Go devem ter apenas um pacote por diretório. Tenha certeza que os arquivos estão sendo organizados corretamente em suas devidas pastas.

A minha está desta forma:

```go
golang
    |
    |-> 01-hello-world
    |    |- hello.go
    |    |- hello_test.go
    |    |- README.md
    |
    |-> 02-integers
    |    |- adder_test.go
    |
    |- go.mod
```

Estou adicionando README’s com estes conteúdos que eu estou entendendo da doc do nosso mano quii (Chris James)

## Escrevendo o teste antes

```go
package integers

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

```

Agora, em vez de usar o %q na função Errorf para carregar as variáveis na string, como o tipo mudou para integer vamos usar o %d.

O nome do pacote agora também não é mais main, estamos usando integers que faz alusão ao que o código está trabalhando.

## Vendo o teste falhar

Ao rodar o teste é exibido:

```go
# hello/02-integers [hello/02-integers.test]
./adder_test.go:6:9: undefined: Add
FAIL	hello/02-integers [build failed]
```

## Escrevendo o minimo de código pro teste compilar e verificar o erro de asserção no output

```go
package integers

func Add(x, y int) int {
	return 4
}
```

Como estamos retornando o número 4 de maneira fixa, podemos escrever um outro teste para falhar se for seguir a risca do TDD mas fica meio em loop isso, não vale a pena.

Por hora, vamos apenas corrigir a função:

```go
package integers

func Add(x, y int) int {
	return x + y
}
```

A gente pode criar documentação das funções com comentários acima da declaração da função e eles vão aparecer no Go Doc que podemos rodar local na nossa máquina:

```go
package integers

// Add recebe dois parâmetros de tipo inteiro e retorna a soma em uma variável inteira
func Add(x, y int) int {
	return x + y
}
```

## Exemplos de execução nos testes

Se você quiser deixar ainda mais claro o que as funções do seu package fazem, é possível definir funções de exemplo nos nossos arquivos de teste.

Elas devem começar com o prefixo `Example` (assim como as funções de teste começam com `Test`) e devem estar em um arquivo de teste `_test.go` :

```go
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
```

Fazer isso deixa o código ainda mais acessível, estas funções também vão para o Go Docs e executando o teste com a flag -v:

```go
=== RUN   TestAdder
--- PASS: TestAdder (0.00s)
=== RUN   ExampleAdd
--- PASS: ExampleAdd (0.00s)
PASS
ok  	golang-with-tests/02-integers	0.002s
```

O exemplo também é executado.

O exemplo é executado toda vez que o arquivo de teste for compilado e aquele comentário com `// Output: 6` está garantindo essa execução. Se a gente remover o comentário, o teste não será mais executado.

## Resumo do aprendizado até aqui

- Mais práticas com TDD;
- Integers;
- Escrevendo documentação de uma forma melhor;
- Exemplos de como usar o código e como ele é executado com os nossos testes