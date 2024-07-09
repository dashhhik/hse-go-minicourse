Задание
=

Необходимо написать http сервер по работе с аккаунтами и балансами. Для этого можно использовать, как стандартную библиотеку net/http, так и фрейворки echo, gin-gonic, fiber, fasthttp и т.д. Должны быть реализованы 5 методов (получение аккаунта, изменение имени аккаунта, изменение баланса аккаунта, создание и удаление аккаунта). Далее должен быть реализован CLI (command-line interface), он же клиент. В нем для реализации интерфейса использовать можно стандартную библиотеку flag или любую внешнюю, например cobra.

API Методы
=

1. Создание аккаунта
```
POST /account
{
  "name": "string",
}
```

2. Получение аккаунта
```
GET /account/{name}
```

3. Изменение баланса аккаунта
```
PATCH /account/{name}
{
    "balance": 0
}
```

4. Получить все аккаунты
```
GET /accounts
```

5. Удаление аккаунта
```
DELETE /account/{name}
```

6. Изменение имени аккаунта
```
PUT /account/{name}
{
    "name": "string"
}
```

CLI
=

```
cd second-hwk/client

go build -o client
```

1. Создание аккаунта
```
./client create --name "string" 
```

2. Получение аккаунта
```
./client get --name "string"
```

3. Изменение баланса аккаунта
```
./client update --name "string" --balance 0
```

4. Получить все аккаунты
```
./client list
```

5. Удаление аккаунта
```
./client delete --name "string"
```

6. Изменение имени аккаунта
```
./client rename --name "string" --newname "string"
```
