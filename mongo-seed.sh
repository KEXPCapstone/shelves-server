#!/bin/bash
# note, for this to work properly in production,
# this script should be copied to the deployment server
echo "running mongoimport"
mongoimport --host mongodb --db library --collection releases --type json -v --file /data/seeds/dalet_deploy.json --jsonArray
echo "creating artists collection from release data"
mongo mongodb/library --eval 'db.releases.aggregate([{$group: {_id: "$KEXPReleaseArtistCredit", releases: {$push: {id: "$_id", KEXPTitle: "$KEXPTitle", KEXPMBID: "$KEXPMBID"}}}},{$sort:{_id:1}},{$out: "artists"}],{collation:{locale: "en", strength: 1}})'
