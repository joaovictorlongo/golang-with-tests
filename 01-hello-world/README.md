## Primeiro de tudo…

Estes posts retratam a minha experiência com o aprendizado da linguagem Go Lang.

Irei construir a mesma linha de raciocínio e seguir os mesmos passos do site: https://quii.gitbook.io/learn-go-with-tests. Aqui irei escrever com as minhas palavras.

Já falando sobre o primeiro passo no aprendizado da linguagem, é comum a gente começar fazendo o famoso “Hello world”. Dado que você já tenha os recursos da linguagem instalado na máquina né? (https://go.dev/doc/install)

E pra começar, eu criei uma pasta de estudos e dentro dela criei a pasta *golang.*

Agora, em relação aos arquivos do go, basicamente a extensão é o próprio nome xD

Então podemos criar o arquivo “hello.go” com o seguinte código:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, world")
}
```

Pra rodar, basta abrirmos no a pasta onde está o arquivo no terminal e digitar `go run hello.go`

## O que acontece aqui?

Basicamente, programas em Go são pacotes então definimos a primeira linha informando que este é o main.

‘fmt’ é uma standard lib do Go que por exemplo, possui a função Println para exibir informações no terminal/console.

A definição de funções se da pela palavra func o que é bem intuitivo, assim como a utilização de () e {} que é comumente visto em outras linguagens.

## Como que testa este código mesmo ele sendo incrivelmente simples?

Uma boa prática que é recomendada em código para auxiliar nos processos de teste é a separação de responsabilidades entre as funções, como por exemplo:

```go
package main

import "fmt"

func Hello() string {
	return "Hello, world"
}

func main() {
	fmt.Println(Hello())
}

```

Aqui eu criei uma função Hello que simplesmente retorna a string que a gente precisa e em vez de retornar a string direto na execução do pacote, ou seja, na função main, a gente usa uma função que faz isso, desta forma podemos tanto chamar a função na execução do pacote quanto chamar ela no arquivo de teste.

Vamos agora criar um arquivo de teste dentro da pasta do nosso hello-world chamado `hello_test.go`

O padrão de escrita dos testes em go segue esta nomenclatura de _test.go fechou?

```go
package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
```

Antes de explicar o arquivo de testes, o comando para executar este código seria o `go test` mas se a gente salvar o arquivo e rodar este comando vamos ver um erro de execução parecido com este:

```bash
go: go.mod file not found in current directory or any parent directory; see 'go help modules'
```

Isso acontece porque não existe nenhum módulo (go modules) inicializado na pasta, para que possa ser concentrada a versão, nome e as dependências do projeto e pra fazer isso é bem simples, basta a gente digitar no terminal:

```bash
	go mod init hello
```

Onde basicamente hello é o nome do nosso módulo.

Feito isso vamos conseguir rodar este arquivo de teste:

```bash
PASS
ok  	hello	0.001s
```

Niceeeeeee.

Resumindo, sempre que a gente for iniciar um novo projeto Go, precisamos rodar este comando de mod init para ‘inicializar’ ele.

## Voltando pros testes

O legal do Go é que você escreve os testes sem precisar de nenhum framework. A linguagem já ter isso de forma nativa já ajuda porque todos vão preferir usar o que a linguagem já tem em vez de importar mais uma lib no projeto e ter que aprender sobre ela e depender dela para tudo funcionar…

### Regras para escrita de testes

Não tem muito segredo e como dizem, é a mesma coisa que escrever funções, com algumas regrinhas…

- Precisa estar em um arquivo com o nome tipo xxx_test.go
- A função precisa começar com a palavra Test
- A função de teste só precisa e espera o argumento t * testing.T
- Pra gente usar o t * testing.T é necessário importar a lib ‘testing’, tipo o que a gente fez com o fmt…

Basicamente o nosso ‘t’ é o hook para tudo que precisamos fazer em relação aos testes no nosso arquivo de testes.

Olhando de novo para o arquivo fica muito easy de entender o que está acontecendo:

```go
package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
```

Importamos a lib testing, definimos a função **Test**Hello com o parâmetro **t *testing.T** e:

- Capturamos o retorno da função Hello do módulo na variável got;
- Definimos o que esperamos na variável want;
- Comparamos com um if statement se eles são diferentes e caso for true (diferentes) exibimos um erro com o hook da função de testes concatenando os valores das variáveis na string que esta sendo exibida.

## Alterando o Hello World para Hello ‘Você’

Agora que a gente tem um arquivo de teste fica mais fácil alterar o código de uma maneira um pouco mais segura.

Até agora a gente escreveu o código, depois escreveu o teste e foi alterando estes arquivos mas agora vamos seguir com uma das premissas do TDD. Escrever os testes e depois escrever o código.

Pra gente fazer isso na prática, basta alterar o que a gente espera do teste mesmo que isso for resultar em um erro:

```go
package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("João")
	want := "Hello, João"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
```

Aqui, eu alterei o teste pra receber a string ‘João’ na função Hello e também alterei o want para ‘Hello, João’.

No meu VSCode, por estar com a extensão do Go instalada, já é apresentado um erro do próprio LSP mas se rodarmos o teste esperando estes argumentos também vamos receber o seguinte erro:

```go
# hello [hello.test]
./hello_test.go:6:15: too many arguments in call to Hello
	have (string)
	want ()
FAIL	hello [build failed]
```

Em linguagens estaticamente tipadas tipo o Go é de suma importância a leitura dos erros, pois o compilador geralmente indica o que a gente precisa fazer para continuar ou resolver o problema que estamos enfrentando.

No caso do nosso Hello, precisamos adequar a função Hello para aceitar argumentos do tipo string. A boa é ir fazendo passo a passo para entender:

```go
func Hello(name string) string {
	return "Hello, world"
}
```

Se a gente pegar isso aqui e só rodar o teste novamente, agora, vamos receber do compilador que a função Hello não está recebendo argumentos suficientes para a execução:

```go
# hello [hello.test]
./hello.go:10:14: not enough arguments in call to Hello
	have ()
	want (string)
FAIL	hello [build failed]
```

Já se eu passar a string que a função espera receber, não terei mais erros no meu LSP mas se eu rodar os testes novamente com o código abaixo, também vou receber um erro:

```go
package main

import "fmt"

func Hello(name string) string {
	return "Hello, world"
}

func main() {
	fmt.Println(Hello("João"))
}

```

```go
--- FAIL: TestHello (0.00s)
    hello_test.go:10: got "Hello, world" want "Hello, João"
FAIL
exit status 1
FAIL	hello	0.001s
```

Agora ficou fácil, o compilador ta dizendo que a gente ta esperando o texto “Hello, João” mas recebemos “Hello, world”. 

Isso ta acontecendo porque o return da nossa função Hello ainda não está usando o parâmetro recebido pela função.

Se a gente concatenar corretamente o parâmetro na string e rodar o teste de novo vamos ver que ele vai passar sem problemas e olha só que massa, até este ponto, nosso código está devidamente testado 🥸

```go
package main

import "fmt"

func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("João"))
}
```

```go
PASS
ok  	hello	0.001s
```

*Neste ponto se estiver usando git para controlar o código feito, é uma boa fazer um commit*

## Constantes

Constantes em GoLang tem a escrita igual do javascript:

```go
const englishHelloPrefix = "Hello, "
```

Uma boa prática é armazenar strings que estão ‘hardcoded’ no código em constantes para um melhor entendimento do código, manutenções e melhorias. Podemos refatorar nosso código para ficar igual a:

```go
package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("João"))
}

```

Os testes devem continuar passando.

## Hello world de novo

O próximo passo é verificar se o parâmetro está vazio e se estiver, retornar a mensagem padrão de “Hello, World” em vez de só “Hello, “.

Primeiro a gente começa escrevendo um novo teste que irá falhar:

```go
package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("sayng hello to people", func(t *testing.T) {
		got := Hello("João")
		want := "Hello, João"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
```

E neste ponto a gente começa a utilizar uma função legal da lib de testing que são os subtests. Então para uma funcionalidade geralmente vamos ter mais de um teste e é útil agrupar eles em um único lugar para descrever diferentes cenários.

Da pra compartilhar código entre os testes também, escrevendo com subtests, desta maneira.

Agora que a gente tem o teste esperando o cenário ideal e agora que a gente viu que ele esta falhando conforme o esperado vamos refatorar a função principal pro teste receber os argumentos corretos:

```go
package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("João"))
}

```

Agora, se a variável name for igual a uma string vazia, World será atribuído, e os testes estarão passando 😀

```go
PASS
ok  	hello	0.001s
```

É bem importante que o teste e a descrição sejam claros quanto ao que está sendo testado, desta maneira garantiremos muito mais qualidade na entrega e nas manutenções em geral.

Refatorar não serve só para os arquivos de prod, coisas que usamos no desenvolvimento como por exemplo os testes, podem ser refatorados para uma melhora geral do produto.

No exemplo dos testes que fizemos ate agora podemos refatorar a asserção para uma função helper.

```go
package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("sayng hello to people", func(t *testing.T) {
		got := Hello("João")
		want := "Hello, João"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

```

A diferença na leitura é inegável, as vezes uma pequena função pode auxiliar e muito no nosso código. Em um arquivo cheio de testes, se ficarmos repetindo o mesmo if, a página vira um quebra cabeças em longas jornadas de leitura.

No parâmetro da função `assertCorrectMessage` a gente ta passando o testing.TB que é uma interface que satisfaz condições tanto do t.testing.T quanto do *testing.B, então você pode chamar funções helper de um teste ou um benchmark.

A chamada de t.Helper() é necessária pra falar pra suite de teste que esse método é um helper. Fazendo isso, se um teste nosso falhar ele vai falar a linha que falhou do teste e não da nossa função helper, por exemplo.

## Disciplina

Na doc o autor fala que é sempre legal manter a disciplina de:

- Escrever um teste
- Fazer o compilador aceitar
- Rodar o teste, ver que ele falhou e checar a mensagem de erro para saber como atuar
- Escrever código suficiente para solucionar o erro e fazer o teste passar
- Refatoração para melhorar a leitura e manutenção futura do código.

## Mais requirimentos no Hello World

Adicionando mais uma pequena camada de complexidade na função Hello, agora, vamos internacionalizar a parada…

Vamos receber um novo parâmetro informando o idioma que iremos exibir o hello world. Se a linguagem passada por parâmetro não for configurada vamos retornar o default English.

Agora a gente tem que implementar seguindo o TDD um pouco mais easy desta vez:

```go
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("João", "Spanish")
		want := "Hola, João"
		assertCorrectMessage(t, want, got)
	})
```

Como de costume, se estamos passando parâmetros a mais que a função espera, ao rodar os testes:

```go
# hello [hello.test]
./hello_test.go:19:25: too many arguments in call to Hello
	have (string, string)
	want (string)
FAIL	hello [build failed]
```

Adicionando o parâmetro na função, agora temos:

```go
# hello [hello.test]
./hello.go:15:20: not enough arguments in call to Hello
	have (string)
	want (string, string)
./hello_test.go:7:16: not enough arguments in call to Hello
	have (string)
	want (string, string)
./hello_test.go:13:16: not enough arguments in call to Hello
	have (string)
	want (string, string)
FAIL	hello [build failed]
```

Agora vamos corrigir passando o parâmetro vazio em todas as chamadas que não passamos o parâmetro language ainda, desta forma será retornado apenas o erro de asserção:

```go
--- FAIL: TestHello (0.00s)
    --- FAIL: TestHello/in_Spanish (0.00s)
        hello_test.go:21: got "Hola, João" want "Hello, João"
FAIL
exit status 1
FAIL	hello	0.002s
```

Bom agora vamos corrigir adicionando um if na função hello:

```go
package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "Spanish" {
		return "Hola, " + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("João", ""))
}

```

E então os testes devem passar.

É hora então de refatorar e seguir a mesma ideia que tivemos para o prefixo em inglês:

```go
package main

import "fmt"

const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == spanish {
		return spanishHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("João", ""))
}

```

Além disso o if para o espanhol tinha uma magic string que em casos de integers ou strings mais complexas atrapalham no processo de leitura do código, por isso, foi movido para uma constante, onde é facilmente identificada e manipulada, além, de manter uma unica escrita para o código em vez de misturar código com string.

## Agora em português

A mesma coisa:

- Escrevemos um teste que irá falhar;
- Analisamos o erro e corrigimos o código para satisfazer o que esperamos no teste;
- O teste passa;
- A gente refatora.

```go
	t.Run("in Portuguese", func(t *testing.T) {
		got := Hello("João", "Portuguese")
		want := "Olá, João"
		assertCorrectMessage(t, want, got)
	})
```

```go
--- FAIL: TestHello (0.00s)
    --- FAIL: TestHello/in_Portuguese (0.00s)
        hello_test.go:27: got "Olá, João" want "Hello, João"
FAIL
exit status 1
FAIL	hello	0.001s
```

```go
package main

import "fmt"

const spanish = "Spanish"
const portuguese = "Portuguese"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const portugueseHelloPrefix = "Olá, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == spanish {
		return spanishHelloPrefix + name
	}
	if language == portuguese {
		return portugueseHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("João", ""))
}

```

## Refatoração com switch

Como agora a gente tem vários ifs checando um valor específico, é comum de ver por ai o uso do statement switch. Podemos usar ele pra refatorar nosso código e manter ele mais legível e de fácil manutenção caso a gente queira adicionar mais linguagens no futuro:

```go
package main

import "fmt"

const spanish = "Spanish"
const portuguese = "Portuguese"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const portugueseHelloPrefix = "Olá, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case portuguese:
		prefix = portugueseHelloPrefix
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("João", ""))
}

```

Com estas mudanças a função pode ir ficando um pouco grande e um pouco difícil de lidar, com isso podemos mover a lógica de prefix para uma função específica:

```go
package main

import "fmt"

const spanish = "Spanish"
const portuguese = "Portuguese"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const portugueseHelloPrefix = "Olá, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case portuguese:
		prefix = portugueseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("João", ""))
}

```

Alguns conceitos novos:

- Na assinatura da nova função nos definimos um retorno com valor nomeado (prefix string)
- Isso cria uma variável prefix na nossa função
    - Para strings esta variável inicia como “” e para int como 0
    - Ela vai ser retornada mesmo se você escrever só o return ali
- Default no swtich determina o valor que será retornado caso não tenha dado match com nenhum case.
- A função que criamos pro prefix começa com a leta minúscula o que quer dizer que ela é uma função privada. Em Go, funções com a inicial maiúscula tem o escopo publico.
- Da pra agrupar as constantes em bloco mas pra leitura é melhor que cada uma tenha sua linha

## Resumindo

Da pra render pra caramba com um Hello, World 😀

### Aprendizados:

- Escrita de testes;
- Declaração de funções com argumentos e tipos de retorno;
- if, const e switch
- Declaração de variáveis e constantes

### O processo do TDD e porque esses passos são importantes

Agora falando como João aqui na minha opinião é que no mundo real, na maioria das vezes, você vai precisar escrever testes para os códigos que você escreve e particularmente escrever eles antes, além da segurança que da na hora de programar e verificar se esta certo é a satisfação do teste já estar pronto.

Querendo ou não é uma etapa super maçante e ela estar alinhada ao seu código te da mais assertividade e agilidade na hora das entregas, fechou?