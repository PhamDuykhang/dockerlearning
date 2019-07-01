.SILENT:
up_service:
		CGO_ENABLED=0 go build main.go ;\
		mv main deployment;\
		cd deployment;\
		docker build -t kbank .;\
		rm -rf main
		docker run --publish 8000:8000 kbank