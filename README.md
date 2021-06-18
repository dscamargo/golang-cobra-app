Criação/Listagem de usuários via web server e CLI.

- Gin Framework
- Sqlite
- Cobra CLI

```sh
go run main.go [command] [argumentos]
```

Comandos:

| Comando     | Argumentos              | Padrão                    | Descrição                                 |
|-------------|-------------------------|---------------------------|-------------------------------------------|
| http        | --port                  | 8080                      | Roda um webserver                         |
| newuser     | --name<br>--email       | user<br>user@email.com.br | Cria um usuário no banco de dados         |
| listUsers   |                         |                           | Lista todos os usuários cadastrados no DB |
| createTable | --dbName<br>--tableName | user<br>test.db           | Cria a tabela no banco SQL                |
