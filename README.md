## Описание решения

Решение представлено через docker-compose. Запуск:
```bash
docker-compose up -d --force-recreate
```
Остановка:
```bash
docker-compose down
```

### Контейнеры:
1. mortgage_backend - бэкенд проектируемого сервиса. Доступен из вне на порту 8000.
2. mockbank_backend - мок-сервис api банка. Доступен из вне на порту 9000.
3. mortgage_db - база данных (postgres). Доступна из вне на порту 15432.

### Тестирование
В файле test.http есть примеры запросов. API сервисов соответствует условиям задачи.

### Принципы работы
Основной сервис в mortgage_backend принимает заявки. Каждая заявка сохраняется в базе данных mortgage_db. \
После сохранения заявка попадает в очередь заявок в параллельном потоке. Сервис делает запросы к mockbank_backend выбирая заявки из очереди. \
После получения ответа от mockbank_backend сервис изменяет статус заяки в базе данных и заявка попадает в очередь проверки статусов в другом параллельном потоке.
Сервис выбирает заявки из очереди статусов и делает запросы к mockbank_backend, после получения ответа сервис изменяет статус заявки в базе данных.
В файле docker-compose.yml возможна настройка сервисов через переменные окружения.

## Задание


**Необходимо реализовать сервис по работе с заявками на ипотеку.**


### Требование к заданию:
1. В качестве языка программирования использовать Golang;
2. Для БД использовать MySQL, PostgreSQL;
3. Должна быть возможность реализованное задание запустить
Функционал
Сервис должен реализовать:
1. возможность добавления заявки на ипотеку;
2. возможность выдачи списка существующих заявок на ипотеку;
3. отправлять заявку в банк;
4. периодически проверять статус заявки пока не будет получен конечный статус (approved, rejected, error);


### Требования
Сервис должен соответствовать следующим требованиям:
1. Поля заявки соответствовать полям, принимаемым API Банка.
2. Поля заявки должны валидироваться: \
a. все поля обязательны для заполнения; \
b. телефон в международном формате (+79123456789); \
c. email должен быть корректным.
3. Заявка должна быть сохранена в БД на стороне реализуемого сервиса.
4. Запросы к API Банка можно осуществлять не чаще 1 раза в секунду. Параметр должен быть конфигурируемым.
5. Запросы к API сервиса должны вызывать запросы к API Банка асинхронно.
6. Необходимо предусмотреть ситуацию, при которой API Банка может быть недоступно.
7. Конфигурация сервиса должна быть осуществлена из переменных окружения (environment) среды, в которым будет запущен.
8. Необходима реализация тестов.


### Приложение 1. API Банка
Статусы заявки
Возможны следующие статусы заявки на ипотеку: \
processing - заявка обрабатывается банком \
approved - заявка одобрена банком \
rejected - заявка отклонена банком \
error - в заявке найдена ошибка 


#### Методы
Добавление заявки
POST /request
```
{
  "request": {
      "lastname": <string>,
      "firstname": <string>,
      "middlename": <string>,
      "phone": <string>,
      "email": <string>,
    }
}
```

Все поля обязательны для заполнения.

При успешном добавлении заявки будет получен ответ с http-статусом 201 вида:
```
{
  "request": {
    "id": <string>,
	"status_code": <string>
  }
}
```

Перечень полей ответа:
    • id - string - id заявки в системе банка
    • status_code - string - статус заявки в системе банка

Например:
```json
{
  "request": {
    "id": "f1305753-2a2d-42da-827c-091897f1feb7",
    "status_code": "processing"
  }
}
```

Проверка статуса заявки
GET /request/:id
В ответ будет получен ответ с http-статусом 200 вида:
```
{
  "request": {
    "id": <uuid>,
    "status_code": <string>,
  }
}
```

Например:
```json
{
  "request": {
    "id": "f1305753-2a2d-42da-827c-091897f1feb7",
    "status_code": "approved"
  }
}
```
В случае если заявка не найдена, то будет возвращен ответ с http-статусом 404.
