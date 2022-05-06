package exchange

type ExchangeClient interface {
	GetOrderBook(marketSymbol string) (*OrderBook, error)
	GetBalanceOfCoin(coinSymbol string) (float64, error)
	NewLimitOrder(marketSymbol, orderType string, amount, price float64) error
	CancelOpenOrder(OrderID string) error
	CancelAllOpenOrders() error
}
