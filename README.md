### Quick start
Требуется Go не ниже 1.13 с нативной поддержкой go.mod

(Опционально)
cd internal/proto && protoc --go_out=plugins=grpc:. *.proto && cd ../../

Install protobuf

Mac: brew install protobuf
Linux: https://gist.github.com/sofyanhadia/37787e5ed098c97919b8c593f0ec44d8

### Сервер

cd cmd && go build -o main && ./main

### Клиент (тестовый)

cd test && go build -o client && ./client

### Описание

Этот проект часть тестового задания:

https://docs.google.com/document/d/1RI11ZQyfok9su7P-M6JuusyTWgs5K4xO6MtfP40zg8s/edit?usp=sharing

Доступен по урлу 80.93.182.105, порт: 50005

Если есть вопросы, писать @nuralimov (Telegram)