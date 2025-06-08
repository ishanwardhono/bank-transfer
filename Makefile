.PHONY: install clean

install:
	docker-compose up -d

deploy:
	docker-compose up -d --no-deps --build app

clean:
	docker-compose down -v