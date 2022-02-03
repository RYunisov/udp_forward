udp_forward
-----------

# About tool

I want to check how GOlang work with network...
It tool to do UPD forward from one host to other

# Prepare

* Go v1.17.3

# Build

* Run from work directory
```sh
go build 
```

# Run

* Run prepared binary
```sh 
udp_forward 127.0.0.1 1024 8.8.8.8 53
```
Args:
* Listen host is 127.0.0.1 and listen port is 53
* Forward to host is 8.8.8.8 and forward port is 53

> Note: Required sudo permission for listen port below 1024 
