package params

type VoteRequest struct {
	PostId string  `json:"post_id" bindding:"required"`
	Vote   float64 `json:"vote" bindding:"required,oneof=1,0,-1"`
}
