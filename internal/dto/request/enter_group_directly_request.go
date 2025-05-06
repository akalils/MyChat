package request

type EnterGroupDirectlyRequest struct {
	OwnerId   string `json:"owner_id"` // 群聊id
	ContactId string `json:"contact_id"`
}
