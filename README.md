# Gache
A distributed cache written by golang

## Why we need Cache?
Reduced the number of queries from database by storing data in local browser.

## Some usually problem of Cache.
#### If we don't have enough memory....
Use LRU(Least Recently Used) strategy to delete data.
#### Dirty read/write about concurrent control.
Use Mutex lock for some operation (Add, update, delete) to avoid dirty read/write.
