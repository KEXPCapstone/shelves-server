#!/bin/bash
# note, for this to work properly in production,
# this script should be copied to the deployment server
echo "running mongoimport"
mongoimport --host mongodb --db library --collection releases --type json -v --file /data/seeds/dalet_deploy.json --jsonArray
echo "creating artists collection from release data"
mongo mongodb/library --eval 'db.releases.aggregate([{$group: {_id: "$KEXPReleaseGroupMBID",artist:{$first:{$arrayElemAt: ["$artist-credit", 0]}},releases: {$push: {id: "$_id"}}}},{$group: {_id: "$artist.artist.id",artistName: {$first: "$artist.artist.name"},artistSortName: {$first: "$artist.artist.sort-name"},disambiguation: {$first: "$artist.artist.disambiguation"},releaseGroups:{$push:{releaseGroupId: "$_id",releases: "$releases"}}}},{$sort:{_id:1}},{$out: "artists"}],{collation:{locale: "en", strength: 1}})'

# mongo aggregate
# (same as above, but for readability)
# db.releases.aggregate([
#     {
#         $group: {
#             _id: "$KEXPReleaseGroupMBID",
#             artist: {
#                 $first: {
#                     $arrayElemAt: ["$artist-credit", 0]
#                 }
#             },
#             releases: {
#                 $push: {
#                     id: "$_id"
#                 }
#             }
#         }
#     },
#     {
#         $group: {
#             _id: "$artist.artist.id",
#             artistName: {
#                 $first: "$artist.artist.name"
#             },
#             artistSortName: {
#                 $first: "$artist.artist.sort-name"
#             },
#             disambiguation: {
#                 $first: "$artist.artist.disambiguation"
#             },
#             releaseGroups: {
#                 $push: {
#                     releaseGroupId: "$_id",
#                     releases: "$releases"
#                 }
#             }
#         }
#     },
#     {
#         $sort:{_id:1}
#     },
#     {
#         $out: "artists"
#     }],{collation:{locale: "en", strength: 1}})
