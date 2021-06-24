package maps

import "testing"

/*
声明 map 的方式有点儿类似于数组。

不同之处是，它以 map 关键字开头，需要两种类型。

第一个是键的类型，写在 [] 中。第二个是值的类型，跟在 [] 之后。


键的类型很特别，它只能是一个可比较的类型，因为如果不能判断两个键是否相等，我们就无法确保我们得到的是正确的值。

可比类型在语言规范中有详细解释。(https://golang.org/ref/spec#Comparison_operators)

另一方面，值的类型可以是任意类型，它甚至可以是另一个 map。
*/

func TestSearch1(t *testing.T) {
	dictionary := map[string]string{"test": "this is just a test"}

	got := Search(dictionary, "test")
	want := "this is just a test"

	assertString(t, got, want)
}

func TestSearch2(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	got, err := dictionary.Search("test")
	if err != nil {
		assertString2(t, err, ErrNotFound)
	}

	want := "this is just a test"

	assertString(t, got, want)
}

/*
Error 类型可以使用 .Error() 方法转换为字符串，我们将其传递给断言时会执行此操作。

我们也用 if 来保护 assertStrings，以确保我们不在 nil 上调用 .Error()。
*/
func TestSearch3(t *testing.T) {
	dictionary := &Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertString(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertString2(t, err, ErrNotFound)
	})
}

/*
我们的 Add 看起来不错。除此之外，我们没有考虑当我们尝试添加的值已经存在时会发生什么！

如果值已存在，map 不会抛出错误。相反，它们将继续并使用新提供的值覆盖该值。

这在实践中很方便，但会导致我们的函数名称不准确。Add 不应修改现有值。它应该只在我们的字典中添加新单词。
*/
func TestAdd(t *testing.T) {
	dictionary := &Dictionary{}
	word := "test"
	definition := "this is just a test"

	dictionary.Add(word, definition)

	assertDefinition(t, dictionary, word, definition)
}

func TestAdd2(t *testing.T) {
	t.Run("new world", func(t *testing.T) {
		dictionary := &Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := &Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	word := "test"
	definition := "this is just a test"
	dictionary := &Dictionary{word: definition}
	newDefinition := "new definition"

	dictionary.Update(word, newDefinition)

	assertDefinition(t, dictionary, word, newDefinition)
}

func TestUpdate2(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		newDefinition := "new definition"
		dictionary := &Dictionary{word: definition}

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := &Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := &Dictionary{word: "test definition"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("Expected '%s' to be deleted", word)
	}
}

func assertString(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func assertString2(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error '%s' want '%s'", got, want)
	}
}

func assertDefinition(t *testing.T, dictionary *Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word: ", err)
	}

	if definition != got {
		t.Errorf("got '%s' want '%s'", got, definition)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error '%s' want '%s'", got, want)
	}
}

/*
总结

在本节中，我们介绍了很多内容。我们为一个字典应用创建了完整的 CRUD API。在整个过程中，我们学会了如何：

  - 创建 map
  - 在 map 中搜索值
  - 向 map 添加新值
  - 更新 map 中的值
  - 从 map 中删除值
  - 了解更多错误相关的知识
  - 如何创建常量类型的错误
  - 对错误进行封装

TODO: 优化 Update 与 Delete
*/
