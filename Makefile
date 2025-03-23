build:
	go build -o ./dist/tracker .

run: build
	./dist/tracker
