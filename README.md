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

---