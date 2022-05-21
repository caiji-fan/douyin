package param

type Video struct {
	Title string `form:"title"  binding:"required" msg:"标题为空标题"`
}
