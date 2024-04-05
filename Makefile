make_server:
	go run main.go

reown:
	sudo chown -R $(USER) .

.PHONY: make_server reown