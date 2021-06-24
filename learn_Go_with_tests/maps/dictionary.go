package maps

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

// var (
// 	ErrNotFound   = errors.New("could not find the word you were looking for")
// 	ErrWordExists = errors.New("cannot add word because it already exists")
// )

func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}

// 此处遇到指针无法寻找到索引的问题
//可以参考此处： https://stackoverflow.com/questions/36463608/go-invalid-operation-type-mapkeyvalue-does-not-support-indexing
// question: 基本的搜索很容易实现，但是如果我们提供一个不在我们字典中的单词，会发生什么呢？
func (d *Dictionary) Search(word string) (string, error) {

	if _, ok := (*d)[word]; !ok {
		return "", ErrNotFound
	}

	return (*d)[word], nil // 以括号括起，表示获取到 map 去寻址 而不是指针
}

/*
Map 有一个有趣的特性，不使用指针传递你就可以修改它们。

  - 这是因为 map 是引用类型。这意味着它拥有对底层数据结构的引用，就像指针一样。
  - 它底层的数据结构是 hash table 或 hash map，你可以在这里阅读有关 hash tables (https://en.wikipedia.org/wiki/Hash_table) 的更多信息。

Map 作为引用类型是非常好的，因为无论 map 有多大，都只会有一个副本。

引用类型引入了 maps 可以是 nil 值。如果你尝试使用一个 nil 的 map，你会得到一个 nil 指针异常，这将导致程序终止运行。

由于 nil 指针异常，你永远不应该初始化一个空的 map 变量：

   - 错误示范：  var m map[string]string

相反，你可以像我们上面那样初始化空 map，或使用 make 关键字创建 map：

   - 正确示范:   m = make(map[string]string)       or        m = map[string]string{}

这两种方法都可以创建一个空的 hash map 并指向 dictionary。这确保永远不会获得 nil 指针异常。
*/

/*
重构后

这里我们使用 switch 语句来匹配错误。如上使用 switch 提供了额外的安全，以防 Search 返回错误而不是 ErrNotFound。
*/
func (d *Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		(*d)[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d *Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil: // Word is Exist
		(*d)[word] = definition
	default:
		return err
	}
	return nil
}

/*
Go 的 map 有一个内置函数 delete。

它需要两个参数。第一个是这个 map，第二个是要删除的键。

delete 函数不返回任何内容，我们基于相同的概念构建 Delete 方法。

由于删除一个不存在的值是没有影响的，与我们的 Update 和 Create 方法不同，我们不需要用错误复杂化 API。
*/
func (d *Dictionary) Delete(word string) {
	delete((*d), word)
}

type Dictionary map[string]string

/*

我们将错误声明为常量，这需要我们创建自己的 DictionaryErr 类型来实现 error 接口。

你可以在 Dave Cheney 的这篇优秀文章中了解更多相关的细节。(https://dave.cheney.net/2016/04/07/constant-errors)

简而言之，它使错误更具可重用性和不可变性。
*/
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}
