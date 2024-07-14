

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


4. Удаление аккаунта
```
./client delete --name "string"
```

5. Изменение имени аккаунта
```
./client rename --name "string" --newname "string"
```
