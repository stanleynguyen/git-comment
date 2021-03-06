start:
	docker stack deploy -c docker-compose.yml gitcommentapp
start_dev:
	docker-compose -f docker-compose.dev.yml -p gitcommentapp up -d
start_dev_w_log:
	docker-compose -f docker-compose.dev.yml -p gitcommentapp up
rebuild_dev:
	docker-compose -f docker-compose.dev.yml -p gitcommentapp build
build: build_comment build_member build_proxy
build_member:
	pushd $(shell pwd)/member-app && \
	docker build -t stanleynguyen/gitcomment_member . && \
	docker image push stanleynguyen/gitcomment_member && \
	popd
build_comment:
	pushd $(shell pwd)/comment-app && \
	docker build -t stanleynguyen/gitcomment_comment . && \
	docker image push stanleynguyen/gitcomment_comment && \
	popd
build_proxy:
	pushd $(shell pwd)/nginx && \
	docker build -t stanleynguyen/gitcomment_proxy . && \
	docker image push stanleynguyen/gitcomment_proxy && \
	popd
stop:
	docker stack rm gitcommentapp
stop_dev:
	docker-compose -f docker-compose.dev.yml -p gitcommentapp down
viz:
	docker container run -p 8080:8080 -v /var/run/docker.sock:/var/run/docker.sock -d dockersamples/visualizer
