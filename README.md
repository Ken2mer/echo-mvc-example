# echo-mvc

## set database url

```go
// main.go

var (
	// set postgres url (example)
	dbURL = "postgres://user:pass@host:port/dbname?sslmode=disable"
)
```

## run app

```zsh
% go run main.go
```

### create user

```zsh
% curl -X POST -H "Content-Type: application/json" -d '{"name":"example 1"}' localhost:1323/users
% curl -X POST -H "Content-Type: application/json" -d '{"name":"example 2"}' localhost:1323/users
% curl -X POST -H "Content-Type: application/json" -d '{"name":"example 3"}' localhost:1323/users
```

### Get users

```zsh
% curl -s http://localhost:1323/users | jq .
[
  {
    "user_id": 1,
    "name": "example 1"
  },
  {
    "user_id": 2,
    "name": "example 2"
  },
  {
    "user_id": 3,
    "name": "example 3"
  },
]
```
