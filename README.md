# go-redis-distributed-lock

Stackoverflow Discussion: https://softwareengineering.stackexchange.com/questions/442130/how-to-solve-duplicate-request-with-distributed-lock-management-in-golang/448124#448124

Library Used: https://github.com/go-redsync/redsync

## redsync
Distributed mutual exclusion lock using Redis for Go 

Redsync provides a Redis-based distributed mutual exclusion lock implementation for Go as described in this post. A reference library (by antirez) for Ruby is available at github.com/antirez/redlock-rb.

## Simulation

Command:
`go run main.go`

Output:
```
instance-2 is locked
instance-2 is unlocked
2023/11/07 12:10:26 instance-2 took 13.119042ms
instance-1 is locked
instance-1 is unlocked
2023/11/07 12:10:26 instance-1 took 70.338792ms
instance-3 is locked
instance-3 is unlocked
2023/11/07 12:10:26 instance-3 took 145.336ms
```