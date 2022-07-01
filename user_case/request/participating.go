package request

type ParticipatingReq struct {
	UserID string `json:"user_id"`
	GameID string `json:"game_id"`
	Nick   string `json:"nick"`
}
