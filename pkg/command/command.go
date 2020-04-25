package command

import "fmt"

type Order interface {
	Exec(int) string
}

/************************************************/
type Stock struct{}

func (s *Stock) Buy(x int) string {
	return fmt.Sprintf("buy stock: %d", x)
}
func (s *Stock) Sell(x int) string {
	return fmt.Sprintf("sell stock: %d", x)
}

/************************************************/
type Buy struct{
	stock Stock
}

func (b Buy) Exec(x int) string {
	return b.stock.Buy(x)
}

/************************************************/
type Sell struct{
	stock Stock
}

func (s Sell) Exec(x int) string {
	return s.stock.Sell(x)
}

/************************************************/
type Broker struct{}

func (c Broker) ExecOrder(o Order, x int) string {
	return o.Exec(x)
}
