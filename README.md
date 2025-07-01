# 💸 ExpenseTracker (Backend API на Golang)

## 📋 Описание

**ExpenseTracker** — это веб-приложение для учета личных финансов. Пользователь может регистрироваться, добавлять доходы, расходы, управлять целями и лимитами, получать отчеты и анализировать свои финансы.

---

## 🚀 Стек технологий

- **Go (Golang)** + Gin
- PostgreSQL
- JWT (аутентификация)
- SQLX
- .env (настройки окружения)
- Graceful shutdown

---

## 🔐 Аутентификация

- Регистрация: `POST /register`
- Логин: `POST /login`
- Авторизация через JWT (в заголовке `Authorization: Bearer <token>`)

---

## 🧾 Основной функционал

### 📌 Расходы
- `POST /api/expenses` — создать
- `GET /api/expenses` — получить все
- `PUT /api/expenses/:id` — обновить
- `DELETE /api/expenses/:id` — удалить
- `GET /api/expenses/categories?from=...` — фильтрация по дате, категории и сумме
- `GET /api/expenses/total` — сумма по фильтру

### 💰 Доходы
- `POST /api/incomes`
- `GET /api/incomes`
- `DELETE /api/incomes/:id`

### 🎯 Финансовые цели
- `POST /api/goals`
- `GET /api/goals`
- `PUT /api/goals/:id` — пополнить цель
- `DELETE /api/goals/:id`

### 📊 Отчеты
- `GET /api/reports/daily`
- `GET /api/reports/weekly`
- `GET /api/reports/monthly`

### 🔒 Бюджет (лимиты)
- `POST /api/budgets`
- `GET /api/budgets`
- `GET /api/budgets/status` — проверка превышения

### 🔁 Повторяющиеся расходы
- `POST /api/recurring`
- `GET /api/recurring`
- `DELETE /api/recurring/:id`

### 🧩 Пользовательские категории
- `POST /api/categories`
- `GET /api/categories`
- `DELETE /api/categories/:id`

---

# 📊 Expense Tracker API

Приложение для отслеживания доходов, расходов, финансовых целей и анализа бюджета.

## 🚀 Возможности

- 🔐 Регистрация и авторизация (JWT)
- 💸 CRUD по расходам и доходам
- 🎯 Финансовые цели
- 🧾 Отчёты: день / неделя / месяц
- 📅 Повторяющиеся расходы (например, аренда)
- 📂 Фильтрация по дате, категории, сумме
- 💰 Бюджеты по категориям
- 🏷️ Пользовательские категории
- 📈 Баланс (доходы - расходы)
- 🔄 Swagger-документация
- ✅ Graceful shutdown сервера

## 📦 Стек технологий

- Go (Golang)
- Gin (Web-фреймворк)
- PostgreSQL (База данных)
- sqlx (Работа с БД)
- JWT (Аутентификация)
- Swagger (Документация)
- Docker (по желанию)

## 🛠 Установка и запуск

1. Клонируй репозиторий:
   ```bash
   git clone https://github.com/your-username/expense-tracker.git
   cd expense-tracker
