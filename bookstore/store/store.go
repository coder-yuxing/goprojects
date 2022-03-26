package store

import "errors"

var (
	ErrNotFound = errors.New("not found")
	ErrExist    = errors.New("exist")
)

// Book 定义图书抽象数据类型
type Book struct {
	Id      string   `json:"id"`      // 图书ISBN ID
	Name    string   `json:"name"`    // 图书名称
	Authors []string `json:"authors"` // 图书作者
	Press   string   `json:"press"`   // 出版社
}

// Store 定义针对 Book 存储的接口类型
// 对于想要进行图书数据操作的一方而言，
// 他只需要得到一个满足 Store 接口的实例就可以实现对图书数据的存取操作
// 而无需关心图书数据采取何种存储方式
// 这就实现了图书存储操作与底层图书数据存储方式的解耦
// 这种面相接口编程也是 Go 组合设计哲学的一个重要体现
type Store interface {
	Create(*Book) error       // 创建一个新图书条目
	Update(*Book) error       // 更新某图书条目
	Get(string) (Book, error) // 获取某图书信息
	GetAll() ([]Book, error)  // 获取所有图书信息
	Delete(string) error      // 删除某图书条目
}
