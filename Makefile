.PHONY : cli factorial clean test

grpc: server client
	./calculater_server

server: factorial
	go build -o calculater_server rpc_factorial/server/*.go 

client:	cli 
	go build -o factorial_client rpc_factorial/client/*.go

factorial: 
	go build factorial/*.go

cli: 
	go build cli/*.go

test:
	go test ./cli/ ./factorial/

clean:
	rm ./factorial_client
	rm ./factorial_client