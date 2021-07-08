
build:
# Build Go and Docker images
# go build -o bin/logger main.go
	docker build -t logger ./

release: build
# Push new image to docker hub as user-specified version AND as latest
# "make ver=0.1 release"
	docker tag logger:latest jabray5/logger:$(ver)
	docker push jabray5/logger:$(ver)
	docker tag logger:latest jabray5/logger:latest
	docker push jabray5/logger:latest


docker-run:
# Run latest image in a docker container
	docker run --name logcontainer jabray5/logger:latest

k8s-deploy:
# Kubernetes
	kubectl apply -f persistent-volume.yaml
	kubectl apply -f pv-claim.yaml
	kubectl apply -f deployment.yaml
	kubectl apply -f service.yaml
	@echo service running on port 30000

k8s-delete:
	kubectl delete svc logger-service
	kubectl delete deployment logger-deployment
	kubectl delete -f pv-claim.yaml
	kubectl delete -f persistent-volume.yaml

k8s-update:
# Restart pods with latest image 
	kubectl rollout restart deployment logger-deployment