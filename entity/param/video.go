package param

type Video struct {
	AuthorId int    `form:"author_id" binding:"required" msg:"无效的用户标识"`
	Title    string `form:"title"  binding:"required" msg:"标题为空标题"`
}
