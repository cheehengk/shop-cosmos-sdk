# Shop
**Shop** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve
```
## Consensus breaking
In the following sections, we will discuss how consensus was broken in the Shop blockchain. A consensus-breaking
change is an update or modification to the blockchain protocol that is not compatible with the existing rules. This
means that nodes running the old version of the software would not agree with the nodes running the new version on
the validity of new transactions or blocks, potentially causing a fork in the blockchain.

Initially, the Shop blockchain was created with a simple Post type which contains metadata and data fields like a
title, description, price, date, seller, and a unique ID. However, the Shop blockchain was later updated to include
a new field, status, in the Post type. The status accepts either "Available", "Reserved", or "Sold".

Since this new field was only implemented as a string datatype (and not an enumerator with default value), this
means that there will be a break in consensus between the old and new versions of the Shop blockchain. Nodes running
the old version of the Shop blockchain will not be able to validate the new blocks created by nodes running the new
version. This is because the old nodes will not be able to recognize the new field, status, in the Post type.
Without a properly integrated status field in the old nodes, the new nodes will also not be able to validate the the
old blocks created by the old nodes, especially since the the handlers are updated to check for the status field by
string comparison.

## Basic commands
### Create Post
shopd tx shop create-post {title} {description} {price} --from {account} --chain-id shop

Example: 

```shopd tx shop create-post item-1 this-is-item-1 100 --from Alice --chain-id shop```

### List Posts
shopd q shop list-post {start_unix_time} {end_unix_time} {min_price} {max_price} --chain-id shop

*Note: start_unix_time and end_unix_time are unix timestamps. If you want to ignore them, use 0.*

Example: 

```shopd q shop list-post 0 0 0 0 --chain-id shop```

### Show Post
shopd q shop show-post {post_id} --chain-id shop

Example: 

```shopd q shop show-post 4 --chain-id shop```

### Update Post
shopd tx shop update-post {title} {description} {price} {post_id} --from {account} --chain-id shop

Example: 

```shopd tx shop update-post new-title new-description 200 4 --from Alice --chain-id shop```

### Delete Post
shopd tx shop delete-post {post_id} --from {account} --chain-id shop

Example: 

```shopd tx shop delete-post 4 --from Alice --chain-id shop```

### Change Status (only on branch: "feat/consensus-breaking")
shopd tx shop change-post-status {post_id} {status} --from {account} --chain-id shop

*Note: status is either "Available", "Reserved", or "Sold". If invalid status is provided, it will default to 
current status.*

Example: 

```shopd tx shop change-post-status 4 Sold --from Alice --chain-id shop```

