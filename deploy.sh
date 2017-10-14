function build {
	ROOTDIR=$(PWD)
	GOOS=linux GOARCH=amd64 go build -o helloapp main.go
	cd $ROOTDIR/front/ && yarn && yarn build
	cd $ROOTDIR && docker build -t sunfmin/helloapp .
	rm helloapp
}

function push {
	TAG=$(git rev-parse HEAD|cut -c 1-6)
	docker tag sunfmin/helloapp hub.c.163.com/sunfmin/helloapp:$TAG
	docker push hub.c.163.com/sunfmin/helloapp:$TAG
}

push
