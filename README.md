# Gache
A distributed cache written by golang

## Why we need cache?
Reduced the number of queries from database by storing data in local browser.

## A main process of using gache
```
                                     true
accept key --> check if it's cached ------> return value(1)
                | false                                   true
                |-------> should be got from remote node ------> contact with remote node --> return value(2)
                            | false
                            |------> call callback function, get value and add in cache --> return value(3)
```


## Some usually problem of cache.
#### If we don't have enough memory....
Use LRU(Least Recently Used) strategy to delete data.
#### Dirty read/write about concurrent control.
Use Mutual exclusion for some operation (Add, update, delete) to avoid dirty read/write.
#### Cache Avalanche & Hotspot Invalid
Internal Package "SingleFlight" can handle this problem.
## Uses of each part
[byteView.go](gache%2FbyteView.go): 
A readonly data structure to show the cache value.

[cache.go](gache%2Fcache.go): 
Use mutex lock to implement concurrent control in LRU cache.

[gache.go](gache%2Fgache.go):
Interactive with external part, solve the process of cache store and get.