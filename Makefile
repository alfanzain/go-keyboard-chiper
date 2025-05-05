.PHONY: *

run:
	-docker compose -p "keyboard-chiper" -f ./deploy/local/run/docker-compose.yml down --remove-orphans
	docker compose -p "keyboard-chiper" -f ./deploy/local/run/docker-compose.yml up --build --attach=server
