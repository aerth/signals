all:
	go build sendsig.go
	sudo install sendsig /usr/local/bin/
