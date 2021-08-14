build:
	docker build -t freetube:latest .

dev:
	docker build --progress=plain -t freetube:latest . && docker run -it --rm -v ${CURDIR}:/usr/src/app/go/src/github.com/mrauer/freetube freetube:latest && docker exec -it freetube:latest

binary:
	env GOOS=linux GOARCH=amd64 go build -i -o freetube

release:
	env GOOS=darwin GOARCH=amd64 go build -i -o bin/osx/freetube
	env GOOS=linux GOARCH=amd64 go build -i -o bin/linux/freetube
	env GOOS=windows GOARCH=amd64 go build -i -o bin/windows/freetube.exe
