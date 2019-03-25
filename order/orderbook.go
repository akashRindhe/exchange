package order

/*OrderBook - Represents the set of orders at each bid and ask price level */
type orderBook struct {
	security    string
	commChannel chan Message
}

func InitalizeOrderBook(security string, commChannel chan Message) {

}

/*AddToOrderBook - Adds the order to the order book*/
func (ob *orderBook) AddToOrderBook() {
	for {

	}
}

func (ob *orderBook) MatchOrders() {

}
