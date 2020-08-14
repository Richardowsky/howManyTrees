# How Many Trees

## Prerequisites:
You need to install - [go 1.13](https://golang.org/dl/)

## How to test:
1. Clone this repo: `git clone https://github.com/Richardowsky/howManyTrees.git`
2. `make test`

## How to use:
```go
package example

import howManyTrees "howManyTrees/src"

func example()  {
var n, h = 3, 3
	howManyTrees.Solution(n, h)
}

```

Алгоритм:
- В функции Solution каждый раз проверем n и h
- Проверем каждый раз повторяющуюся ветку
- Если проверки ложны, создаем цикл
- В этом цикле разделяем главное дерево на множество маленьких деревьев с рекурсией
- Разделение результатов вычисляет общее количество деревьев.

Алгоритм решения = O(N*log(N))