#CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o testzombie .
default: lowercase
	docker build -f Dockerfile -t valeriogheri/lowercase:latest .

lowercase:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o lowercase

clean:
	rm lowercase
