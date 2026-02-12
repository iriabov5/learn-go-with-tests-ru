# Настройка окружения

---

**← [Почему Go и TDD](02-why-go-and-tdd.md)** | **[Глава 1: Hello, World](../chapters/01-hello-world/README.md) →**

---

Прежде чем начать писать код, нужно настроить рабочее окружение. К счастью, для Go это довольно просто.

## Установка Go

### Проверьте, установлен ли Go

Откройте терминал и выполните:

```bash
go version
```

Если вы видите версию (например, `go version go1.21.5 darwin/amd64`), Go уже установлен.

### Установка

#### macOS

С помощью Homebrew:
```bash
brew install go
```

Или скачайте установщик с [golang.org/dl](https://golang.org/dl/)

#### Linux

```bash
wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz
```

Добавьте в `~/.bashrc` или `~/.zshrc`:
```bash
export PATH=$PATH:/usr/local/go/bin
```

#### Windows

Скачайте установщик с [golang.org/dl](https://golang.org/dl/) и запустите его.

### Проверка установки

```bash
go version
go env GOROOT
go env GOPATH
```

## Понимание рабочей области Go

Go имеет специфичную структуру рабочей области:

- **GOROOT:** Путь к установленной Go (обычно `/usr/local/go` или `C:\Go`)
- **GOPATH:** Путь к вашей рабочей области (по умолчанию `~/go`)
- **Go Modules:** Современный способ управления зависимостями (рекомендуется)

### Go Modules (рекомендуемый подход)

Начиная с Go 1.11, модули — это стандартный способ управления зависимостями. Вам **не нужно** настраивать `$GOPATH` для работы с модулями.

Создайте новый проект в любой директории:

```bash
mkdir myproject
cd myproject
go mod init myproject
```

Это создаст файл `go.mod`:

```go
module myproject

go 1.21
```

## Выбор редактора кода

### Visual Studio Code (рекомендуется)

VS Code — отличный выбор для Go с богатой экосистемой расширений.

1. Установите [VS Code](https://code.visualstudio.com/)
2. Установите расширение **Go** от Google

Установка расширения:
- Откройте VS Code
- Нажмите `Cmd/Ctrl + Shift + X`
- Найдите "Go"
- Нажмите "Install"

#### Полезные функции расширения Go:

- **Автодополнение кода**
- **Навигация по коду** (Go to Definition)
- **Быстрое исправление ошибок**
- **Запуск тестов** из редактора
- **Отладка**
- **Форматирование** (自动 на save)

### Альтернативы

#### GoLand (JetBrains)

- Платный IDE с мощными возможностями
- Отличный для больших проектов
- Интегрированный отладчик и профайлер

#### Vim/Neovim

С плагином `vim-go`:
```vim
:PluginInstall 'fatih/vim-go'
```

#### Emacs

С пакетом `go-mode`:
```elisp
(use-package go-mode)
```

## Полезные инструменты

### gofmt — форматирование кода

Go имеет официальный форматировщик кода:

```bash
gofmt -w yourfile.go
```

В VS Code форматирование происходит автоматически при сохранении.

### go vet — статический анализ

```bash
go vet ./...
```

Находит распространённые ошибки, которые компилятор пропускает.

### golint — проверка стиля

```bash
go install golang.org/x/lint/golint@latest
golint ./...
```

Проверяет соответствие стилю Go.

### golangci-lint — набор линтеров

```bash
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
golangci-lint run
```

Мощный инструмент, объединяющий множество линтеров.

## Первый тест

Давайте напишем первый тест, чтобы убедиться, что всё настроено правильно.

### Создайте директорию проекта

```bash
mkdir hello
cd hello
go mod init hello
```

### Создайте файл теста `hello_test.go`

```go
package hello

import "testing"

func TestHello(t *testing.T) {
    got := Hello()
    want := "Hello, World!"

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
```

### Запустите тест

```bash
go test
```

Вы должны увидеть ошибку (потому что функция `Hello` ещё не реализована):

```
# hello [hello.test]
./hello_test.go:6:9: undefined: Hello
FAIL    hello [build failed]
```

### Реализуйте функцию в `hello.go`

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

Теперь файл `hello.go` содержит пакет `main` и функцию `main()`, что позволяет запустить программу с помощью команды `go run hello.go`.

### Запустите тест снова

```bash
go test
```

Теперь вы должны увидеть:

```
PASS
ok      hello   0.002s
```

**Поздравляем!** Вы написали первый тест и первую функцию на Go.

## Настройка для Java-разработчиков

Если вы пришли из Java, вот несколько советов:

| Java | Go |
|------|-----|
| `src/main/java` | Любая директория (обычно корень проекта) |
| `src/test/java` | Файлы `*_test.go` рядом с исходным кодом |
| Maven/Gradle | `go mod` для зависимостей |
| `mvn test` / `gradle test` | `go test` |
| JUnit | Встроенный пакет `testing` |
| `System.out.println` | `fmt.Println` или `log.Printf` |
| IDE автодополнение | `gopls` (Language Server Protocol) |

### Полезные команды для сравнения

| Задача | Maven/Gradle | Go |
|--------|--------------|-----|
| Сборка | `mvn compile` / `gradle build` | `go build` |
| Запуск тестов | `mvn test` / `gradle test` | `go test` |
| Запуск приложения | `mvn exec:java` / `gradle run` | `go run main.go` |
| Установка зависимостей | `mvn install` / `gradle build` | `go mod tidy` |
| Форматирование | `mvn spotless:apply` / `gradle spotlessApply` | `gofmt -w .` |
| Статический анализ | `mvn checkstyle:check` / `gradle checkstyleMain` | `go vet ./...` |

## Проверьте всё

Убедитесь, что всё работает:

```bash
# Версия Go
go version

# Создайте тестовый проект
mkdir test-setup && cd test-setup
go mod init test-setup

# Создайте простой файл
cat > main.go << 'EOF'
package main

import "fmt"

func main() {
    fmt.Println("Go is ready!")
}
EOF

# Запустите
go run main.go

# Создайте тест
cat > main_test.go << 'EOF'
package main

import "testing"

func TestReady(t *testing.T) {
    if true != true {
        t.Error("Something is wrong")
    }
}
EOF

# Запустите тест
go test
```

Если всё работает, вы готовы к первой главе!

## Что дальше?

Теперь, когда окружение настроено, переходим к первой главе: **Hello, World**, где мы напишем первую программу и первый тест по-настоящему.

---

**← [Почему Go и TDD](02-why-go-and-tdd.md)** | **[Глава 1: Hello, World](../chapters/01-hello-world/README.md) →**
