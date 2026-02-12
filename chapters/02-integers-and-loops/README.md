# Глава 2: Целые числа и циклы

---

**← [Глава 1: Hello, World](../chapters/01-hello-world/README.md)** | **[Глава 3: Структуры и методы](../chapters/03-structures-and-methods/README.md) →**

---

В этой главе мы изучим работу с целыми числами и циклами в Go. Мы будем использовать методологию TDD (Test-Driven Development) для написания кода.

## Что вы узнаете

- Как работать с целыми числами в Go
- Как использовать циклы в Go
- Сравнение с Java: как те же концепции реализуются в Java
- Практические примеры и упражнения

## Целые числа в Go

В Go есть несколько типов для представления целых чисел:

- `int`: целое число со знаком, размер зависит от платформы (обычно 64 бита)
- `int8`, `int16`, `int32`, `int64`: целые числа со знаком с фиксированным размером
- `uint`: целое число без знака, размер зависит от платформы
- `uint8`, `uint16`, `uint32`, `uint64`: целые числа без знака с фиксированным размером

### Пример: Сложение двух чисел

Давайте напишем функцию, которая складывает два целых числа.

#### Go

```go
package main

import "fmt"

func Add(a, b int) int {
    return a + b
}

func main() {
    result := Add(2, 3)
    fmt.Println(result) // Вывод: 5
}
```

#### Java

```java
public class Main {
    public static int add(int a, int b) {
        return a + b;
    }

    public static void main(String[] args) {
        int result = add(2, 3);
        System.out.println(result); // Вывод: 5
    }
}
```

## Циклы в Go

В Go есть только один тип цикла: `for`. Он может использоваться для различных видов итераций.

### Пример: Цикл for

Давайте напишем функцию, которая выводит числа от 1 до 5.

#### Go

```go
package main

import "fmt"

func PrintNumbers() {
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
    }
}

func main() {
    PrintNumbers()
}
```

#### Java

```java
public class Main {
    public static void printNumbers() {
        for (int i = 1; i <= 5; i++) {
            System.out.println(i);
        }
    }

    public static void main(String[] args) {
        printNumbers();
    }
}
```

## Сравнение Go и Java

| Характеристика | Go | Java |
|----------------|-----|------|
| **Целые числа** | `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64` | `byte`, `short`, `int`, `long` |
| **Циклы** | Только `for` | `for`, `while`, `do-while` |
| **Инициализация переменной в цикле** | `for i := 1; i <= 5; i++` | `for (int i = 1; i <= 5; i++)` |
| **Бесконечный цикл** | `for { ... }` | `while (true) { ... }` |

## Практическое упражнение

Попробуйте написать функцию, которая вычисляет факториал числа, используя цикл `for`.

### Go

```go
package main

import "fmt"

func Factorial(n int) int {
    result := 1
    for i := 1; i <= n; i++ {
        result *= i
    }
    return result
}

func main() {
    result := Factorial(5)
    fmt.Println(result) // Вывод: 120
}
```

### Java

```java
public class Main {
    public static int factorial(int n) {
        int result = 1;
        for (int i = 1; i <= n; i++) {
            result *= i;
        }
        return result;
    }

    public static void main(String[] args) {
        int result = factorial(5);
        System.out.println(result); // Вывод: 120
    }
}
```

## Итоги главы

- ✅ Вы узнали, как работать с целыми числами в Go
- ✅ Вы узнали, как использовать циклы в Go
- ✅ Вы сравнили концепции с Java
- ✅ Вы решили практическое упражнение

## Что дальше?

В следующей главе мы изучим **структуры и методы**, где узнаем, как создавать пользовательские типы данных и методы для них.

---

**← [Глава 1: Hello, World](../chapters/01-hello-world/README.md)** | **[Глава 3: Структуры и методы](../chapters/03-structures-and-methods/README.md) →**