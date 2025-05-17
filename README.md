# TodoList gRPC/HTTP Microservice 

Очень простой микросервис для управления задачами с двойным интерфейсом: gRPC (высокая производительность) и HTTP/REST (совместимость). Полностью контейнеризированное решение. Масштабируем.

## Особенности

- **Двойной API**: 
  - gRPC на порту `50051` 
  - REST через gRPC-Gateway на `8080`
- **Автогенерация** кода из `.proto`-файлов
- **Готовность к продакшену**:
  - Docker-образы Alpine

## 🚀 Быстрый старт

### Требования
- Go 1.24.3+
- Docker 28.1.1+
- protoc 31+

### Запуск
```bash
# 1. Клонировать репозиторий
git clone https://github.com/yourusername/grpc-todo.git
cd grpc-todo

# 2. Запустить в Docker
docker-compose up --build

# 3. Проверить работу
curl -X POST http://localhost:8080/v1/tasks \
  -H "Content-Type: application/json" \
  -d '{"title": "Купить молоко"}'
```
📚 API Endpoints
```
REST (HTTP/JSON)
Метод	    |   Путь	           |   Описание
POST	    |   /v1/tasks	       |   Добавить задачу
GET	      |   /v1/tasks	       |   Список задач
PATCH	    |   /v1/tasks/{id}   |   Обновить задачу
DELETE	  |   /v1/tasks/{id}   |	 Удалить задачу
```
📦 Структура проекта
```
.
├── proto/               # Protobuf-файлы
├── server/              # gRPC-сервер
├── gateway/             # HTTP-шлюз
├── client/              # Пример клиента (опционально)
├── docker-compose.yml
├── Makefile
└── README.md
```
