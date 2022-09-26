run: env
	go run main.go

test:
	go test ./tests

env:
	[ -e ".env" ] && echo "Env Exists" || cp .env.example .env

clean:
	rm .env
