# todo_list

ハッカソンの練習用

## 動かし方
```sh
$ docker compose up -d
```
- http://localhost:8080 -> バックエンドアプリ
- http://localhost:8000 -> adminer
  - mysqlを操作できるやつ
  - host: mysql
  - user: root
  - password: root
  - database: todo_list

### リクエストを送る
例
```sh
$ curl -v http://localhost:8080/todos/1
$ curl -v -H "Content-Type: application/json" -d '{"task": "do my homework", "status": "processing"}' http://localhost:8080/todos
$ curl -v -X PATCH -H "Content-Type: application/json" -d '{"status": "done"}' http://localhost:8080/todos/1
```
