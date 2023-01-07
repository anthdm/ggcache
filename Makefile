build:
	go build -o bin/ggcache

run: build
	./bin/ggcache

runfollower: build
	./bin/ggcache --listenaddr :4000 --leaderaddr :3000