server:
	go run main.go
start-dm:
	CompileDaemon -command="./go-mongo-tut"

reown:
	sudo chown -R $(USER) .

.PHONY: server reown start-dm