# TCP Key-Value Store

A simple concurrent TCP-based key-value store written in Go.

## Features

* Server listens on a configurable TCP port.
* Supports commands:

  * `SET key value`
  * `GET key`
  * `DELETE key`
* `GET` on a missing key returns **NOT_FOUND**.
* Multiple clients supported concurrently (goroutine per connection).
* Thread-safe storage using a mutex-protected map.
* Simple text-based protocol (one command per line).
* Robust against client disconnects or malformed input.
* Optional TTL support:
  `SET key value 30` → key expires after 30 seconds.
* Background cleanup removes expired keys periodically.

## Example

Client request:

```
SET key1 hello
GET key1
DELETE key1
```

Server response:

```
SET_OK
GET_OK hello
DELETE_OK
```

## Run

Start server: naviagte to server folder

```

go run .
```
navigate to client folder
go run .

```
go run .
```

