# Aprenda Go

![Go Version](https://img.shields.io/badge/Go-1.24.2-blue.svg)

Este repositório contém minha jornada de estudos da linguagem Go (Golang), desde conceitos básicos até tópicos avançados como concorrência, interfaces e paralelismo. O material está organizado em capítulos progressivos com exercícios práticos e exemplos de código.

## Fonte

Material baseado em diversos recursos de aprendizado Go, incluindo:

##### Principal 
- Curso: Aprenda Go (Ellen Korbes)
- Documentação oficial do Go
- Go by Example

**Canais extras**
- Full Cycle
- HunCoding
- Aprenda Golang

### Temas abordados

#### **Capítulo 3 - Fundamentos da Linguagem**

- `ex01/` - Declaração de variáveis com operador curto (`:=`)
- `ex02/` - Variáveis globais com `var`
- `ex03/` - Tipos personalizados e conversão de tipos
- `ex04/` - Valores zero (zero values)
- `ex05/` - Conversão entre tipos

#### **Capítulo 4 - Tipos Booleanos e Operadores**

- `ex01/` - Tipo booleano e operadores relacionais (`==`, `!=`, `<`, `>`, `<=`, `>=`)
- `ex02/` - Operadores lógicos e expressões booleanas

#### **Capítulo 5 - Constantes e Operadores Bit a Bit**

- `ex01/` - Declaração e uso de constantes
- `ex02-part1/` a `ex02-part3/` - Operadores relacionais e tipos diferentes
- `ex03/` - Constantes tipadas e não tipadas
- `ex04/` - Operadores bit a bit (shift operators)
- `ex05/` - Mais operadores bit a bit
- `ex06/` - Constante `iota` para enumerações

#### **Capítulo 7 - Estruturas de Controle**

- `ex01/` - Loop `for` básico (1 a 1000)
- `ex02/` - Loop `for` com condições específicas
- `ex03/` - Loop estilo "while"
- `ex04/` - Loop infinito com `break`
- `ex05/` - Operador módulo (`%`)
- `ex06/` - Loops aninhados
- `ex07/` - Estruturas condicionais `if/else`
- `ex08/` - Switch cases
- `ex09/` - Switch sem expressão
- `ex-surpresa/` - Valores ASCII, decimal, hexadecimal e Unicode

#### **Capítulo 8 - Estruturas de Dados**

- `slice/` - Trabalhando com slices
- `map/` - Mapas (chave-valor)

#### **Capítulo 9 - Arrays, Slices e Maps**

- `ex01/` - Arrays básicos
- `ex02/` - Slices e iteração
- `ex03/` - Slicing (fatiamento)
- `ex04/` - Função `append`
- `ex05/` - Manipulação avançada de slices
- `ex06/` - Make e capacidade de slices
- `ex07/` - Slices multidimensionais
- `ex08/` - Maps básicos
- `ex09/` - Maps com slices como valores
- `ex010/` - Deletando elementos de maps

#### **Capítulo 10 - Structs**

- `struct/test1/` - Structs básicas e inicialização
- `struct/test2/` - Structs aninhadas e composição

#### **Capítulo 11 - Funções**

- `ex01/` - Declaração básica de funções
- `ex02/` - Funções com múltiplos parâmetros
- `ex03/` - Funções com retorno múltiplo
- `ex04/` - Funções variádicas (`...`)

#### **Capítulo 12 - Funções Avançadas**

- `func/teste1/` - Funções como tipos
- `func/teste2/` - Funções como parâmetros
- `func/teste3-interface/` - Interfaces básicas
- `func/teste4-interface/` - Interfaces e polimorfismo
- `func/teste5/` - Funções anônimas e expressões
- `func/teste6/` - Funções que retornam funções
- `func/teste7/` - Callbacks
- `func/teste8/` - Closures
- `func/teste9/` - Recursão

#### **Capítulo 13 - Funções na Prática**

- `ex01/` - Funções variádicas básicas
- `ex02/` - Funções com slices
- `ex03/` - Defer
- `ex04/` - Métodos
- `ex05/` - Interfaces na prática
- `ex06/` - Funções anônimas
- `ex07/` - Expressões de função
- `ex08/` - Retornando funções
- `ex09/` - Callbacks
- `ex010/` - Closures na prática
- `ex011/` - Recursão

#### **Capítulo 14 - Ponteiros**

- `ponteiros-procedural/` - Estilo procedural com ponteiros
- `ponteiros-poo/` - Estilo orientado a objetos

#### **Capítulo 15 - Ponteiros na Prática**

- `ex01/` - Conceitos básicos de ponteiros
- `ex02/` - Ponteiros e funções (estilo procedural)
- `ex02-poo/` - Ponteiros com métodos (estilo POO)

#### **Capítulo 16 - Bibliotecas Padrão**

- `bcrypt/` - Criptografia de senhas com bcrypt
- `json/1-marshal/` - Serialização JSON
- `json/2-unmarshal/` - Deserialização JSON
- `json/1-decode-encode/` - Encoding/Decoding JSON
- `sort/basic/` - Ordenação básica
- `sort/sort.Sort/` - Interface sort.Sort

#### **Capítulo 17 - Ordenação e Sort**

- `ex01/` - Sort com tipos básicos
- `ex02/` - Sort customizado
- `ex03/` - Sort com structs
- `ex04/` - Sort com strings e inteiros
- `ex05/` - Implementação completa da interface Sort

#### **Capítulo 18 - Concorrência Básica**

- `1-concorrencia-paralelismo/` - Conceitos de concorrência vs paralelismo
- `2-concorrencia-na-pratica/` - Race conditions
- `3-mutex/` - Mutexes para sincronização
- `4-atomic/` - Operações atômicas

#### **Capítulo 19 - Organização de Código**

- `pkg1/` - Criando pacotes
- `pkg2/` - Usando pacotes externos

#### **Capítulo 20 - Concorrência Avançada**

- `ex01-part1/` - Goroutines básicas
- `ex01-part2/` - Goroutines complexas
- `ex02/` - Interfaces com métodos de ponteiro
- `ex03-part1/` - Race conditions demonstradas
- `ex03-part2/` - Soluções para race conditions
- `ex04-part1/` - Mutex básico
- `ex04-part2/` - Mutex avançado
- `ex05/` - Operações atômicas
- `ex06/` - Informações do sistema (OS/ARCH)

#### **Capítulo 21 - Channels (Canais)**

- `1-unbuffered-channels/` - Canais não bufferizados
- `2-buffered-channels/` - Canais bufferizados
- `3-canais-direcionais/` - Canais direcionais (send-only/receive-only)
- `4-canais-direcionais/` - Mais canais direcionais
- `5-range-close/` - Range e close em canais
- `6-select/ex1/` - Select básico
- `6-select/ex2/` - Select com quit channel
- `6-select/ex3/` - Select com múltiplos canais
- `7-comma-ok-basic/` - Comma ok idiom básico
- `8-comma-ok/` - Comma ok avançado
- `9-convergencia-basic/` - Padrão fan-in básico
- `10-convergencia/` - Padrão fan-in avançado
- `teste-channels/` - Testes diversos com channels

## Como Usar Este Repositório

### Pré-requisitos

- Go 1.24.2 ou superior instalado
- Editor de código (recomendado: VS Code com extensão Go)

### Executando os Exemplos

1. Clone o repositório:

```bash
git clone https://github.com/j-ordep/go-exercises.git
cd aprenda-go
```

2. Execute qualquer exemplo:

```bash
# Exemplo: executar o primeiro exercício do capítulo 3
cd cap3/ex01
go run ex01.go

# Exemplo: executar um exercício de channels
cd cap21/1-unbuffered-channels
go run main.go
```

3. Para exercícios que usam dependências externas:

```bash
# O projeto já tem go.mod configurado
go mod tidy  # baixa as dependências
cd cap16/bcrypt
go run ex.go
```

## Progressão de Aprendizado

### Iniciante (Capítulos 3-8)

- ✅ Variáveis e tipos
- ✅ Operadores e expressões
- ✅ Estruturas de controle (if, for, switch)
- ✅ Arrays, slices e maps

### Intermediário (Capítulos 9-16)

- ✅ Structs e composição
- ✅ Funções e métodos
- ✅ Interfaces
- ✅ Ponteiros
- ✅ Bibliotecas padrão

### Avançado (Capítulos 17-21)

- ✅ Algoritmos de ordenação
- ✅ Concorrência e paralelismo
- ✅ Goroutines e WaitGroups
- ✅ Channels e patterns de comunicação
- ✅ Sincronização (Mutex, Atomic)

## Principais Conceitos Cobertos

### Fundamentos

- **Variáveis**: Declaração, tipos, zero values
- **Constantes**: Tipadas, não tipadas, iota
- **Operadores**: Aritméticos, relacionais, lógicos, bit a bit

### Estruturas de Dados

- **Arrays**: Estruturas de tamanho fixo
- **Slices**: Arrays dinâmicos
- **Maps**: Estruturas chave-valor
- **Structs**: Tipos personalizados

### Programação Funcional

- **Funções**: Declaração, parâmetros, retornos múltiplos
- **Funções Variádicas**: Parâmetros variáveis (`...`)
- **Funções de Primeira Classe**: Como valores e parâmetros
- **Closures**: Captura de escopo
- **Callbacks**: Funções como argumentos

### Orientação a Objetos

- **Métodos**: Funções com receivers
- **Interfaces**: Contratos de comportamento
- **Polimorfismo**: Múltiplas implementações
- **Composição**: Embedding de structs

### Concorrência

- **Goroutines**: Threads leves
- **Channels**: Comunicação entre goroutines
- **Select**: Multiplexação de canais
- **Patterns**: Fan-in, fan-out, worker pools
- **Sincronização**: Mutex, WaitGroups, operações atômicas

### Bibliotecas Padrão

- **JSON**: Marshaling/Unmarshaling
- **Crypto**: Hashing de senhas (bcrypt)
- **Sort**: Algoritmos de ordenação
- **Runtime**: Informações do sistema

## Tópicos Especiais

### Race Conditions e Sincronização

- Demonstração de condições de corrida
- Soluções com Mutex
- Operações atômicas
- Patterns de sincronização

### Patterns de Channels

- **Unbuffered vs Buffered**: Diferenças e casos de uso
- **Directional Channels**: Send-only e receive-only
- **Select Statement**: Multiplexação não-bloqueante
- **Comma Ok Idiom**: Verificação de estado de canais
- **Fan-in/Fan-out**: Patterns de distribuição

### Boas Práticas

- Organização de código em pacotes
- Tratamento de erros
- Testes e benchmarks
- Performance e otimização

---