package request

type UsersRequest struct {
	ID uint64 `form:"userId" binding:"required,gte=1"`
	//State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UsersListRequest struct {
	Page int `form:"page" binding:"gte=1"`
	//State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}
