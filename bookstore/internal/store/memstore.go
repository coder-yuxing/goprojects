package store

import (
	mystore "github.com/yuxing/hellomodule/bookstore/store"
	"github.com/yuxing/hellomodule/bookstore/store/factory"
	"sync"
)

func init()  {
	factory.Register("mem", &MemStore{
		books: make(map[string]*mystore.Book),
	})
}

// MemStore 内存存储
// 基于内存实现 store.Book 的存取操作
type MemStore struct {
	sync.RWMutex
	books map[string]*mystore.Book
}

func (m *MemStore) Create(b *mystore.Book) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.books[b.Id]; ok {
		return mystore.ErrExist
	}
	book := *b
	m.books[b.Id] = &book

	return nil
}

func (m *MemStore) Update(b *mystore.Book) error {
	m.Lock()
	defer m.Unlock()

	oldBook, ok := m.books[b.Id]
	if !ok {
		return mystore.ErrNotFound
	}
	book := *oldBook
	if b.Name != "" {
		book.Name = b.Name
	}
	if b.Authors != nil {
		book.Authors = b.Authors
	}

	if b.Press != "" {
		book.Press = b.Press
	}

	m.books[book.Id] = &book

	return nil
}


func (m *MemStore) Get(id string) (mystore.Book, error) {
	m.RLock()
	defer m.RUnlock()

	t, ok := m.books[id]
	if ok {
		return *t, nil
	}
	return mystore.Book{}, mystore.ErrNotFound
}

func (m *MemStore) GetAll() ([]mystore.Book, error) {
	m.RLock()
	defer m.RUnlock()

	allBooks := make([]mystore.Book, 0, len(m.books))
	for _, book := range m.books {
		allBooks = append(allBooks, *book)
	}
	return allBooks, nil
}

func (m *MemStore) Delete(id string) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.books[id]; !ok {
		return mystore.ErrNotFound
	}

	delete(m.books, id)
	return nil
}