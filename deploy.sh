#!/usr/bin/env bash
cd gateway
./build.sh
cd ../libray
./build.sh
cd ../shelves
./build.sh
cd ..
echo "Deploying to DigitalOcean Droplet; You may be prompted to enter your SSH passphrase:"
ssh root@api.kexpshelves.com << HERE
	docker rm -f shelves
	docker rm -f library
	docker rm -f gateway
	docker rm -f redissvr
	docker network rm dylan
	docker network create dylan
	docker run -d --name redissvr --network dylan redis
	docker run -d --name mongodb -v /srv/data:/data/seeds --network dylan mongo
	docker pull kexpcapstone/shelves
	docker pull kexpcapstone/library
	docker run -d --name shelves --network dylan -e DBADDR=mongodb:27017 kexpcapstone/shelves
	docker run -d \ 
	--name library \ 
	-e DBADDR=mongodb:27017 \
	-e RELEASEDB=library \
	-e RELEASECOLL=releases \
	--network dylan \
	kexpcapstone/library
	docker pull kexpcapstone/gateway
	docker run -d \
	-p 443:443 \
	--name gateway \
	-v /etc/letsencrypt:/tls:ro \
	-e SHELVESSVCADDR=shelves:80 \
	-e LIBRARYSVCADDR=library:80 \
	-e TLSCERT=/tls/live/api.kexpshelves.com/fullchain.pem \
	-e TLSKEY=/tls/live/api.kexpshelves.com/privkey.pem \
	-e DBADDR=mongodb:27017 \
	-e REDISADDR=redissvr:6379 \
	-e SESSIONKEY=blackcoffee \
	--network dylan \
	kexpcapstone/gateway
	exit
HERE

