project_name = mbase
image_name = gofiber:latest

run-local:
	go run app.go

requirements:
	go mod tidy

clean-packages:
	go clean -modcache

docker-stop:
	docker-compose -f docker-compose.yaml stop

docker-rm:
	docker-compose -f docker-compose.yaml stop && docker-compose -f docker-compose.yaml rm -f

docker-start:
	docker-compose -f docker-compose.yaml up -d --build mbase

docker-start-only-redis:
	docker-compose -f docker-compose.yaml up -d --build mbase-redis

clean-uploads:
	@if [  -d "static/public/uploads" ];then     \
		rm -rf static/public/uploads;           \
    fi