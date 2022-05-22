package param

// 投稿
type Video struct {
	Title string `form:"title"  binding:"required" msg:"标题为空标题"`
}

// 查询用户稿件列表
type VideoList struct {
	UserID int `form:"user_id" binding:"required" msg:"无效的用户标识"`
}
