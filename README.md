# hedron

A Distributed Key-Value store in Go.

## Properties

1. Strongly-Consistent (Data is always consistent on majority of cluster)
2. Leaderless (No elections, comparatively more available)
3. Both gRPC and REST Endpoints.
4. High Throughput.


## Installation and Running

1. Clone the repositoty.
2. Navigate to the cloned folder.
3. Run ```$ ./cmd/build_image.sh``` (Docker must be installed for this to work.)
4. Run ```docker-compose up -d```

This should create a three-node cluster of Hedron.

## How to Use

Every Hedron node exposes two ports, one gRPC and one REST port.
(Go through the docker-compose.yaml file in the base directoy of the repository for more.)

### Communicating using REST

Hedron exposes four APIs:

* /create
* /read
* /update
* /delete

All of them have to be POST requests.
The request body for create and update request should look like:
```
{
    "key": "keyName",
    "value": "anyValue"
}
```
And request body for read and delete requests:
```
{
    "key": "keyName"
}
```

## Insights

* Currently, it is not possible to add a new node to a running cluster. So before starting a cluster, make sure (through HEDRON_FELLOW_NODES env variable) that a node knows all the other nodes that will join the cluster.

* A Hedron node will stop processing requests if majority of the cluster is down. This is done to maintain strong-consistency.

* In very rare scenarios, a node will reply with "Request timed out. Unconfirmed write". It obviously means that the request timed out, but the catch is that the node is uncertain whether the write was completed or not. This has to be checked by another read request.
