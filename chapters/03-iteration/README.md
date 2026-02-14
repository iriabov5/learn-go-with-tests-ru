# Глава 3: Итерации

---

**← [Глава 2: Целые числа](../02-integers/README.md)** | **[Глава 4: Структуры и методы](../04-structures-and-methods/README.md) →**

---

В этой главе мы изучим итерации (циклы) в Go. Мы будем использовать методологию TDD (Test-Driven Development) для написания кода.

## Что вы узнаете

- Как использовать циклы в Go
- Как применять TDD при работе с циклами
- Сравнение с Java: как циклы реализуются в Java
- Практические примеры и упражнения

## Циклы в Go

В Go есть только один тип цикла: `for`. Он может использоваться для различных видов итераций.

### Пример: Повторение строки с использованием TDD

Давайте напишем функцию, которая повторяет строку заданное количество раз, используя TDD подход.

#### Шаг 1: Напишите тест

Сначала создадим тест для функции повторения:

```go
package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
```

#### Шаг 2: Запустите тест и убедитесь, что он падает

Когда мы запустим тест, мы получим ошибку компиляции, потому что функция `Repeat` еще не существует:

```
./repeat_test.go:6:12: undefined: Repeat
```

#### Шаг 3: Напишите минимальный код для запуска теста

Создадим функцию `Repeat`, которая пока возвращает пустую строку:

```go
package iteration

func Repeat(s string, count int) string {
	return ""
}
```

Теперь тест запустится, но упадет с сообщением:

```
repeat_test.go:10: expected "aaaaa" but got ""
```

#### Шаг 4: Напишите достаточно кода, чтобы пройти тест

Теперь реализуем простую логику с конкатенацией строк:

```go
package iteration

func Repeat(s string, count int) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated += s
	}
	return repeated
}
```

Теперь тест должен пройти успешно.

#### Шаг 5: Рефакторинг

Хотя наша реализация работает, она неэффективна из-за многократного создания новых строк. В Go строки неизменяемы, поэтому каждый оператор `+=` создает новую строку.

Давайте улучшим производительность, используя `strings.Builder`:

```go
package iteration

import "strings"

// Repeat takes a string and repeats it count times.
func Repeat(s string, count int) string {
	var builder strings.Builder
	for i := 0; i < count; i++ {
		builder.WriteString(s)
	}
	return builder.String()
}
```

#### Go

```go
// repeat.go
package iteration

import "strings"

// Repeat takes a string and repeats it count times.
func Repeat(s string, count int) string {
	var builder strings.Builder
	for i := 0; i < count; i++ {
		builder.WriteString(s)
	}
	return builder.String()
}
```

```go
// repeat_test.go
package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
```

#### Java

```java
public class StringRepeater {
    /**
     * Повторяет строку заданное количество раз
     */
    public static String repeat(String s, int count) {
        StringBuilder sb = new StringBuilder();
        for (int i = 0; i < count; i++) {
            sb.append(s);
        }
        return sb.toString();
    }

    public static void main(String[] args) {
        String result = repeat("a", 5);
        System.out.println(result); // Вывод: aaaaa
    }
}
```

```java
// Тест на JUnit
import org.junit.Test;
import static org.junit.Assert.assertEquals;

public class StringRepeaterTest {
    @Test
    public void testRepeat() {
        String result = StringRepeater.repeat("a", 5);
        assertEquals("aaaaa", result);
    }
}
```

## Типы циклов for в Go

В Go цикл `for` может использоваться несколькими способами:

### 1. Классический for цикл

```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

### 2. While цикл

```go
i := 0
for i < 10 {
    fmt.Println(i)
    i++
}
```

### 3. Бесконечный цикл

```go
for {
    // делаем что-то бесконечно
    if условие {
        break
    }
}
```

## Бенчмарки

Мы можем использовать бенчмарки для измерения производительности нашей функции:

```go
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
```

Для запуска бенчмарков используйте команду:

```
go test -bench=.
```

## Сравнение Go и Java

| Характеристика | Go | Java |
|----------------|-----|------|
| **Циклы** | Только `for` | `for`, `while`, `do-while` |
| **Инициализация переменной в цикле** | `for i := 1; i <= 5; i++` | `for (int i = 1; i <= 5; i++)` |
| **Бесконечный цикл** | `for { ... }` | `while (true) { ... }` |
| **Выход из цикла** | `break` | `break` |
| **Пропуск итерации** | `continue` | `continue` |

## Практическое упражнение

Попробуйте написать функцию `CountDown(n int) string`, которая создает строку обратного отсчета от n до 0, разделенную запятыми, следуя TDD подходу:
1. Сначала напишите тест
2. Убедитесь, что он падает
3. Напишите минимальный код для запуска
4. Реализуйте правильную логику
5. Добавьте документацию

### Решение

```go
// countdown.go
package iteration

import (
	"strconv"
	"strings"
)

// CountDown takes an integer and returns a string counting down from that number to 0.
func CountDown(n int) string {
	var parts []string
	for i := n; i >= 0; i-- {
		parts = append(parts, strconv.Itoa(i))
	}
	return strings.Join(parts, ",")
}
```

```go
// countdown_test.go
package iteration

import "testing"

func TestCountDown(t *testing.T) {
	result := CountDown(3)
	expected := "3,2,1,0"

	if result != expected {
		t.Errorf("expected %q but got %q", expected, result)
	}
}
```

## Итоги главы

- ✅ Вы узнали, как использовать циклы в Go
- ✅ Вы изучили TDD подход при работе с циклами
- ✅ Вы сравнили концепции с Java
- ✅ Вы решили практическое упражнение

## Что дальше?

В следующей главе мы изучим **структуры и методы**, где узнаем, как создавать пользовательские типы данных и методы для них.

---

**← [Глава 2: Целые числа](../02-integers/README.md)** | **[Глава 4: Структуры и методы](../04-structures-and-methods/README.md) →**