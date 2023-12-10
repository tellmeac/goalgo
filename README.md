# goalgo

## Как запустить (для демо)

Для простого запуска сервера без настройки (без бота)

```bash
docker compose up -d postgres app ml-functions
```

Для остановки

```bash
docker compose down
```

Для отображения вебсервив необходимо обратиться к адресу localhost:8000

Для настройки бота - создать бота в тг и узнать id канала для публикации уведомлений.
Нужно создать в корневой директории файл .env и указать токен и идентификатор чата (См .env.example)

```bash
# Запуск всех сервисов
docker compose up -d
```

## Стек

### Front-end

- react
- apexcharts

### Back-end

- Go
- REST
- postgres

### Machine learning

- python
- catboost
- psycopg2
- непараметрическая регрессия

### Dev-Ops

- Docker (compose)
