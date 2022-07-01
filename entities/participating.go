package entities

import "time"

type Participating struct {
	UserID string `json:"user_id"`
	GameID string `json:"game_id"`
	Nick   string `json:"nick"`
}

type ParticipatingByUser struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	FormattedDate   string    `json:"formatted_date"`
	StartDate       time.Time `json:"start_date"`
	EndDate         time.Time `json:"end_date"`
	EndGame         time.Time `json:"end_game"`
	State           int       `json:"state"`
	Photo           string    `json:"photo"`
	TypeGame        int       `json:"type_game"`
	PointsGG        int       `json:"points_gg"`
	ParticipatingID string    `json:"participating_id"`
	Nick            string    `json:"nick"`
}

//func (p ParticipatingByUser) ToDomain() ParticipatingByUser {
//	return ParticipatingByUser{
//		ID:              p.ID,
//		Name:            p.Name,
//		FormattedDate:   p.FormattedDate,
//		StartDate:       p.StartDate,
//		EndDate:         p.EndDate,
//		EndGame:         p.EndGame,
//		State:           p.State,
//		Photo:           p.Photo,
//		TypeGame:        p.TypeGame,
//		PointsGG:        p.PointsGG,
//		ParticipatingID: p.ParticipatingID,
//		Nick:            p.Nick,
//	}
//}
