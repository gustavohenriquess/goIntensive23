# Go 


## start project
```
  go mod init github.com/gustavohenriquess/go-intensive23
```

## Test project
```
  go test ./...
```

## Install packages
```
  go mod tidy
```

## Close DB
```
  defer db.Close()
```
O defer é usado para executar a ação após o retorno da função que o chamou.