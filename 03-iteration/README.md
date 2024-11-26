# 03 - Iteration

Pra fazer iterações em Go tudo que você precisa é da funcionalidade for. No Go, não existe while, do until, apenas o for para todos os casos, o que pode ser uma coisa boa a final, precisamos aprender somente este método de iteração.

Agora vamos criar um teste que irá repetir um caractere 5 vezes.

Aqui vamos usar o mesmo workflow que estamos acostumados dos outros capítulos

## Escrevendo o teste antes

```go
package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
```

Rodando o test ou pelo LSP do VSCode vamos ver que Repeat é undefined.

## Escrevendo o mínimo para validar o teste

```go
package iteration

func Repeat(character string) string {
	return ""
}
```

Com isso temos o retorno do teste igual a:

```bash
--- FAIL: TestRepeat (0.00s)
    iteration_test.go:10: expected "aaaaa" but got ""
FAIL
exit status 1
FAIL	golang-with-tests/03-iteration	0.001s
```

## Agora vamos escrever o código para o teste passar, de fato

O for em Go é bem similar as outras linguagens, com exceção de:

- Não é escrita com parêntesis em volta dos 3 componentes;
- {} sempre são necessárias

```go
package iteration

func Repeat(character string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated = repeated + character
	}
	return repeated
}

```

Até agora nós inicializamos e atribuímos valores em variáveis através do `:=` mas neste caso a gente declarou a variável repeated antes do for, mas inicializamos ela dentro dele quando atribuímos valor.

## Aquela refatorada

No Go também temos o “the Add AND assignment operator” → `+=`

Basicamente ele soma o operador do lado direito ao operador do lado esquerdo e atribui o resultado ao lado esquerdo.

```go
package iteration

const repeatCount = 5

func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
```

Desta forma não precisamos repetir a variável na frente do igual.

Uma boa também é definir o quanto irá repetir o laço sem usar magic numbers 😀

## Benchmarking

A suite de testes do Go Lang possui uma função de benchmark para termos noção de quanto tempo leva para o nosso código executar N vezes.

O próprio motor de teste do Go determina quantas vezes o código será executado e no fim, ele exibe o tempo e quantas vezes o código foi executado

Dando nome aos bois a gente vai usar o testing.B do Go para ter acesso ao b.N que é esse recurso que vai determinar quantas vezes o código será executado e capturar os logs pra gente.

Pra ser bem sincero isso é a unica coisa que eu acho meio estranho no Go até agora, estas tipagens de teste 🧙🏽‍♂️

```go
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
```

A função é bem simples, basta a gente implementar um for onde o i vai iterar até b.N e chamar a execução da nossa função

Pra rodar o benchmarking junto com o teste é só digitar:

```bash
go test -bench=.
```

E temos o seguinte resultado aqui na minha máquina:

```bash
goos: linux
goarch: amd64
pkg: golang-with-tests/03-iteration
cpu: 12th Gen Intel(R) Core(TM) i5-1235U
BenchmarkRepeat-12    	 6270073	       169.0 ns/op
PASS
ok  	golang-with-tests/03-iteration	1.260s
```

Demorou 169 nanosegundos para rodar 6270073 vezes

Niceeeee

## Pra finalizar

A gente pode escrever a função de exemplo aqui no nosso teste também e o arquivo de testes fica desta forma:

```go
package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

func ExampleRepeat() {
	wordToRepeat := "j"
	fmt.Println(Repeat(wordToRepeat))
	// Output: jjjjj
}

```

## Resumo do que foi aprendido

- Mais praticas de TDD;
- for
- Como escrever benchmarks