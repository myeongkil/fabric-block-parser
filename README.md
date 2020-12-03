# Hyperledger Fabric BlockParser

## Quick start
```
go get
go build
./fabric-block-parser -h
```

## Example
``` shell
# Fetching block
# docker exec -i -t <your_peer_container> peer channel fetch <newest|oldest|config|(number)> [outputfile] [flags]

docker exec -i -t peer0.org1.example.com peer channel fetch newest dc_newest.block -c dc

# logs
# ------
# 2020-12-03 07:55:16.211 UTC [channelCmd] InitCmdFactory -> INFO 001 Endorser and orderer connections initialized
# 2020-12-03 07:55:16.214 UTC [cli.common] readBlock -> INFO 002 Received block: 1
```

``` shell
# Copy block to host
# docker cp <your_peer_container>:/<block_path> ./

docker cp peer0.org1.example.com:/opt/gopath/src/github.com/hyperledger/fabric/peer/dc_newest.block ./
```

``` shell
# Get block info
# ./fabric-block-parser info <block_name>

./fabric-block-parser info dc_newest.block

# logs
# ------
# CurrentHash - using LSH: PD+1AReN2lI4Rqan4AoaCHN5i4xsiggyCIMzIKSAixc=
# Header:
#     Number: 1
#     PreviousHash: JJZLIU7NUrZu488klZQ1qXKCrotVFdqCRJV0BqPoas4=
#     DataHash: SyjV20FFUsxhX9d0+s6rYNoZspWV3jL6/GlePRekkDk=
# Data:
#     Transaction 0
#         Signature: MEQCIBy4319QiWbETcZTe82Quqm5HqqVwANkSkbmKhzcwGtKAiAv5xxWSZxRV/DlWkQLLg44B8ZV2LPEKrru90RacJLDYg==
#         Version: 0
#         Timestamp: seconds:1606981628  nanos:476270289
#         TxId: aa17fc3397a67f1c5ef3976b5c01f199742ec802d6b7a4e77fbc16eee3d9a86d
#         Channel ID: dc
#         Creator:
# devOrg�-----BEGIN CERTIFICATE-----
# MIICLzCCAdagAwIBAgIRAP/u4NrhlG8JIgPoxr3wIw4wCgYIKoZIzj0EAwIweDEL
# MAkGA1UEBhMCS1IxDjAMBgNVBAgTBUJ1c2FuMQ4wDAYDVQQHEwVCdXNhbjEbMBkG
# A1UEChMSZGV2LnNtYXJ0bTJtLmNvLmtyMQwwCgYDVQQLEwNkZXYxHjAcBgNVBAMT
# FWNhLmRldi5zbWFydG0ybS5jby5rcjAeFw0yMDEyMDMwNzQxMDBaFw0zMDEyMDEw
# NzQxMDBaMGwxCzAJBgNVBAYTAktSMQ4wDAYDVQQIEwVCdXNhbjEOMAwGA1UEBxMF
# QnVzYW4xGjAKBgNVBAsTA2RldjAMBgNVBAsTBWFkbWluMSEwHwYDVQQDDBhBZG1p
# bkBkZXYuc21hcnRtMm0uY28ua3IwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATB
# pUS49roKCirFr0OUu7DhML79d3DaHFhCkw9ffXBErCe3irQcexZLnC3F6/UClNcl
# ZlPqyzbwFVsfx8twzN08o00wSzAOBgNVHQ8BAf8EBAMCB4AwDAYDVR0TAQH/BAIw
# ADArBgNVHSMEJDAigCBeiGdyh/I4DeRJchYB8w77iTpj6c1/h5unSOcACJdSpTAK
# BggqhkjOPQQDAgNHADBEAiBzH9fCVD/1vw6DKNz7NrSRoVNaiagN74RwnHsBhEVR
# BwIgYqtBEVbpdwOw926wJvja74/jYo40adBRACdGa0fd0Ds=
# -----END CERTIFICATE-----

```

``` shell
# Get block hash
# ./fabric-block-parser hash <block_name>

./fabric-block-parser hash dc_newest.block

# logs
# ------
# CurrentHash: PD+1AReN2lI4Rqan4AoaCHN5i4xsiggyCIMzIKSAixc=
# PreviousHash: JJZLIU7NUrZu488klZQ1qXKCrotVFdqCRJV0BqPoas4=
# Computed Data Hash: SyjV20FFUsxhX9d0+s6rYNoZspWV3jL6/GlePRekkDk=
```