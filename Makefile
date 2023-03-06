project_name = mbase
image_name = gofiber:latest

run-local:
	go run app.go

requirements:
	go mod tidy

clean-packages:
	go clean -modcache

delete-container-if-exist:
	docker stop $(project_name) || true && docker rm $(project_name) || true

docker-stop:
	docker stop $(project_name)

docker-start:
	docker-compose -f docker-compose.yaml up -d --build mbase