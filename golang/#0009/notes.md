ponteiro é com se fosse uma variável que vai salvar na memoria.

ponteiro é uma referência de memória

com o ponteiro nós colocamos o endereço de memória onde a variável que estmoas usando está, como se fosse uma gaveta.

com ponteiros não passamos o contéudo da variável e, sim, onde ela está.

quando colocamos um \* na frente da varíavel em ponteiro ele desferencia e imprime de uma forma "legível", não com o endereço de memória.

    variavelthree = 100
    ponteiro = &variavelthree

    fmt.Println(variavelthree, *ponteiro)
