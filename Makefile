default: lowercase
	docker build -f Dockerfile -t valeriogheri/lowercase:latest .
	docker login -u valeriogheri -p MetopA_2016
	docker tag valeriogheri/lowercase:latest valeriogheri/lowercase:1.0.5
	docker push valeriogheri/lowercase

lowercase:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o lowercase

clean:
	rm lowercase
