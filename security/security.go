package security

import (
	"exchange/order"
)

type Security struct {
	name        string
	commChannel chan order.Message
}
