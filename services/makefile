docker-up: docker-run-ping docker-run-pong

docker-run-ping: docker-build-ping
	docker run -dp 8001:8001 ping

docker-run-pong: docker-build-pong
	docker run -dp 8002:8002 pong

docker-build-ping:
	docker build ./ping --tag ping

docker-build-pong:
	docker build ./pong --tag pong