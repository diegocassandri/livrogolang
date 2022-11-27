# Ponteiros em GO

O operador & fornece o endereço de uma variável, e o operador * recupera a variável a qual o ponteiro se refere.

um ponteiro é o endereço de uma variável. Desse, modo, um ponteiro é o local em que um valor é armazenado. Nem todo valor tem um endereço, mas toda variável tem. Com um ponteiro, podemos ler ou atualizar o valor de uma variável indiretamente, sem usar ou sequer saber o nome da variável, se é que ela tem um nome.

Se uma variável for declarada como var x int, a expressão &x (Endereço de x) fornece um ponteiro para uma variável inteira, isto é, um valor do tipo *int, lido como "ponteiro para int". Se esse valor se chamar p, dizemos que "p aponta para x" ou de modo equivalente "p contém o endereço de x". A variável a qual p aponta é escrita como ***p**, essa expressão fornece o valor dessa variável, que é um int, mas como *p representa uma variável, ela também pode aparecer do lado esquerdo de uma atribuição, caso em que a atribuição atualiza a variável

```go
x := 1
p := &x //p, do tipo *int, aponta para x
 
fmt.Println(*p) //1
*p = 2 // equivalente a x = 2
fmt.Println(x) //2
```
