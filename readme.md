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


# Kubernets

- Criar Cluster com Kind
```
  kind create cluster
```

- Pegar as informações do cluster
```
  kubectl cluster-info
```

- Visualizar os nodes
```
  kubectl get nodes
```

- Criar o Pod
```
  kubectl create -f pod.yaml
```

- Apagar o Pod
```
  kubectl delete -f pod.yaml
```

- Visualizar os Pods
```
  kubectl get pods
```

- Visualizar os Pods com mais detalhes
```
  kubectl get pods -o wide
```

- Criar um service
```
  kubectl create -f service.yaml
```

- Visualizar os services
```
  kubectl get svc
```

- Visualizar os services com mais detalhes
```
  kubectl get svc -o wide
```

- Apagar o service
```
  kubectl delete -f service.yaml
```


- Rodar o serviço localmente
```
  kubectl port-forward svc/goapp-service 8888:8888
```