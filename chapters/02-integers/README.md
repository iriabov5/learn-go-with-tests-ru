# Глава 2: Целые числа

---

**← [Глава 1: Hello, World](../01-hello-world/README.md)** | **[Глава 3: Итерации](../03-iteration/README.md) →**

---

В этой главе мы изучим работу с целыми числами в Go. Мы будем использовать методологию TDD (Test-Driven Development) для написания кода.

## Что вы узнаете

- Как работать с целыми числами в Go
- Как использовать TDD при работе с целыми числами
- Сравнение с Java: как те же концепции реализуются в Java
- Практические примеры и упражнения

## Целые числа в Go

В Go есть несколько типов для представления целых чисел:

- `int`: целое число со знаком, размер зависит от платформы (обычно 64 бита)
- `int8`, `int16`, `int32`, `int64`: целые числа со знаком с фиксированным размером
- `uint`: целое число без знака, размер зависит от платформы
- `uint8`, `uint16`, `uint32`, `uint64`: целые числа без знака с фиксированным размером

### Пример: Сложение двух чисел с использованием TDD

Давайте напишем функцию, которая складывает два целых числа, используя TDD подход.

#### Шаг 1: Напишите тест

Сначала создадим тест для функции сложения:

```go
package integers

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}
```

#### Шаг 2: Запустите тест и убедитесь, что он падает

Когда мы запустим тест, мы получим ошибку компиляции, потому что функция `Add` еще не существует:

```
./adder_test.go:6:9: undefined: Add
```

#### Шаг 3: Напишите минимальный код для запуска теста

Создадим функцию `Add`, которая пока возвращает 0:

```go
package integers

func Add(x, y int) int {
	return 0
}
```

Теперь тест запустится, но упадет с сообщением:

```
adder_test.go:10: expected '4' but got '0'
```

#### Шаг 4: Напишите достаточно кода, чтобы пройти тест

Теперь реализуем правильную логику:

```go
package integers

func Add(x, y int) int {
	return x + y
}
```

Теперь тест должен пройти успешно.

#### Шаг 5: Рефакторинг

Добавим документацию к функции:

```go
// Add takes two integers and returns the sum of them.
func Add(x, y int) int {
	return x + y
}
```

#### Тестируемые примеры

Можно также добавить тестируемый пример:

```go
func ExampleAdd() {
	sum := Add(1, 5)
	print(sum)
	// Output: 6
}
```

#### Go

```go
// adder.go
package integers

// Add takes two integers and returns the sum of them.
func Add(x, y int) int {
	return x + y
}
```

```go
// adder_test.go
package integers

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	print(sum)
	// Output: 6
}
```

#### Java

```java
public class IntegerUtils {
    /**
     * Складывает два целых числа
     */
    public static int add(int a, int b) {
        return a + b;
    }

    public static void main(String[] args) {
        int result = add(2, 2);
        System.out.println(result); // Вывод: 4
    }
}
```

```java
// Тест на JUnit
import org.junit.Test;
import static org.junit.Assert.assertEquals;

public class IntegerUtilsTest {
    @Test
    public void testAdd() {
        int result = IntegerUtils.add(2, 2);
        assertEquals("2 + 2 should equal 4", 4, result);
    }
}
```

## Сравнение Go и Java

| Характеристика | Go | Java |
|----------------|-----|------|
| **Целые числа** | `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64` | `byte`, `short`, `int`, `long` |
| **Объявление переменной** | `var x int` или `x := 5` | `int x = 5;` |
| **Типизация** | Статическая | Статическая |
| **Подход к тестированию** | Встроенный пакет `testing` | JUnit, TestNG |
| **Синтаксис тестов** | Функции с префиксом `Test` | Аннотации `@Test` |

## Практическое упражнение

Попробуйте написать функцию `Multiply(x, y int) int`, которая перемножает два целых числа, следуя TDD подходу:
1. Сначала напишите тест
2. Убедитесь, что он падает
3. Напишите минимальный код для запуска
4. Реализуйте правильную логику
5. Добавьте документацию

### Решение

```go
// multiply.go
package integers

// Multiply takes two integers and returns the product of them.
func Multiply(x, y int) int {
	return x * y
}
```

```go
// multiply_test.go
package integers

import "testing"

func TestMultiply(t *testing.T) {
	result := Multiply(3, 4)
	expected := 12

	if result != expected {
		t.Errorf("expected '%d' but got '%d'", expected, result)
	}
}
```

## Итоги главы

- ✅ Вы узнали, как работать с целыми числами в Go
- ✅ Вы изучили TDD подход при работе с целыми числами
- ✅ Вы сравнили концепции с Java
- ✅ Вы решили практическое упражнение

## Что дальше?

В следующей главе мы изучим **итерации**, где узнаем, как использовать циклы в Go.

---

**← [Глава 1: Hello, World](../01-hello-world/README.md)** | **[Глава 3: Итерации](../03-iteration/README.md) →**