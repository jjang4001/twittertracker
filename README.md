# Set up
1. Clone into ~/go/src directory
2. For each dependency, run:
```
go get <dependency>
ex: go get github.com/gorilla/handlers
```
3. To start the server, run
```
go run main.go
```

To test the websocket connection, open socket.html, and click "Test socket".