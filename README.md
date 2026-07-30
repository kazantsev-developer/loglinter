# loglinter

`loglinter` is a high-performance static analysis tool for Go designed to enforce unified logging standards across highload microservices. It operates with zero-allocations on critical string validation paths and supports standard `log/slog` and `go.uber.org/zap` loggers.

## Features

- **Casing Enforcement**: Validates that all log messages start with a lowercase letter.
- **Localization Control**: Restricts log messages to the English language only (ASCII/Latin safety zone).
- **Format Validation**: Disallows emojis, math symbols, and forbidden punctuation (`!`, `?`) in log messages.
- **Security & Compliance**: Prevents potential sensitive data leaks (e.g., passwords, API keys, tokens, secrets) via low-allocation pattern matching.

## Architecture & Layout

The repository strictly follows the Standard Go Project Layout:

- `cmd/loglinter/`: Standalone CLI tool entry point using `singlechecker`.
- `plugin/`: Entry point for modern `golangci-lint` Module-based plugin integration.
- `pkg/linter/`: Core analysis logic, abstract syntax tree (AST) traversal, and linting rules.

## Installation & Build

Build automation is managed entirely via the provided `Makefile`. To run tests and compile all artifacts, execute:

```bash
make
```

This single command triggers:

1. Integrated test suites via `analysistest` with active race detection (`-race`).
2. Standalone CLI building with linker optimization flags (`-ldflags="-s -w"`).
3. Dynamic shared library compilation (`.so`) for plugin systems.

Individual targets:

- `make test`: Run code verification suites.
- `make build`: Compile binary targets into the `bin/` directory.
- `make clean`: Purge pre-compiled artifacts.

## Usage Examples

### Standalone CLI Execution

You can run the `loglinter` binary directly over your package tree:

```bash
./bin/loglinter ./...
```

### Modern golangci-lint Integration

To enable this linter as a custom plugin inside your `.golangci.yml` configuration:

```yaml
linters-settings:
  custom:
    loglinter:
      path: ./bin/loglinter.so
      description: Checks log messages for style, language, and security guidelines

linters:
  enable:
    - loglinter
```

---

## Правила валидации логов

Линтер автоматически сканирует абстрактное синтаксическое дерево (AST) и пресекает следующие типы некорректных лог-записей для логгеров `slog` и `zap`:

### 1. Лог-сообщения должны начинаться со строчной буквы

- **Неправильно**: `log.Info("Starting server on port 8080")`
- **Правильно**: `log.Info("starting server on port 8080")`

### 2. Лог-сообщения должны быть только на английском языке

- **Неправильно**: `log.Error("ошибка подключения к базе данных")`
- **Правильно**: `log.Error("failed to connect to database")`

### 3. Лог-сообщения не должны содержать спецсимволы или эмодзи

- **Неправильно**: `log.Info("server started!🚀")`
- **Правильно**: `log.Info("server started")`

### 4. Лог-сообщения не должны содержать потенциально чувствительные данные

- **Неправильно**: `log.Info("user password: " + password)` или `log.Debug("api_key=" + apiKey)`
- **Правильно**: `log.Info("user authenticated successfully")`
