/*
Package redis implement a pure redis cluster client, meaning it doesn't
support any cluster commands.

Create a new cluster client with specified options:

    cluster, err := redis.NewCluster(
        &redis.Options{
    	StartNodes: []string{"127.0.0.1:7000", "127.0.0.1:7001", "127.0.0.1:7002"},
    	ConnTimeout: 50 * time.Millisecond,
    	ReadTimeout: 50 * time.Millisecond,
    	WriteTimeout: 50 * time.Millisecond,
    	KeepAlive: 16,
    	AliveTime: 60 * time.Second,
    })

For basic usage:

    cluster.Do("SET", "foo", "bar")
    cluster.Do("INCR", "mycount", 1)
    cluster.Do("LPUSH", "mylist", "foo", "bar")
    cluster.Do("HMSET", "myhash", "f1", "foo", "f2", "bar")

Use convert help functions to convert replies to int, float, string, etc:

    reply, err := Int(cluster.Do("INCR", "mycount", 1))
    reply, err := String(cluster.Do("GET", "foo"))
    reply, err := Strings(cluster.Do("LRANGE", "mylist", 0, -1))
    reply, err := StringMap(cluster.Do("HGETALL", "myhash"))

Use batch interface to pack multiple commands for pipelining:

    batch := cluster.NewBatch()
    batch.Put("LPUSH", "country_list", "France")
    batch.Put("LPUSH", "country_list", "Italy")
    batch.Put("LPUSH", "country_list", "Germany")
    batch.Put("INCRBY", "countries", 3)
    batch.Put("LRANGE", "country_list", 0, -1)
*/
package redis
