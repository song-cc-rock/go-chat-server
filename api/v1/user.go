package v1

type AddUserResponse struct {
	ID       string `json:"id"`                                // ID
	Name     string `json:"name"`                              // 账户名
	NickName string `json:"nickName"`                          // 用户名
	Phone    string `json:"phone"`                             // 电话
	Mail     string `json:"mail"`                              // 邮箱
	Avatar   string `gorm:"type:varchar(1000);" json:"avatar"` // 头像
	Status   string `json:"status"`                            // 状态	未申请 => not-applied , 申请待通过 => pending, 申请已拒绝 => rejected, 通过 => approved
}
