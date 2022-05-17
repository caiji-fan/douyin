/**
 * @Author yg
 * @Date 2022-05-17
 * @Description
 **/
package param

//Relation 关注参数
type Relation struct {
	UserID     int  `json:"user_id" form:"user_id"`
	ToUserID   int  `json:"to_user_id" form:"to_user_id"`
	ActionType byte `json:"action_type" form:"action_type"  binding:"required" msg:"无效的操作类型"`
}
