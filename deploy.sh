# deploy script for continuous deployment
# after a succesful travis build on branch master
# this will live on a digital ocean droplet that is the target for our deployment
docker-compose down
docker-compose pull
docker-compose up -d
