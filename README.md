# Online chat

![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/ShatAlex/TradingApp)
![Static Badge](https://img.shields.io/badge/gin-v1.9.1-brightgreen)
![Static Badge](https://img.shields.io/badge/gorilla/websocket-v1.5.0-yellow)
![Static Badge](https://img.shields.io/badge/sqlx-v1.3.5-brown)



## :sparkles: Project Description
The project is designed to exchange user messages in chat rooms over websocket connections.

Implemented an authentication system, creating new chat rooms and adding new users to them.
___

## :clipboard: Usage
To Run the application use:
```
make run
```
If the application is running for the first time, you need to apply migrations to the db:
```
make migrate
```