package params

type VoteRequest struct {
	PostId int64 `json:"post_id" bindding:"required"`
	Vote   int64 `json:"vote" bindding:"required,oneof=1,0,-1"`
}
