
# Gophkeeper

Gophkeeper - это инструмент CLI для управления секретами. Ниже приведены инструкции по установке и использованию.

Для тестирования запустить файл из /cmd/server:
```sh
go run main.go
```

## Генерация CLI

### Linux
```sh
GOOS=linux GOARCH=amd64 go build -o gophkeeper-linux -ldflags "-X main.version=1.0.0 -X 'main.buildDate=$(date)'"
```

### Windows
```sh
GOOS=windows GOARCH=amd64 go build -o gophkeeper-windows.exe -ldflags "-X main.version=1.0.0 -X 'main.buildDate=$(date)'"
```

### MacOS
```sh
GOOS=darwin GOARCH=amd64 go build -o gophkeeper-mac -ldflags "-X main.version=1.0.0 -X 'main.buildDate=$(date)'"
```

## Использование

### Получение версии
```sh
./gophkeeper-linux version
```

### Регистрация пользователя
```sh
./gophkeeper-linux register testuser testpassword
```

### Логин
```sh
./gophkeeper-linux login testuser testpassword
```

### Логин с записью токена в переменную
```sh
TOKEN=$(./gophkeeper-linux login testuser testpassword | awk '{print $NF}')
```

### Проверка токена
```sh
echo $TOKEN
```

### Сохранение данных

#### Сохранение пары логин/пароль
##### Формат данных: website, тип, данные в формате json
```sh
./gophkeeper-linux store "website.com" "login" '{"login":"testlogin","password":"testpassword"}' --token $TOKEN
```

#### Сохранение произвольного текста
##### Формат данных: ключ, тип, строка данных
```sh
./gophkeeper-linux store "somekey" "text" "sometextdata" --token $TOKEN
```

#### Сохранение данных банковской карты
##### Формат данных: имя банка, тип, данные в формате json
```sh
./gophkeeper-linux store "bankname" "bank card" '{"number":"1234123412341234","date":"04-23","cvv":"032"}' --token $TOKEN
```

### Получение данных

#### Получение всех данных
```sh
./gophkeeper-linux get --token $TOKEN
```

#### Получение всех данных по ключу
```sh
./gophkeeper-linux get --key "website.com" --token $TOKEN
```

#### Получение пары логин/пароль по ключу и идентификатору (логину)
```sh
./gophkeeper-linux get --key "website.com" --type "login" --identifier "testlogin" --token $TOKEN
```

#### Получение произвольного текста по ключу
```sh
./gophkeeper-linux get --key "somekey" --token $TOKEN
```

#### Получение данных банковской карты по ключу и номеру карты
```sh
./gophkeeper-linux get --key "bankname" --type "bank card" --identifier "1234123412341234" --token $TOKEN
```

### Обновление данных

#### Обновление пары логин/пароль
```sh
./gophkeeper-linux update --key "website.com" --type "login" --identifier "testlogin" --new_data '{"login":"testlogin","password":"newpassword"}' --token $TOKEN
```

#### Обновление произвольного текста
```sh
./gophkeeper-linux update --key "somekey" --type "text" --new_data "newtextdata" --token $TOKEN
```

#### Обновление данных банковской карты
```sh
./gophkeeper-linux update --key "bankname" --type "bank card" --identifier "1234123412341234" --new_data '{"number":"1234123412341234","date":"05-24","cvv":"123"}' --token $TOKEN
```

### Удаление данных

#### Удаление пары логин/пароль
```sh
./gophkeeper-linux delete --key "website.com" --type "login" --identifier "testlogin" --token $TOKEN
```

#### Удаление произвольного текста
```sh
./gophkeeper-linux delete --key "somekey" --type "text" --token $TOKEN
```

#### Удаление данных банковской карты
```sh
./gophkeeper-linux delete --key "bankname" --type "bank card" --identifier "1234123412341234" --token $TOKEN
```
