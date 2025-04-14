# Busca binária

* A pesquisa binária é muito mais rápida que a pesquisa simples.
* O(log n) é mais rápido do que O(n), e O(log n) fica ainda mais rápido conforme os elementos da lista aumentam.
* A rapidez de um algoritmo não é medida em segundos.
* O tempo de execução de um algoritmo é medido por meio de seu crescimento.
* O tempo de execução de um algoritmo é expresso na notação Big O.

# Big O notation
* O Big O é uma notação especial que diz o quão rápido é um algoritmo (tempo de execução)
* A rapidez de um algoritmo não é medida em segundos, mas pelo crescimento do número de operações

# Exemplos de tempo de execução Big 0
* O(log n): tempo logarítmo (pesquisa binária)
* O(n): tempo linear (pesquisa simples)
* O(n * log n): algoritmo rápido de ordenação (quicksort)
* O(n^2): algoritmo lento de ordenação (ordenação por seleção)
* O(n!): caixeiro viajante (um algoritmo bastante lento)

# Exercicios

1.1 Suponha que você tenha uma lista com 128 nomes e esteja fazendo uma pesquisa binária. Qual seria o número máximo de etapas que você levaria para encontrar o nome desejado? 


```bash
  Cálculo dos 128 nomes: 

  log 128 = 7

  totalizando 7 etapas até encontrar o número desejado.
```

1.2 Suponha que você duplique o tamanho da lista. Qual seria o número máximo de etapas agora? 

```bash
  Cálculo dos 256 nomes: 

  log 1256 = 8

  totalizando 8 etapas até encontrar o número desejado.
```

Forneça o tempo de execução para cada um dos casos a seguir em termos da notação Big O.

1.3 Você tem um nome e deseja encontrar o número de telefone para esse nome em um agenda telefônica. ```O(log n)```

1.4 Você tem um número de telefone e deseja encontrar o dono dele em uma agenda telefônica (Dica: Deve procurar pela agenda inteira). ```O(n)```

1.5 Você quer ler o número de cada pessoa na agenda telefônica. ```O(n)```

1.6 Você quer ler os números apenas dos nomes que começam com A. (Isso é complicado! Esse algoritmo envolve conceitos que serão abordados mais profundamente no Capítulo 4: Quicksort. Leia e responda - você ficará surpreso). ```O(n)```

