package Users

type User struct {
	Id              int    `json:"id" gorm:"primary_key"`
	UserName        string `json:"user_name" gorm:"comment:'姓名' " `
	UserSex         int    `json:"user_sex" gorm:"comment:'性别 0男 1女' " `
	UserTel         string `json:"user_tel" gorm:"type:char(20);not null; default ''; comment:'电话' " `
	ReceivableMoney string `json:"receivable_money" gorm:"comment:'应收金额' " `
	ReceivedMoney   string `json:"received_money" gorm:"comment:'实收金额' " `
}
