all: clean build run

build:
	go build -ldflags "-s -w -extldflags '-static'" -o searcher

clean:
	rm -f searcher

run:
	./searcher server