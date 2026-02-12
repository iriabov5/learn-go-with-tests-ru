# Глава 1: Hello, World

---

**← [Настройка окружения](../../introduction/03-setup.md)** | **[Глава 2: Целые числа и циклы](../chapters/02-integers-and-loops/README.md) →**

---

Добро пожаловать в первую главу! Здесь мы напишем свою первую программу на Go и свой первый тест.

## Что вы узнаете

- Как создать тест в Go
- Как запустить тесты
- Базовый синтаксис Go
- Цикл TDD: красный → зелёный → рефакторинг

## Ваш первый тест

Создайте новый каталог для этой главы:

```bash
mkdir hello
cd hello
go mod init hello
```

### Шаг 1: Напишите тест

Создайте файл `hello_test.go`:

```go
package main

import "testing"

func TestHello(t *testing.T) {
    got := Hello()
    want := "Hello, World!"

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
```

### Разбор теста

- `package hello` — объявляет пакет. В Go тесты находятся в том же пакете, что и тестируемый код
- `import "testing"` — импортирует пакет для тестирования
- `func TestHello(t *testing.T)` — функция теста. Имя должно начинаться с `Test`
- `t *testing.T` — указатель на структуру тестирования, которая предоставляет методы для отчёта
- `t.Errorf` — сообщает об ошибке, но не останавливает тест

> **Примечание:** В Go имена функций и переменных, которые экспортируются (доступны извне пакета), должны начинаться с **заглавной буквы**. Поэтому `TestHello` — это экспортируемая функция теста.

### Шаг 2: Запустите тест

```bash
go test
```

Вы должны увидеть ошибку:

```
# hello [hello.test]
./hello_test.go:6:9: undefined: Hello
FAIL    hello [build failed]
```

Это **красный** этап TDD. Тест не компилируется, потому что функция `Hello` ещё не существует.

### Шаг 3: Напишите минимальный код для прохождения теста

Создайте файл `hello.go`:

```go
package main

import "fmt"

func Hello() string {
    return "Hello, World!"
}

func main() {
    fmt.Println(Hello())
}
```

### Шаг 4: Запустите тест снова

```bash
go test
```

Теперь вы должны увидеть:

```
PASS
ok      hello   0.002s
```

Это **зелёный** этап TDD. Тест проходит!

### Шаг 5: Рефакторинг

В данном случае код настолько прост, что рефакторинг не нужен. Но если бы мы написали более сложный код, сейчас было бы время его улучшить, сохраняя прохождение тестов.

Это **рефакторинг** этап TDD.

## Как работает `go test`

Команда `go test`:

1. Ищет все файлы, заканчивающиеся на `_test.go`
2. Находит все функции, начинающиеся с `Test`
3. Компилирует тестовый пакет
4. Запускает тесты
5. Сообщает результаты

## Полезные флаги

```bash
# Подробный вывод
go test -v

# Запуск конкретного теста
go test -run TestHello

# Покрытие кода
go test -cover

# Покрытие кода с детализацией
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Структура проекта

В этой главе у вас должна быть такая структура:

```
hello/
├── go.mod          # файл модуля Go
├── hello.go        # реализация
└── hello_test.go   # тесты
```

## Аналог в Java

### Тест на Go

```go
package main

import "testing"

func TestHello(t *testing.T) {
    got := Hello()
    want := "Hello, World!"

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
```

### Тест на Java (JUnit 5)

```java
import org.junit.jupiter.api.Test;
import static org.junit.jupiter.api.Assertions.assertEquals;

class HelloTest {

    @Test
    void hello() {
        String result = Hello.hello();
        String expected = "Hello, World!";
        assertEquals(expected, result);
    }
}
```

### Реализация на Go

```go
package main

import "fmt"

func Hello() string {
    return "Hello, World!"
}

func main() {
    fmt.Println(Hello())
}
```

### Реализация на Java

```java
public class Hello {
    public static String hello() {
        return "Hello, World!";
    }
}
```

## Сравнение Go и Java

| Характеристика | Go | Java |
|----------------|-----|------|
| **Пакет тестов** | Тот же пакет, что и код | Отдельная директория `src/test/java` |
| **Импорт тестирования** | `import "testing"` (встроенный) | `import org.junit.jupiter.api.Test` (внешняя библиотека) |
| **Обозначение теста** | `func TestXxx(t *testing.T)` | `@Test` аннотация |
| **Проверка (assertion)** | `if got != want { t.Errorf(...) }` | `assertEquals(expected, actual)` |
| **Запуск тестов** | `go test` | `mvn test` или `gradle test` |
| **Сообщение об ошибке** | `t.Errorf("got %q want %q", got, want)` | `assertEquals(expected, actual, message)` |
| **Строковые литералы** | `"Hello, World!"` | `"Hello, World!"` |

## Ключевые различия

### 1. Нет отдельного пакета для тестов

В Go тесты обычно находятся в том же пакете, что и тестируемый код. Это означает, что тесты имеют доступ к **неэкспортируемым** (приватным) функциям и переменным.

В Java тесты находятся в отдельном пакете и могут тестировать только `public` методы (если не использовать рефлексию).

### 2. Ручные проверки

В Go нет встроенных функций типа `assertEquals`. Вы пишете проверки вручную с помощью `if`. Это кажется лишним кодом, но даёт больше гибкости.

В JUnit есть богатый набор утверждений: `assertEquals`, `assertTrue`, `assertThrows` и т.д.

### 3. Встроенный пакет тестирования

В Go пакет `testing` — часть стандартной библиотеки. Ничего устанавливать не нужно.

В Java JUnit — внешняя библиотека, которую нужно добавить в зависимости.

### 4. Простота запуска

`go test` — это одна команда для всего. В Maven/Gradle вам нужно настроить плагины, зависимости и т.д.

## Практическое упражнение

Попробуйте изменить тест, чтобы он проверял другое сообщение:

```go
func TestHello(t *testing.T) {
    got := Hello()
    want := "Привет, мир!"

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
```

Запустите тест. Он должен упасть. Теперь измените функцию `Hello`, чтобы тест прошёл.

## Итоги главы

- ✅ Вы написали первый тест на Go
- ✅ Вы узнали цикл TDD: красный → зелёный → рефакторинг
- ✅ Вы поняли базовый синтаксис Go
- ✅ Вы увидели различия между Go и Java в тестировании

## Что дальше?

В следующей главе мы изучим **целые числа и циклы**, научимся работать с числами и напишем более сложные тесты.

---

**← [Настройка окружения](../../introduction/03-setup.md)** | **[Глава 2: Целые числа и циклы](../chapters/02-integers-and-loops/README.md) →**
