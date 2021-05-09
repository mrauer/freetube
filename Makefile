build:
	docker build -t freetube:latest .

run:
	docker build -t freetube:latest . && docker run -it --rm -v ${CURDIR}:/usr/src/app/go/src/github.com/mrauer/freetube freetube:latest && docker exec -it freetube:latest

linux:
	go build -i -o bin/linux/freetube main.go

osx:
	env GOOS=darwin GOARCH=amd64 go build -i -o bin/osx/freetube
