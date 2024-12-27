# Calculator API

API для решения математических выражений.

## Запуск
### Docker
Можно запустить приложение через [Docker](https://www.docker.com/).

- Docker run
```shell
# образ с GitHub
docker pull ghcr.io/linuxfight/yandexcalcapi:main
docker run -p 8080:8080 -d ghcr.io/linuxfight/yandexcalcapi:main

# сборка локально
docker build -t test .
docker run -p 8080:8080 -d test
```
- Docker compose
```shell
docker compose -f compose.yml up --build -d
``` 

### Обычная сборка
```shell
# unix
go mod download -x
go build -o ./server ./cmd/main.go

./server

# windows
go mod download -x
go build -o ./server.exe ./cmd/main.go

./server.exe
```

## Использование
Для использования вам нужен http клиент (к примеру [Postman](https://www.postman.com/downloads/) или [Insomnia](https://insomnia.rest/download)).
Также можно использовать любой http клиент для вашего языка программирования.

Можно использовать уже запущенную версию - https://solve.linuxfight.me/api/v1/calculate.

Отправьте запрос на ```http://localhost:8080/api/v1/calculate```

Пример на typescript
```typescript
const sendPostRequest = async () => {
  const url = 'http://localhost:8080/api/v1/calculate';
  const data = {
    expression: '2+2*2',
  };

  try {
    const response = await fetch(url, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    });

    if (!response.ok) {
      throw new Error('Network response was not ok');
    }

    const responseData = await response.json();
    console.log('Response:', responseData);
  } catch (error) {
    console.error('Error:', error);
  }
};

sendPostRequest();
```

Такой же пример с помощью curl
```shell
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2"
}'
```

## Структура проекта
- ```.github``` - CI для проверки кода, автоматической сборки Docker, обновления зависимостей

- ```cmd``` - пакет main, код для запуска

- ```internal``` - код веб сервера
  - ```logger``` - логирование с помощью [Zap](https://github.com/uber-go/zap)
  - ```handlers``` - функции для обработки http запросов с помощью [Fiber](https://github.com/gofiber/fiber)

- ```pkg/calc``` - код для обработки выражений

## Документация API

```expression``` - строка-выражение состоящее из односимвольных идентификаторов и знаков арифметических действий.

Входящие данные - цифры(рациональные), операции +, -, *, /, операции приоритизация ( или )

В случае ошибки записи выражения сервер выдает ошибку 422.
При успешном выполнении будет получен код 200.
Если будет вызвана ошибка на сервере, то будет возвращён код 500. 

### Примеры:
#### 200 - успешное выполнение

Запрос: 
```json
{
  "expression": "2+2*2"
}
```
Ответ:
```json
{
    "result": "6"
}
```

#### 422 - входные данные не соответствуют требованиям приложения
Если запрос не содержит заголовок ```Content-Type``` = ```application/json```, то вы получите:
```json
{
  "error": "Content-Type must be application/json"
}
```

Запрос:
```json
{
  "testing": 123
}
```
или
```json
{
  "expression": ""
}
```
Ответ:
```json
{
    "error": "invalid JSON"
}
```

Также может быть вызвано неверным выражением:
```json
{
  "expression": "a + 523"
}
```
или
```json
{
  "expression": "+521*21()())))))"
}
```
Ответ:
```json
{
  "error": "invalid character/expression/etc"
}
```

#### 500 - ошибка веб сервера
Ошибка может произойти, если выполнение Solve вызвало панику.

```json
{
    "error": "Internal server error"
}
```