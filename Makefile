build:
	docker build -t freetube:latest .

run:
	docker build -t freetube:latest . && docker run -it --rm -v ${CURDIR}:/usr/src/app/go/src/github.com/mrauer/freetube freetube:latest && docker exec -it freetube:latest

binary:
	go build -i -o freetube main.go
