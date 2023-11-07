package main

import (
	"log"
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	goredislib "github.com/redis/go-redis/v9"
	"golang.org/x/sync/errgroup"
)

func main() {
	eg := errgroup.Group{}
	eg.Go(func() error {
		start := time.Now()
		printNumber("instance-1")

		elapsed := time.Since(start)
		log.Printf("instance-1 took %s", elapsed)

		return nil
	})
	eg.Go(func() error {
		start := time.Now()

		printNumber("instance-2")

		elapsed := time.Since(start)
		log.Printf("instance-2 took %s", elapsed)

		return nil
	})
	eg.Go(func() error {
		start := time.Now()

		printNumber("instance-3")

		elapsed := time.Since(start)
		log.Printf("instance-3 took %s", elapsed)

		return nil
	})
	eg.Wait()
}

func printNumber(instanceName string) {
	// Create a pool with go-redis (or redigo) which is the pool redisync will
	// use while communicating with Redis. This can also be any pool that
	// implements the `redis.Pool` interface.
	client := goredislib.NewClient(&goredislib.Options{
		Addr:       "localhost:6379",
		ClientName: instanceName,
	})
	pool := goredis.NewPool(client) // or, pool := redigo.NewPool(...)

	// Create an instance of redisync to be used to obtain a mutual exclusion
	// lock.
	rs := redsync.New(pool)

	// Obtain a new mutex by using the same name for all instances wanting the
	// same lock.
	mutexname := "my-global-mutex"
	mutex := rs.NewMutex(mutexname)

	// Obtain a lock for our given mutex. After this is successful, no one else
	// can obtain the same lock (the same mutex name) until we unlock it.
	if err := mutex.Lock(); err != nil {
		panic(err)
	}

	// Do your work that requires the lock.
	println(instanceName, "is locked")

	// Release the lock so other processes or threads can obtain a lock.
	if ok, err := mutex.Unlock(); !ok || err != nil {
		panic("unlock failed")
	}

	println(instanceName, "is unlocked")
}
