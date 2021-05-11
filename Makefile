all: clean build run

build:
	go build -ldflags "-s -w -extldflags '-static'" -o crm-users

clean:
	rm -f crm-users

run:
	./crm-users server