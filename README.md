levenshtein-search-lib
======================
### TODOs
- [ ] offload to async background IO, e.g. **smaller** binary indexed files stored in system `$TMP` (to eliminate network to database)
- [ ] [**BK-tree - Wikipedia**](https://en.wikipedia.org/wiki/BK-tree)
- [ ] monitor backend server and do indexing when it is low usage
- [ ] should include throttling (search text box may send text fast)
- [ ] skip searching
  - >= 70% match = if search key is shorter than the current data:
    - `1.0 - (levenshtein.ComputeDistance(searchString, data) / len(data) )`
      - the bracket should gives less then 0.3 -> let say "peter" for data, len is 5, 0.3 * 5 = 1.5, so the comparing should skip once it is larger than 1.5 


### Cache libraries
- [patrickmn/go-cache: An in-memory key:value store/cache (similar to Memcached) library for Go, suitable for single-machine applications.](https://github.com/patrickmn/go-cache)
- [go-cache](https://patrickmn.com/projects/go-cache/)
- [eko/gocache: ☔️ A complete Go cache library that brings you multiple ways of managing your caches](https://github.com/eko/gocache)
- [**allegro/bigcache: Efficient cache for gigabytes of data written in Go.**](https://github.com/allegro/bigcache)
  - [Writing a very fast cache service with millions of entries in Go | blog.allegro.tech](https://blog.allegro.tech/2016/03/writing-fast-cache-service-in-go.html)
- [dgraph-io/ristretto: A high performance memory-bound Go cache](https://github.com/dgraph-io/ristretto)
  - [Introducing Ristretto: A High-Performance Go Cache - Dgraph Blog](https://dgraph.io/blog/post/introducing-ristretto-high-perf-go-cache/)
- [Go缓存系列之: GoCache使用与分析-CSDN博客](https://blog.csdn.net/baidu_32452525/article/details/118199304)
