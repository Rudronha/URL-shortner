package cache

import (
    "github.com/hashicorp/golang-lru/v2"
    "log"
)

var lruCache *lru.Cache[string, string]

func InitLRUCache(size int) {
    var err error
    lruCache, err = lru.New[string, string](size)
    if err != nil {
        log.Fatalf("Failed to create LRU cache: %v", err)
    }
}

func GetMemory(key string) (string, bool) {
    return lruCache.Get(key)
}

func SetMemory(key, value string) {
    lruCache.Add(key, value)
}
