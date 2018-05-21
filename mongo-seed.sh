#!/bin/bash
# note, for this to work properly in production,
# this script should be copied to the deployment server
echo "running mongoimport"
mongoimport --host mongodb --db library --collection releases --type json -v --file /data/seeds/dalet_deploy.json --jsonArray
echo "creating artists collection from release data"
mongo mongodb/library --eval 'db.releases.aggregate([{$group:{_id:"$KEXPReleaseGroupMBID",artist:{$first:{$arrayElemAt:["$artist-credit",0]}},title:{$first:"$KEXPTitle"},releases:{$push:{id:"$_id",title:"$KEXPTitle",coverArtArchive: "$cover-art-archive"}}}},{$group:{_id:"$artist.artist.id",artistName:{$first:"$artist.artist.name"},artistSortName:{$first:"$artist.artist.sort-name"},disambiguation:{$first:"$artist.artist.disambiguation"},releaseGroups:{$push:{releaseGroupId:"$_id",title:"$title",releases:"$releases"}}}},{$sort:{_id:1}},{$out:"artists"}],{allowDiskUse: true,collation:{locale:"en",strength:1}})'
echo "creating labels collection from release data"
mongo mongodb/library --eval 'db.releases.aggregate([{$addFields:{label:{$arrayElemAt:["$label-info",0]},artist:{$arrayElemAt:["$artist-credit",0]}}},{$group:{_id:"$label.label.id",labelName:{$first:"$label.label.name"},disambiguation:{$first:"$label.label.disambiguation"},releases:{$push:{releaseID:"$_id",releaseGroupID:"$KEXPReleaseGroupMBID",title:"$title",artistName:"$artist.name",artistID:"$artist.artist.id",catalogNumber:"$label.catalog-number"}}},},{$out:"labels"}])'

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
#             title: {
#                 $first: "$KEXPTitle"
#             },
#             releases: {
#                 $push: {
#                     id: "$_id",
#                     title: "$KEXPTitle",
#                     coverArtArchive: "$cover-art-archive"
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
#                     title: "$title",
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

# labels collection
# db.releases.aggregate(
#     [
#         {
#             $addFields: {
#                 label: {
#                     $arrayElemAt: ["$label-info", 0]
#                 },
#                 artist: {
#                       $arrayElemAt: ["$artist-credit", 0]
#                 }
#             }
#         },
#         {
#             $group: {
#                 _id: "$label.label.id",
#                 labelName: {
#                     $first: "$label.label.name"
#                 },
#                 disambiguation: {
#                     $first: "$label.label.disambiguation"
#                 },
#                 releases: {
#                     $push: {
#                         releaseID: "$_id",
#                         releaseGroupID: "$KEXPReleaseGroupMBID",
#                         title: "$title",
#                         artistName: "$artist.name",
#                         artistID: "$artist.artist.id",
#                         catalogNumber: "$label.catalog-number"
#                     }
#                 }
#             },
#         },
#         {
#             $out: "labels"
#         }
#     ]
# )