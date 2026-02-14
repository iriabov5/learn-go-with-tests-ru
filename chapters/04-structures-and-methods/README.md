# Глава 4: Структуры и методы

---

**← [Глава 3: Итерации](../03-iteration/README.md)** | **[Глава 5: Интерфейсы](../05-interfaces/README.md) →**

---

В этой главе мы изучим структуры и методы в Go. Мы будем использовать методологию TDD (Test-Driven Development) для написания кода.

## Что вы узнаете

- Как создавать структуры в Go
- Как определять методы для структур
- Как использовать TDD при работе со структурами и методами
- Сравнение с Java: как те же концепции реализуются в Java
- Практические примеры и упражнения

## Структуры в Go

Структуры позволяют группировать связанные данные вместе. Давайте начнем с создания простой структуры `Rectangle` (прямоугольник).

### Пример: Прямоугольник с использованием TDD

Давайте напишем тест для метода, который вычисляет площадь прямоугольника.

#### Шаг 1: Напишите тест

Сначала создадим тест для метода вычисления площади:

```go
package geometry

import "testing"

func TestRectangle(t *testing.T) {
	rectangle := Rectangle{Width: 10, Height: 5}
	area := rectangle.Area()
	expectedArea := 50.0

	if area != expectedArea {
		t.Errorf("Expected area %f, but got %f", expectedArea, area)
	}
}
```

#### Шаг 2: Запустите тест и убедитесь, что он падает

Когда мы запустим тест, мы получим ошибки компиляции, потому что структура `Rectangle` и метод `Area` еще не существуют.

#### Шаг 3: Напишите минимальный код для запуска теста

Создадим структуру `Rectangle` и метод `Area`, которые пока возвращают заглушку:

```go
package geometry

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return 0
}
```

Теперь тест запустится, но упадет с сообщением о неверной площади.

#### Шаг 4: Напишите достаточно кода, чтобы пройти тест

Теперь реализуем правильную логику вычисления площади:

```go
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
```

Теперь тест должен пройти успешно.

#### Шаг 5: Рефакторинг

Мы можем добавить метод для вычисления периметра:

```go
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}
```

## Методы в Go

Методы в Go - это функции с получателем. Получатель указывается между ключевым словом `func` и именем метода.

Синтаксис:
```go
func (receiverName ReceiverType) MethodName(args) (returnType) {
    // тело метода
}
```

### Go

```go
// shape.go
package geometry

// Rectangle представляет собой прямоугольник с шириной и высотой
type Rectangle struct {
	Width  float64
	Height float64
}

// Area вычисляет площадь прямоугольника
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter вычисляет периметр прямоугольника
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}
```

```go
// shape_test.go
package geometry

import "testing"

func TestRectangle(t *testing.T) {
	rectangle := Rectangle{Width: 10, Height: 5}
	area := rectangle.Area()
	expectedArea := 50.0

	if area != expectedArea {
		t.Errorf("Expected area %f, but got %f", expectedArea, area)
	}

	perimeter := rectangle.Perimeter()
	expectedPerimeter := 30.0

	if perimeter != expectedPerimeter {
		t.Errorf("Expected perimeter %f, but got %f", expectedPerimeter, perimeter)
	}
}
```

### Java

```java
public class Rectangle {
    private double width;
    private double height;

    public Rectangle(double width, double height) {
        this.width = width;
        this.height = height;
    }

    public double getArea() {
        return width * height;
    }

    public double getPerimeter() {
        return 2 * (width + height);
    }

    public static void main(String[] args) {
        Rectangle rectangle = new Rectangle(10, 5);
        System.out.println("Area: " + rectangle.getArea()); // Вывод: Area: 50.0
        System.out.println("Perimeter: " + rectangle.getPerimeter()); // Вывод: Perimeter: 30.0
    }
}
```

```java
// Тест на JUnit
import org.junit.Test;
import static org.junit.Assert.assertEquals;

public class RectangleTest {
    @Test
    public void testArea() {
        Rectangle rectangle = new Rectangle(10, 5);
        assertEquals(50.0, rectangle.getArea(), 0.001);
    }

    @Test
    public void testPerimeter() {
        Rectangle rectangle = new Rectangle(10, 5);
        assertEquals(30.0, rectangle.getPerimeter(), 0.001);
    }
}
```

## Добавление других фигур

Теперь давайте добавим другие фигуры: круг и треугольник.

```go
// shape.go
package geometry

import "math"

// Circle represents a circle with radius
type Circle struct {
	Radius float64
}

// Area calculates the area of a circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter calculates the perimeter of a circle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Triangle represents a triangle with three sides
type Triangle struct {
	Base   float64
	Height float64
}

// Area calculates the area of a triangle
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

// Perimeter calculates the perimeter of a triangle
func (t Triangle) Perimeter() float64 {
	// For simplicity, we'll just return the base + height + hypotenuse
	// assuming it's a right triangle
	hypotenuse := math.Sqrt(t.Base*t.Base + t.Height*t.Height)
	return t.Base + t.Height + hypotenuse
}
```

## Таблицы тестов

Мы можем использовать таблицы тестов для эффективного тестирования нескольких случаев:

```go
func TestAreaWithInterface(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			area := tt.shape.Area()
			if area != tt.hasArea {
				t.Errorf("%#v: Expected area %f, but got %f", tt.shape, tt.hasArea, area)
			}
		})
	}
}
```

## Сравнение Go и Java

| Характеристика | Go | Java |
|----------------|-----|------|
| **Структуры** | `struct` | `class` |
| **Методы** | Определяются с получателем | Определяются внутри класса |
| **Получатель метода** | `(r Rectangle) MethodName()` | `this` внутри метода |
| **Интерфейсы** | Неявная реализация | Явная реализация с `implements` |
| **Конструкторы** | Функции вроде `NewRectangle()` | Конструкторы с `new` |
| **Именование** | `camelCase` для экспортируемых, `snake_case` не используется | `camelCase` для переменных и методов |
| **Экспорт** | Заглавная буква делает элемент экспортируемым | Модификаторы доступа: `public`, `private`, `protected` |

### Именование в Go и Java

В Go и Java существуют различные подходы к именованию элементов кода:

#### Go:
- Экспортируемые (публичные) функции, типы, переменные и константы начинаются с заглавной буквы: `Area`, `Rectangle`, `GetWidth`
- Неэкспортируемые (приватные) элементы начинаются со строчной буквы: `area`, `rectangle`, `getWidth`
- Принято использовать `camelCase` для именования функций и типов
- `snake_case` не используется в Go

#### Java:
- Имена классов используют `UpperCamelCase`: `Rectangle`, `Circle`
- Имена методов и переменных используют `lowerCamelCase`: `getArea()`, `calculatePerimeter()`
- Константы используют `UPPER_SNAKE_CASE`: `MAX_VALUE`, `DEFAULT_SIZE`
- Модификаторы доступа (`public`, `private`, `protected`) определяют, какие элементы доступны из других пакетов

## Практическое упражнение

Попробуйте создать структуру `Square` (квадрат) и методы для вычисления площади и периметра, следуя TDD подходу:
1. Сначала напишите тест
2. Убедитесь, что он падает
3. Напишите минимальный код для запуска
4. Реализуйте правильную логику
5. Добавьте документацию

### Решение

```go
// Square представляет собой квадрат
type Square struct {
	Side float64
}

// Area вычисляет площадь квадрата
func (s Square) Area() float64 {
	return s.Side * s.Side
}

// Perimeter вычисляет периметр квадрата
func (s Square) Perimeter() float64 {
	return 4 * s.Side
}
```

```go
// Тест для квадрата
func TestSquare(t *testing.T) {
	square := Square{Side: 5}
	area := square.Area()
	expectedArea := 25.0

	if area != expectedArea {
		t.Errorf("Expected area %f, but got %f", expectedArea, area)
	}

	perimeter := square.Perimeter()
	expectedPerimeter := 20.0

	if perimeter != expectedPerimeter {
		t.Errorf("Expected perimeter %f, but got %f", expectedPerimeter, perimeter)
	}
}
```

## Итоги главы

- ✅ Вы узнали, как создавать структуры в Go
- ✅ Вы изучили, как определять методы для структур
- ✅ Вы изучили TDD подход при работе со структурами и методами
- ✅ Вы сравнили концепции с Java
- ✅ Вы решили практическое упражнение

## Что дальше?

В следующей главе мы изучим **интерфейсы**, которые позволяют создавать абстракции и упрощать код.

---

**← [Глава 3: Итерации](../03-iteration/README.md)** | **[Глава 5: Интерфейсы](../05-interfaces/README.md) →**