test:
	go test -tags testing ./...

cover:
	go test -tags testing -coverprofile=cover.out ./...
	go tool cover -html=cover.out -o cover.html

clean:
	rm -rf bin/*
	rm cover.*