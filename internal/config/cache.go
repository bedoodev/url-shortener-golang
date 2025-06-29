package config

import (
	"time"

	"github.com/bedoodev/high-level-url-shortener/internal/cache"
)

var HotKeyCache *cache.HotKeyCache
var StopCleanupChan chan struct{}

func InitCache() {
	// local cache for 10 seconds
	HotKeyCache = cache.NewHotKeyCache(10 * time.Second)

	StopCleanupChan = make(chan struct{})
	HotKeyCache.StartCleanup(30*time.Second, StopCleanupChan)
}
