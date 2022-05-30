package presenter

import "time"

type Project struct {
	ID          uint64    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartAt     time.Time `json:"start_at"`
	FinishAt    time.Time `json:"finish_at"`
	ConcludedAt time.Time `json:"concluded_at"`
	Status      string    `json:"status"`
}
