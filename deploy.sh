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
	docker network rm dylan
	docker network create dylan
	docker run -d --name redisServer --network dylan redis
	docker run -d --name mongodb --network dylan mongo
	docker pull kexpcapstone/shelves
	docker rm -f shelves
	docker pull kexpcapstone/library
	docker rm -f library
	docker run -d --name shelves --network dylan -e DBADDR=mongodb:27017 kexpcapstone/shelves
	docker run -d --name library --network dylan -e DBADDR=mongodb:27017 kexpcapstone/library
	docker pull kexpcapstone/gateway
	docker rm -f gateway
	docker run -d \
	-p 443:443 \
	--name gateway \
	-v /etc/letsencrypt:/tls:ro \
	-e SHELVESSVCADDR=shelves:80 \
	-e LIBRARYSVCADDR=library:80 \
	-e TLSCERT=/tls/live/api.kexpcapstone.com/fullchain.pem \
	-e TLSKEY=/tls/live/api.kexpcapstone.com/privkey.pem \
	-e DBADDR=mongodb:27017 \
	-e REDISADDR=redisServer:6379 \
	-e SESSIONKEY=blackcoffee \
	--network dylan \
	kexpcapstone/gateway
	exit
HERE

