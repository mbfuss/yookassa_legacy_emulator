# Yookassa Legacy Emulator 0.0.1-alpha

**Yookassa Legacy Emulator** — это эмулятор для работы с устаревшим API ЮKassa. Проект позволяет обрабатывать запросы на регистрацию заказа и отображать страницу оплаты.

## Стек технологий

- **Язык**: Go
- **Фреймворки и библиотеки**:
  - `net/http` — для обработки HTTP-запросов.
  - `github.com/google/uuid` — для генерации уникальных идентификаторов.
  - `github.com/joho/godotenv` — для работы с переменными окружения.
- **Шаблоны**: HTML + CSS.

## Структура проекта

- `internal/business/parser/` — логика парсинга и валидации входящих данных.
- `internal/business/utils/` — вспомогательные утилиты, например, генерация уникального идентификатора заказа.
- `internal/config/` — загрузка конфигурации из `.env` файла.
- `internal/di/` — инициализация зависимостей и запуск HTTP-сервера.
- `internal/handlers/` — обработчики HTTP-запросов.
- `internal/resource/` — структуры для входящих и исходящих данных.
- `templates/` — HTML и CSS для отображения страницы оплаты.

## Установка и запуск

1. **Клонируйте репозиторий:**
   ```bash
   git clone https://github.com/mbfuss/yookassa_legacy_emulator.git
   ```

2. **Установите зависимости:**
   ```bash
   go mod tidy
   ```
3. **Запустите сервер:**
   ```bash
   go run main.go
   ```

## Основные маршруты

- **Регистрация заказа**: `POST /payment/rest/registerPreAuth.do`
   - Принимает данные заказа и возвращает `orderId` и `formUrl`.
- **Страница оплаты**: `GET /payment/rest/payments?orderId={orderId}`
   - Отображает страницу оплаты с кнопкой для проведения оплаты.

## Пример запроса

### Регистрация заказа
**URL**: `http://localhost:9010/payment/rest/registerPreAuth.do`

**Метод**: `POST`

**Тело запроса в формате x-www-form-urlencoded**:
```json
  "userName": "412345",
  "password": "QWERTY123",
  "orderNumber": "12345",
  "amount": "1000",
  "returnUrl": "http://example.com/success",
  "failUrl": "http://example.com/fail",
  "email": "null",
  "phone": "91234567890",
  "jsonParams": {"email":null, "phone":"91234567890"},
  "orderBundle": {"customerDetails":{"email":null,"phone":"91234567890","fullName":"\u0422\u0435\u0441\u0442\u0438\u0440\u043e\u0432\u0430\u043d\u0438\u0435 \u0422\u0435\u0441\u0442\u0438\u0440\u043e\u0432\u0430\u043d\u0438\u0435 "},"cartItems":{"items":[{"positionId":0,"name":"\u041c\u043e\u0434\u0443\u043b\u044c \u043f\u0430\u043c\u044f\u0442\u0438 eMMC 32 GB (\u0422\u0435\u0441\u0442\u043e\u0432\u044b\u0439 \u0442\u043e\u0432\u0430\u0440)","quantity":{"value":"1","measure":"0"},"tax":{"taxType":0},"itemPrice":9800,"itemAttributes":{"attributes":[{"name":"paymentMethod","value":"1"},{"name":"paymentObject","value":"1"}]}},{"positionId":1,"name":"\u0414\u043e\u0441\u0442\u0430\u0432\u043a\u0430 \u0421\u0414\u042d\u041a","quantity":{"value":1,"measure":"0"},"itemPrice":24500,"tax":{"taxType":0},"itemAttributes":{"attributes":[{"name":"paymentMethod","value":"1"},{"name":"paymentObject","value":"4"}]}}]}}, 
  "taxSystem": "0"
```

**Пример ответа**:
```json
{
  "orderId": "550e8400-e29b-41d4-a716-446655440000",
  "formUrl": "http://localhost:8080/payment/rest/payments?orderId=550e8400-e29b-41d4-a716-446655440000"
}
```
C устаревшим API ЮKassa возможно ознакомиться по ссылке:
[устаревшее API ЮKassa](https://yoomoney.ru/i/forms/yc-program-interface-api-sberbank.pdf).
## Лицензия

Проект распространяется под лицензией MIT.
