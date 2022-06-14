# Clean Architecture With Golang

- When init a new project
```shell
go mod init github.com/samuelterra22/clean-architecture-go
```

- Run testes
```shell
go test ./...
```

- Generate a mock
```shell
mockgen -source=entity/repository.go -destination=entity/mock/mock.go
```

```shell
 go run cmd/main.go 
```

```shell
docker-compose up -d
```

```shell
docker exec -it aluno_app_1 bash
```

```shell
docker exec -it aluno_kafka_1 bash 
kafka-console-producer --bootstrap-server=localhost:9092 --topic=transactions
kafka-console-consumer --bootstrap-server=localhost:9092 --topic=transactions_result
{"id": "123","account_id": "1","credit_card_number": "40000000000000000","credit_card_name": "Samuel Terra","credit_card_expiration_month": 12,"credit_card_expiration_year": 2024,"credit_card_expiration_cvv": 123,"amount": 1200}
{"id": "123","account_id": "1","credit_card_number": "4193523830170205","credit_card_name": "Samuel Terra","credit_card_expiration_month": 12,"credit_card_expiration_year": 2024,"credit_card_expiration_cvv": 123,"amount": 1200}
```


```shell
nest g controller MyController
```

```shell
nest g resource orders
```

```shell
nest g service accounts/account-storage
```

```shell
nest g guard accounts/token 
```

```shell
127.0.0.1       host.docker.internal
```