LOAD_BALANCER_NAME="load_balancer"
BACKEND_NAME="backend"
BACKEND_PORT=8000

run_backend:
	docker build -t $(BACKEND_NAME) ./backend
	docker run -p $(BACKEND_PORT):$(BACKEND_PORT) --name $(BACKEND_NAME) $(BACKEND_NAME); \
	docker stop $(BACKEND_NAME); \
	docker rm $(BACKEND_NAME); \
	docker rmi $(BACKEND_NAME);