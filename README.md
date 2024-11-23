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
- [ ] [PostgreSQL: Documentation: 17: F.33. pg_trgm — support for similarity of text using trigram matching](https://www.postgresql.org/docs/current/pgtrgm.html)
- [ ] [Levenshtein module — Levenshtein 0.23.0 documentation](https://rapidfuzz.github.io/Levenshtein/levenshtein.html)

### Indexing
- [The ugrep file pattern searcher](https://ugrep.com/)
- [Genivia/ugrep-indexer: A monotonic indexer to speed up grepping by >10x (ugrep-indexer is now part of ugrep 6.0)](https://github.com/Genivia/ugrep-indexer)

### Tutorials
- [Understanding the Levenshtein Distance Equation for Beginners | by Ethan Nam | Medium](https://medium.com/@ethannam/understanding-the-levenshtein-distance-equation-for-beginners-c4285a5604f0)

### ChatGPT
```
Yes, there are several established techniques and algorithms to make Levenshtein distance calculations faster and more efficient. Here are some of the most common methods:

### 1. **Dynamic Programming Optimization**
   - **Space Reduction**: The traditional dynamic programming approach uses a 2D matrix, which can consume a lot of memory. Instead, you can optimize this to use only two one-dimensional arrays (or even one) to store the current and previous row of distances, reducing space complexity from O(m*n) to O(min(m, n)), where `m` and `n` are the lengths of the two strings.

### 2. **Early Exit**
   - Implement an early exit condition where if the cost exceeds a certain threshold during the computation, you can terminate the algorithm early. This is particularly useful in applications like fuzzy searching where you only care about distances up to a certain limit.

### 3. **Use of BK-Trees or Tries**
   - As discussed, BK-trees are effective for organizing strings in a way that allows for quick approximate matching. By structuring your data with a BK-tree or a Trie, you can significantly reduce the number of Levenshtein distance calculations needed.

### 4. **Approximate Matching with Hashing**
   - Use techniques like **MinHash** or **SimHash** to create hashed representations of strings, allowing for quick similarity assessments. This can help narrow down potential matches before performing the more expensive Levenshtein calculation.

### 5. **Parallelization**
   - If you need to compute distances for multiple pairs of strings, consider parallelizing the calculations. Using multi-threading or distributed computing can significantly speed up processing time in large datasets.

### 6. **Weighted Edit Distances**
   - Instead of a simple edit distance, you can implement a weighted version where different operations (insertions, deletions, substitutions) have different costs. This can help prioritize certain matches and optimize the search space.

### 7. **Memoization**
   - If you are frequently calculating distances between the same pairs of strings, store previously computed distances in a cache (memoization) to avoid redundant calculations.

### 8. **Use of Bitwise Operations**
   - For small alphabets, you can use bitwise operations to represent characters and compute distances more quickly. This is especially useful in certain constrained problems.

### 9. **Other Distance Metrics**
   - Depending on your application, consider using alternative string similarity measures that may be faster to compute, such as Jaro-Winkler or Hamming distance (for fixed-length strings). These can sometimes provide sufficiently accurate results with lower computational overhead.

### Conclusion

By combining these methods, you can significantly improve the efficiency of Levenshtein distance calculations, making them more suitable for real-time applications or large datasets. The choice of optimization techniques will depend on your specific use case, including the size of the data, the frequency of queries, and the importance of accuracy versus speed.
```
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
