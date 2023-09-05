run:
	go run .

build:
	docker build -t skdf:latest .

it:
	docker run -it skdf:latest sh

swag:
	swagger generate spec -o swagger/swagger.json --scan-models

push:
	git add -A && git commit -m "deploy" && git push origin Imomali/develop