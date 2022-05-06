package exchange

import (
	"fmt"
	"strconv"
)

type Order struct {
	Quantity float64
	Price    float64
}

type OrderBook struct {
	Bid []Order
	Ask []Order
}

func (book *OrderBook) GetAmountBuy(balanceToBuy float64) (amount float64) {
	balanceremaining := balanceToBuy
	for _, x := range book.Ask {
		if x.Price <= 0 {
			return 0
		}
		if balanceremaining < x.Quantity*x.Price {
			amount += balanceremaining / x.Price
			return
		}
		amount += x.Quantity
		balanceremaining -= x.Quantity * x.Price
	}
	return
}

func (book *OrderBook) GetAmountSell(amountToSale float64) (balance float64) {
	balanceremaining := amountToSale
	for _, x := range book.Bid {
		if x.Price <= 0 {
			return 0
		}
		if balanceremaining < x.Quantity {
			balance += balanceremaining * x.Price
			return
		}
		balanceremaining -= x.Quantity
		balance += x.Quantity * x.Price
	}
	return
}

func (book *OrderBook) String(count, decimalsQ, decimalsR int) string {
	result := "\nSells:\n"
	result += addst(reverse(book.Ask[:min(len(book.Ask), count)]), decimalsQ, decimalsR)
	result += "\nBuys:\n"
	result += addst(book.Bid[:min(len(book.Bid), count)], decimalsQ, decimalsR)
	return result
}

func reverse(ords []Order) []Order {
	od := make([]Order, len(ords))
	for i, x := range ords {
		od[len(od)-i-1] = x
	}
	return od
}

func addst(orders []Order, decimalsQ, decimalsR int) (result string) {
	for i := 0; i < len(orders); i++ {
		quantSt := strconv.FormatFloat(orders[i].Quantity, 'f', decimalsQ, 64)
		rate := strconv.FormatFloat(orders[i].Price, 'f', decimalsR, 64)
		result += fmt.Sprintf("%s => %s\n", quantSt, rate)
	}
	return
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
