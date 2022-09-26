run: env
	go run main.go

test:
	go test ./tests

env:
	[ -e "./config/default.yaml" ] && echo "Env Exists" || cp ./config/example.default.yaml ./config/default.yaml

clean:
	rm ./config/default.yaml
