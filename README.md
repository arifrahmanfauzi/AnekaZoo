# Aneka Zoo Golang CRUD API

### Requirement
- Golang 1.17.8 or latest
- Mysql 8.0 or latest

### Installation Instruction
- craete MYSQL database name it `anekazoo`
- import sql file `anekazoo.sql`
- run  `go run main.go`

### List End Point
- /v1/animal/create
- /v1/animal/get
- /v1/animal/get/{id}
- /v1/animal/update/{id}
- /v1/animal/delete/{id}