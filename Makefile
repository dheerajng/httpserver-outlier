.PHONY: build clean

build:
	docker build -t httpserver-outlier:latest -f Dockerfile .

clean:
	docker rm -f httpserver-outlier

run: 
	docker run --name httpserver-outlier -dt --rm -p 4040:4040 httpserver-outlier:latest 

docker-build-run: build
	(docker rm -f httpserver-outlier) || true
	docker run --name httpserver-outlier -dt --rm -p 4040:4040 httpserver-outlier:latest 
	 
