package resp

type Order struct {
	Id          string `json:"id"`
	Key         string `json:"key"`
	OrderId     string `json:"orderId"`
	UserId      string `json:"userId"`
	NickName    string `json:"nickName"`
	Mobile      string `json:"mobile"`
	TotalPrice  int64  `json:"totalPrice"`
	PayStatus   int    `json:"payStatus"`
	PayType     int    `json:"payType"`
	PayTime     string `json:"payTime"`
	OrderStatus int    `json:"orderStatus"`
	ExtraInfo   string `json:"extraInfo"`
	UserAddress string `json:"userAddress"`
	IsDeleted   bool   `json:"isDeleted"`
}
