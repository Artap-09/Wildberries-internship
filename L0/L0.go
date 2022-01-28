package main

func main() {

}

type data struct {
	OrderUid    string `json: "order_uid"`
	TrackNumber string `json: "track_number"`
	Entry       string `json: "entry"`
	Delivery    struct {
		Name    string `json: "name"`
		Phone   string `json: "phone"`
		Zip     string `json: "zip"`
		City    string `json: "city"`
		Address string `json: "address"`
		Region  string `json: "region"`
		Email   string `json: "email"`
	} `json: "delivery"`
	Payment struct {
		Transaction  string `json: "transaction"`
		RequestId    string `json: "request_id"`
		Currency     string `json: "currency"`
		Provider     string `json: "provider"`
		Amount       int    `json: "amount"`
		PaymentDt    int    `json: "payment_dt"`
		Bank         string `json: "bank"`
		DeliveryCost int    `json: "delivery_cost"`
		GoodsTotal   int    `json: "goods_total"`
		CustomFee    int    `json: "custom_fee"`
	} `json: "payment"`
	Items []struct {
		ChrtId      int    `json: "chrt_id"`
		TrackNumber string `json: "track_number'`
		Price       int    `json: "price"`
		Rid         string `json: "rid"`
		Name        string `json: "name"`
		Sale        int    `json: "sale"`
		Size        string `json: "size"`
		Total_price int    `json: "total_price"`
		Nm_id       int    `json: "nm_id"`
		Brand       string `json: "brand"`
		Status      int    `json: "status"`
	}
}
