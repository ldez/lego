package internal

import (
	"fmt"
	"time"
)

type APIResponse[T any] struct {
	Status Status `json:"status"`
	Data   T      `json:"data,omitempty"`
}

type Status struct {
	Code            int    `json:"code"`
	Title           string `json:"title"`
	Message         string `json:"message"`
	ID              string `json:"id"`
	TitleOriginal   string `json:"titleOriginal"`
	MessageOriginal string `json:"messageOriginal"`
	LogID           string `json:"logId"`
	Path            string `json:"path"`
	Time            int    `json:"time"`
	Version         string `json:"version"`
}

func (s Status) Error() string {
	return fmt.Sprintf("%d: %s %s (%s %d %s)", s.Code, s.Title, s.Message, s.ID, s.Time, s.Path)
}

type ZoneListResponse struct {
	Zone Zone `json:"zone"`
}

type RecordCreateResponse struct {
	ID     int64  `json:"id"`
	Record string `json:"record"`
}

type Zone struct {
	ID        string    `json:"id"`
	ProjectID string    `json:"projectId"`
	Title     string    `json:"title"`
	Zone      string    `json:"zone"`
	CreatedAt time.Time `json:"createdAt,omitzero"`
	UpdatedAt time.Time `json:"updatedAt,omitzero"`
}

type Record struct {
	ZoneID   string `json:"zoneId,omitempty"`
	Name     string `json:"name,omitempty"`
	Type     string `json:"type,omitempty"`
	Content  string `json:"content,omitempty"`
	TTL      int    `json:"ttl,omitempty"`
	Priority int    `json:"priority,omitempty"`
	Disabled bool   `json:"disabled,omitempty"`
}
