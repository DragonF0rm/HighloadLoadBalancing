LOAD_BALANCER_NAME="load_balancer"
LOAD_BALANCER_PORT=80
BACKEND_NAME="backend"
BACKEND_PORT=8000

run_backend:
	docker build -t $(BACKEND_NAME) ./backend
	docker run -p $(BACKEND_PORT):$(BACKEND_PORT) --name $(BACKEND_NAME) $(BACKEND_NAME)

stop_backend:
	docker stop $(BACKEND_NAME)

clear_backend:
	docker rm $(BACKEND_NAME)
	docker rmi $(BACKEND_NAME)

run_lb:
	docker build -t $(LOAD_BALANCER_NAME) ./balancer
	docker run -p $(LOAD_BALANCER_PORT):$(LOAD_BALANCER_PORT) -p 9901:9901 \
	--name $(LOAD_BALANCER_NAME) $(LOAD_BALANCER_NAME)

stop_lb:
	docker stop $(LOAD_BALANCER_NAME)

clear_lb:
	docker rm $(LOAD_BALANCER_NAME)
	docker rmi $(LOAD_BALANCER_NAME)
