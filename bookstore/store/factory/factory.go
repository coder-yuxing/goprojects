package factory

import (
	"fmt"
	"github.com/yuxing/goprojects/bookstore/store"
	"sync"
)

var (
	providersMu sync.RWMutex
	providers = make(map[string]store.Store)
)

// Register 注册 store.Store 实现
func Register(name string, p store.Store) {
	providersMu.Lock()
	defer providersMu.Unlock()

	if p == nil {
		panic("store: Register provider is nil")
	}

	if _, dup := providers[name]; dup {
		panic("store: Register called twice for provider " + name)
	}

	providers[name] = p
}

func New(providerName string) (store.Store, error) {
	providersMu.RLock()
	p, ok := providers[providerName]
	providersMu.RUnlock()

	if !ok {
		return nil, fmt.Errorf("store: unknown provider %s", providerName)
	}

	return p, nil
}