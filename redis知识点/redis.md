# Redis面试考点

## 一、基础

redis基本概念：
    Redis (Remote Dictionary Server) 是一个使用 C 语言 编写的，开源的 (BSD许可) 高性能 非关系型 (NoSQL) 的 键值对数据库。

Redis 可以存储 键 和 不同类型数据结构值 之间的映射关系。键的类型只能是字符串，而值除了支持最 基础的五种数据类型 外，还支持一些 高级数据类型：

- 基本数据类型：
  - string字符串(支持的方法)：SET/GET/GETSET/SETEX/SETNX

  - list列表(支持的方法)：EPUSH/LPUSH/RPOP/LPOP

  - hash字典(支持的方法)：HSET/HGET/HGETTALL

  - set集合(支持的方法)：SADD/SMEBERS/SCARD/SPOP

  - zset有序列表(支持的方法)：ZADD/ZRANGE/ZCARD/ZSCORE/ZRANK/ZREM

- 高级数据类型：
  - bitMap位图

  - Hyperloglog

  - 布隆过滤器

  - GeoHash

  - Pub/Sub

  - Stream


### Redis小结：
    与传统数据库不同的是Redis的数据是存在内存中，所以读写速度非常快，因此 Redis 被广泛应用于 缓存 方向，每秒可以处理超过 10 万次读写操作，是已知性能最快的 Key-Value 数据库。另外，Redis 也经常用来做 分布式锁。

除此之外，Redis 支持事务 、持久化、LUA脚本、LRU驱动事件、多种集群方案。

## Redis优缺点

### 优点

- 读写性能优异， Redis能读的速度是 110000 次/s，写的速度是 81000 次/s。

- 支持数据持久化，支持 AOF 和 RDB 两种持久化方式。

- 支持事务，Redis 的所有操作都是原子性的，同时 Redis 还支持对几个操作合并后的原子性执行。

- 数据结构丰富，除了支持 string 类型的 value 外还支持 hash、set、zset、list 等数据结构。

- 支持主从复制，主机会自动将数据同步到从机，可以进行读写分离。

### 缺点

- 数据库 容量受到物理内存的限制，不能用作海量数据的高性能读写，因此 Redis 适合的场景主要局限在较小数据量的高性能操作和运算上。

- Redis 不具备自动容错和恢复功能，主机从机的宕机都会导致前端部分读写请求失败，需要等待机器重启或者手动切换前端的 IP 才能恢复。

- 主机宕机，宕机前有部分数据未能及时同步到从机，切换 IP 后还会引入数据不一致的问题，降低了 系统的可用性。

- Redis 较难支持在线扩容，在集群容量达到上限时在线扩容会变得很复杂。为避免这一问题，运维人员在系统上线时必须确保有足够的空间，这对资源造成了很大的浪费。