package pointeur_and_error

import (
	"errors"
	"fmt"
)

/*
*我们在测试代码和 Withdraw 代码中都有重复的错误消息。
*
*如果有人想要重新定义这个错误，那么测试就会失败，这将是非常恼人的，而对于我们的测试来说，这里有太多的细节了。
*
* 我们并不关心具体的措辞是什么，只是在给定条件的情况下返回一些有意义的错误。
*
*在 Go 中，错误是值，因此我们可以将其重构为一个变量，并为其提供一个单一的事实来源。
 */

var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

// 买入操作
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// 获取当前虚拟币的值
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// 提取操作
func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return InsufficientFundsError
	}

	w.balance -= amount
	return nil
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
