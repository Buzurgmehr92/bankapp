package types

//Money представляет собой денежную сумму в минимальных (центы, копейки, дирамы и т. д.)
type Money int64

//Currency представляет код валюты
type Currency string

//Коды валют
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

//PAN представляет информацию о платёжной карте
type PAN string

//Сфкв представляет информацию о платёжной карте
type Card struct {
	ID         int
	PAN        PAN
	Balance    Money
	Currency   Currency
	Color      string
	Name       string
	MinBalance Money
	Active     bool
}

//Payment представляет информацию о платеже
type Payment struct {
	ID     int
	Amount Money
}

type PaymentSource struct {
	Type string
	Number string
	Balance Money
}