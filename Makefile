server=124.223.8.237
all: build deploy  clean
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server .
deploy:
	ssh -P 3322  root@$(server)  "systemctl stop server"
	scp -P 3322 -r server root@$(server):/server
	ssh -P 3322 root@$(server)  "systemctl start server"
#deploy_config:
#	scp  config.yml root@124.223.8.237:/server
clean:
	rm -rf server

# 本地编译打包上传
#TAG=v0.0.1
#all: build deploy
#
#build:
#	docker build -t server:${TAG} .
#	docker save -o server.tar server:${TAG}
#
#deploy:
#	scp -r server.tar root@124.223.8.237:/server
#	ssh root@124.223.8.237 "docker load -i /server/server.tar &&\
# 							docker ps -af name=apiserver -q | xargs --no-run-if-empty docker rm -f &&\
# 							docker run -d --name=apiserver --restart=always -p 8989:8989 -v /etc/localtime:/etc/localtime -v /server/files:/opt/files server:${TAG}"

# 服务上编译打包
#TAG=v0.0.1
#all: build deploy clean
#
#build:
#	scp -r ../server root@124.223.8.237:/tmp
#	ssh root@124.223.8.237 "cd /tmp/server && docker build -t server:${TAG} . && rm -rf /tmp/server"
#
#deploy:
#	ssh root@124.223.8.237 "docker ps -af name=apiserver -q | xargs --no-run-if-empty docker rm -f &&\
# 							docker run -d --name=apiserver --restart=always -p 8989:8989 -v /etc/localtime:/etc/localtime -v /server/files:/opt/files server:${TAG}"
#
#clean:
#	ssh root@124.223.8.237 "rm -rf /tmp/server"