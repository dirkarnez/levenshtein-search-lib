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

### BK Tree
```python
class BKTreeNode:
    def __init__(self, word):
        self.word = word
        self.children = {}

def levenshtein_distance(s1, s2):
    # Implementation of Levenshtein distance calculation
    if len(s1) < len(s2):
        return levenshtein_distance(s2, s1)

    if len(s2) == 0:
        return len(s1)

    previous_row = range(len(s2) + 1)
    for i, c1 in enumerate(s1):
        current_row = [i + 1]
        for j, c2 in enumerate(s2):
            insertions = previous_row[j + 1] + 1
            deletions = current_row[j] + 1
            substitutions = previous_row[j] + (c1 != c2)
            current_row.append(min(insertions, deletions, substitutions))
        previous_row = current_row
    return previous_row[-1]

class BKTree:
    def __init__(self):
        self.root = None

    def insert(self, word):
        if not self.root:
            self.root = BKTreeNode(word)
        else:
            self._insert_rec(self.root, word)

    def _insert_rec(self, node, word):
        distance = levenshtein_distance(node.word, word)
        if distance in node.children:
            self._insert_rec(node.children[distance], word)
        else:
            node.children[distance] = BKTreeNode(word)

    def search(self, word, threshold):
        results = []
        self._search_rec(self.root, word, threshold, results)
        return results

    def _search_rec(self, node, word, threshold, results):
        if node is None:
            return
        distance = levenshtein_distance(node.word, word)
        if distance <= threshold:
            results.append(node.word)
        for d in range(max(0, distance - threshold), distance + threshold + 1):
            if d in node.children:
                self._search_rec(node.children[d], word, threshold, results)

# Example usage
bk_tree = BKTree()
words = ["about", "above", "bought", "brought"]
for word in words:
    bk_tree.insert(word)

matches = bk_tree.search("uo", 2)  # Searching for words within a Levenshtein distance of 2
print(matches)
```

### Cache libraries
- [patrickmn/go-cache: An in-memory key:value store/cache (similar to Memcached) library for Go, suitable for single-machine applications.](https://github.com/patrickmn/go-cache)
- [go-cache](https://patrickmn.com/projects/go-cache/)
- [eko/gocache: ☔️ A complete Go cache library that brings you multiple ways of managing your caches](https://github.com/eko/gocache)
- [**allegro/bigcache: Efficient cache for gigabytes of data written in Go.**](https://github.com/allegro/bigcache)
  - [Writing a very fast cache service with millions of entries in Go | blog.allegro.tech](https://blog.allegro.tech/2016/03/writing-fast-cache-service-in-go.html)
- [dgraph-io/ristretto: A high performance memory-bound Go cache](https://github.com/dgraph-io/ristretto)
  - [Introducing Ristretto: A High-Performance Go Cache - Dgraph Blog](https://dgraph.io/blog/post/introducing-ristretto-high-perf-go-cache/)
- [Go缓存系列之: GoCache使用与分析-CSDN博客](https://blog.csdn.net/baidu_32452525/article/details/118199304)
