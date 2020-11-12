fmt:
	go fmt -x ./...

lint:
	golangci-lint run ./...

tests:
	go test -v -count=1 ./...

build_web:
	(cd ./web && docker build -t lumio_thetearoundpicker_web:latest .)

build_backend:
	docker build -t lumio_thetearoundpicker_backend:latest .

build: build_web build_backend

run:
	docker-compose up -d

down:
	docker-compose down
