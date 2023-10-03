.PHONY: pull run

pull:
	docker pull mongo:latest

run:
	docker run -d -p 27017:27017 --name=mongo-example mongo:latest
