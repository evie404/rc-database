# rc-database

## Prompt

Write a program that runs a server that is accessible on http://localhost:4000/. When your server receives a request on http://localhost:4000/set?somekey=somevalue it should store the passed key and value in memory. When it receives a request on http://localhost:4000/get?key=somekey it should return the value stored at `somekey`.

## Features

- Database, server that supports Get/Set
- Client that interfaces with the server
- Supports concurrent access

## TODO

- Disk persistence
