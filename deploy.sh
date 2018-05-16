# deploy script for continuous deployment
# after a succesful travis build on branch master
# note, for this to work properly in production,
# this script should be copied to the deployment server
docker-compose down
docker-compose pull
docker-compose up -d
# clear out unused images
yes | docker image prune -a
# restart the library service (race condition workaround)
docker restart library
