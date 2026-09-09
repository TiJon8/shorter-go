# 🔪 Background Link Shortener

Простой сервер работающий в фоне по сокращению ссылок. Приложение имеет базу данных, так что если нужно хранить много ссылок, это удобное решение, плюсом можно установить свое осмысленное название для исходной ссылки

## Usage

Можно использовать и как библиотеку с инициализацией из кода, и как исполняемый файл

#### Скачать через `go get`

```bash
go get -u github.com/TiJon8/shorter-go
```

Базовый пример использования с инициализацией из кода

```go
import (
	"context"
	"os"
	"os/signal"
	"syscall"

	c "github.com/TiJon8/shorter-go/pkg/config"
	f "github.com/TiJon8/shorter-go/pkg/features"

	l "github.com/TiJon8/shorter-go/pkg/logger"
	srv "github.com/TiJon8/shorter-go/pkg/server"
	s "github.com/TiJon8/shorter-go/pkg/storage"
)

func main() {
	appCfg := c.AppConfigMust()
	logger := l.NewLogger(appCfg.Env, appCfg.LogLevel)
	storage, err := s.Init(appCfg.StoragePath)
	if err != nil {
		logger.Error("Open connection to databse failed", l.ErrorAttr("error", err))
		os.Exit(1)
	}
	router := f.InitRouter(storage, nil)
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	serverCfg := c.ServerConfigMust()
	Server := srv.NewHTTPServer(serverCfg, nil, router, storage)
	if err := Server.Run(ctx); err != nil {
		logger.Error("Server error", l.ErrorAttr("err", err))
	}
}
```

**Важно** AppConfigMust ищет определенные переменные окружения в частности требует `storage_path` - путь до файла sqlite

При использовании AppConfigMust и ServerConfigMust нужно определить переменные окружения. В основном только `HTTP_ADDR` и `STORAGE_PATH` являются обязательными

```env
ENV=prod                                # prod|dev (для настройки вывода логгера)
HTTP_ADDR=:8089                         # порт для запуска сервера в формате :uint16
HTTP_GRACEFULL_SHUTDOWNN_DURATION=9s    # время для gracefull остановки
STORAGE_PATH=./database.db              # путь по которому будет файл sqlite
LOG_LEVEL=warn                          # debug|info|warn|error (минимальный уровень логирования)
```

Чтобы не экспортировать всех по отдельности рекомендуется использовать Makefile

```Makefile
include .env
export

run:
	@go run main.go &

```

Запустив `make run` сервер запустится на определенном порту и будет работать в фоне и не привязан к сессии терминала

#### Скачать через `go install`

```bash
go install github.com/TiJon8/shorter-go/cmd/shorter@latest
```

Так как исполняемый файл также треубет переменные окружения, при запуске можно определить `--use-env`, нужно ли искать .env файл, если `--use-env` не задана то необходимо указать как минимум флаг `--addr=:port`. Для справки и возможные опции запустите `shorter --help`

Пример команды

```bash
shorter --addr :4090 --storage-path ./store.db --env prod --log-level warn &
```

При `--use-env` приложение будет смотреть на переменные окржуения из .env, переданные значения в флагах учитываться не будут

Если сервер запущен, можно перейти на localhost:port/ping и в ответе получить `Hello Chi`

Чтобы остановить фоновый сервер отправьте сигнал -2 или -15 для процесса, id процесса будет записанно в файл `pid` при старте сервера

```bash
kill -2 pid
```

## API

Использовать API сервера можно из любого клиента который позволяет использовать http протокол без политики CORS

#### Сократить ссылку

```http
  POST /short
```

В body указывается

| Parameter | Type     | Description                                                                                   |
| :-------- | :------- | :-------------------------------------------------------------------------------------------- |
| `url`     | `string` | **Required**. исходая ссылка                                                                  |
| `alias`   | `string` | **Unique**. алиас по которому будет происходить redirect, по умолчанию генерируется случайный |

#### Изменить ссылку (для существуещего алиаса)

```http
  PATCH /short
```

В body указывается

| Parameter | Type     | Description                                      |
| :-------- | :------- | :----------------------------------------------- |
| `new_url` | `string` | **Required**. новая ссылка                       |
| `alias`   | `string` | **Required**. алиас для которого имзенить ссылку |

#### Удалить ссылку

```http
  DELETE /short
```

В body указывается

| Parameter | Type     | Description                                      |
| :-------- | :------- | :----------------------------------------------- |
| `alias`   | `string` | **Required**. алиас по которому удаляется запись |

## Зависимости

Данный проект построен на самых френдли бибилотеках: go-chi и go-sqlite3, которые в свою очередь совместимы со встроенными в go интерфейсами, как net/http и database/sql. Так что обновить и добавить в исходный код свое очень просто

Here are some related projects

[go-chi документация и примеры](https://go-chi.io/#/README)

[go-sqlite3](https://pkg.go.dev/gopkg.in/mattn/go-sqlite3.v2)

[go database/sql интерфейс](https://pkg.go.dev/database/sql)

## Feedback

Если вы наткнулись на данный репозиторий и у вас есть вдохновение и мотивация его улучшить, прошу, эксперементируйте :)

Буду рад предложениям и идеям в [telegram](https://t.me/stesnyashka)
