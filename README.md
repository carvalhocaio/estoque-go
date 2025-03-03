# Métodos em Go

Go permite definir métodos para tipos personalizados dentro do bloco do pacote:

```go
type Person struct {
    FirstName string
    LastName  string
    Age       int
}

func (p Person) String() string {
    return fmt.Sprintf("%s %s, idade %d", p.FirstName, p.LastName, p.Age)
}
```

## Definição de Métodos

- Métodos são semelhantes a funções, mas incluem um **receptor**, que aparece antes do nome do
método.

- O receptor deve ter um nome curto, geralmente uma abreviação do tipo (ex.: `p` para `Person`), e 
**não se usa** `this` ou `self`.

- Métodos só podem ser definidos no **nível do pacote**.

## Convenções e Restrições

- Sem sobrecarga de métodos: um tipo não pode ter dois métodos com o mesmo nome.

- Métodos não podem ser adicionados a tipos externos: é necessário estar no mesmo pacote
onde o tipo foi definido.

- Código organizado: recomenda-se manter a definição do tipo e seus métodos juntos.

Essas práticas reforçam a clareza e simplicidade do Go, evitando padrões tradicionais de
orientação a objetos.

# Logging em Go

O **logging** é essencial para monitoramento e rastreamento de erros em sistemas de estoque. A escolha da biblioteca correta impacta desempenho e escalabilidade. Aqui estão três opções populares:

## 1. Log (Pacote Padrão)

O pacote `log` é simples e ideal para aplicações menores.

Funcionalidades:

- Registro de mensagens com níveis básicos (Info, Error, Fatal).

- Suporte para formatação simples.

- Escrita de logs em arquivos e console.

```go
package main

import (
    "log"
)

func main() {
    log.Println("Log de informação")
    log.Fatal("Log de erro")
}
```

## 2. Logrus

O Logrus oferece logs estruturados e flexibilidade na saída de logs.

Funcionalidades:

- Formatos estruturados (JSON, texto).

- Suporte a múltiplos destinos (console, arquivos, monitoramento).

- Hooks para integração com ferramentas externas.

- Níveis de log personalizáveis.

```go
package main

import (
    "github.com/sirupsen/logrus"
)

func main() {
    log := logrus.New()
    log.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})

    log.Info("Log de informação")
    log.Warn("Log de aviso")
    log.Error("Log de erro")
}
```

## 3. ZAP

O Zap, criado pelo Uber, é otimizado para alta performance e eficiência.

Funcionalidades:

- Alto desempenho e baixa latência.

- Logs estruturados para sistemas de grande escala.

- Recomendado para aplicações de alto tráfego.

```go
package main

import (
    "go.uber.org/zap"
)

func main() {
    logger, _ := zap.NewProduction()
    defer logger.Sync()

    logger.Info("Log de informação", zap.String("user", "clienteX"))
    logger.Warn("Log de aviso")
    logger.Error("Log de erro")
}
```

A escolha depende do projeto:

- Pacote log: Simples e direto.

- Logrus: Flexibilidade e integração com ferramentas externas.

- Zap: Melhor desempenho para sistemas de grande escala.

---