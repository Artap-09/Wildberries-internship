# Level 0
## Содержание
1. [Описание](#описание)
2. [Запуск](#запуск)  
    * [Postgres](#postgres)
    * [NATS-streaming](#nats-streaming)
    * [Сервис](#сервис)  
3. [Интерфейс](#интерфейс)
4. [Docker-compose](#docker-compose)
5. [TODO](#todo)

## Описание
![https://img.shields.io/badge/version-v1.0.0-green](https://img.shields.io/badge/version-v1.0.0-green) ![https://img.shields.io/badge/Golang-v1.17.6-blue](https://img.shields.io/badge/Golang-v1.17.6-blue) ![https://img.shields.io/badge/Docker--compose-does%20not%20work-red](https://img.shields.io/badge/Docker--compose-does%20not%20work-red)

ТЗ можно посмотреть в файле [L0.pdf](L0.pdf)

Кратко:
- [x] Сервис который получает данные по подписке в NATS-streaming и отправляет на сервер по id
  - [x] Данные хранит в Кеш и пишет в postgres
  - [x] При падении берет данные из БД
- [x] Сервер выдает данные по id
  - [x] Интерфейс (html)

## Запуск
Запуск условно можно разделить на 3 этапа:
1. [Postgres](#postgres)
2. [NATS-streaming](#nats-streaming)
3. [Сервис](#сервис)

:exclamation: запуск через docker-copmpose пока не работает
### Postgres
`Postgres` поднимаем в `docker` командой:
```
docker run --name PSQL -p 5432:5432 -e POSTGRES_USER=myuser -e POSTGRES_PASSWORD=qwerty -e POSTGRES_DB=l0 -d postgres:14
```
Можете поменять `--name PSQL` эта опция дает название контейнеру с `Postgres`  
`-e POSTGRES_USER=myuser` можете заменить `myuser` на своего user, так же можно изменить пароль `-e POSTGRES_PASSWORD=qwerty`

:exclamation::exclamation::exclamation: Не меняйте `-p 5432:5432` `-e POSTGRES_DB=l0` эти параметры пока что невозможно настроить в сервисе

После создания заходим в `bash` в контейнере:
```
docker exec -it PSQL bash
```
Если меняли `--name PSQL` введите свое название контейнера или его `ID`

Теперь заходим в `PSQL`
```
psql -U myuser -W -d l0
```
Появится строка
```
Password: 
```
Введите пароль, каторый указывали при запуске контейнера.  
:exclamation: Вы не будете видеть вводимые символы

Теперь создаем таблицу в БД `l0`:
```SQL
CREATE TABEL taskl0(
order_uid text,
data json);
```
На этом все, можете выходить из `PSQL` и `bash`, сделать это можно командами:  
Для PSQL:
```
\q
```
Для Bash
```
exit
```

### NATS-streaming
`NATS-streaming` тоже поднимаем в `docker`  
Для этого надо запустить команду:
```
docker run --name NATS -p 4222:4222 -p 8222:8222 -d nats-streaming 
```

:exclamation::exclamation::exclamation: Тут не меняйте порты `-p 4222:4222 -p 8222:8222`
### Сервис

Запустим сервис, находясь в папке `Исходники`, командой:
```
go run L0.go
``` 
В терминале вы должны увидеть:  
![Alt-текс](png/%D0%A1%D0%BD%D0%B8%D0%BC%D0%BE%D0%BA%20%D1%8D%D0%BA%D1%80%D0%B0%D0%BD%D0%B0%20%D0%BE%D1%82%202022-02-08%2017-33-39.png)

Введите `POSTGRES_USER` а затем `POSTGRES_PASSWORD` которые указывали выше.  
Если все успешно, то:  
![Alt-текс](png/Снимок%20экрана%20от%202022-02-08%2017-43-50.png)

:exclamation: сервис подписывается на `foo1` в `NATS-streaming`, поэтому для отправки данных используйте это же название.

## Интерфейс
Для подключения на сервер в браузере введите:
```
localhost:4969
```

Вам должен открыться интерфейс сервиса (html страница) 
![Alt-текс](png/Снимок%20экрана%20от%202022-02-08%2017-56-24.png)

После в поле введите нужный `Order UID`
![Alt-текс](png/Снимок%20экрана%20от%202022-02-08%2017-56-30.png)

и нажмите кнопку `Показать`
![Alt-текс](png/Снимок%20экрана%20от%202022-02-08%2017-56-36.png)

## Docker-compose

Сервис компилится в контейнере  
Но при запуске не хочет подключаться к БД (не совпадают пароли)

## TODO

- [ ] Сделать запуск через `docker-copmpose`
- [ ] Добавить скрипт создания таблицы
- [ ] Добавить очистку экрана для этапа входа в БД
- [ ] Сделать возможность полной ностройки сервиса
  - [ ] Postgres
  - [ ] NATS-streaming
  - [ ] Сервера
- [ ] Добавить в интерфейс возможность посмотреть какие `Order UID` существуют
- [ ] Написать тесты
- [ ] Провести проверку на прочность (сервера, сервиса)
