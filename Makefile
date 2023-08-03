add-network:
	docker network create c_network
build:
	docker-compose build
build-nc:
	docker-compose build --no-cache
up:
	docker-compose up
down:
	docker-compose down
sh:
	docker-compose exec c-bakend sh
log:
	docker-compose logs -f c-bakend
mysql:
	docker-compose exec mysql mysql -uroot -ppasswordroot -Ddb
