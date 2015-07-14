build:
	go get
	go build

build/docker:
	docker build -t jmervine/getdown .
