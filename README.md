# Junior_Backend_To-Do

Этот проект создан для демонстрации навыков разработки бэкенда и включает базовый функционал для создания, обновления, удаления и получения задач.

## Функции

- Создание новых задач
- Получение списка задач
- Обновление существующих задач
- Удаление задач

## Технологии

Проект использует следующие технологии и инструменты:

- **Language**: Golang
- **Framework**: gorilla, gorm
- **Database**: PostgreSQL

## Установка

1. Клонируйте репозиторий:

    ```bash
    git clone https://github.com/5afar/Junior_Backend_To-Do.git
    cd Junior_Backend_To-Do
    ```

2. Установите необходимые зависимости:

    ```bash
    go mod tidy
    ```

4. Запустите приложение:

    ```bash
    go run cmd/main.go
    ```

5. Приложение будет доступно по адресу: `http://localhost:8080`

## Использование

После запуска приложения, вы можете использовать следующие эндпоинты:

- `GET /tasks` — Получить список всех задач
- `POST /tasks` — Создать новую задачу (передайте данные задачи в теле запроса)
- `PUT /tasks/<id>` — Обновить задачу по ID (передайте обновленные данные в теле запроса)
- `DELETE /tasks/<id>` — Удалить задачу по ID
