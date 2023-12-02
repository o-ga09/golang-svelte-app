package domain

import "time"

type Note struct {
	UserId     UserId `json:"user_id,omitempty"`
	NoteId     NoteId `json:"note_id,omitempty"`
	Contents   Contents `json:"contents,omitempty"`
	Image_url  Image_url `json:"image_url,omitempty"`
	Created_at Created_at `json:"created_at,omitempty"`
	Updated_at Updated_at `json:"updated_at,omitempty"`
	Deleted_at Deleted_at `json:"deleted_at,omitempty"`
}

type User struct {
	UserId     UserId `json:"user_id,omitempty"`
	Created_at Created_at `json:"created_at,omitempty"`
	Updated_at Updated_at `json:"updated_at,omitempty"`
	Deleted_at Deleted_at `json:"deleted_at,omitempty"`
}

type RequestJson struct {
	UserId     UserId `json:"user_id,omitempty"`
	Contents   Contents `json:"contents,omitempty"`
	Image_url  Image_url `json:"image_url,omitempty"`
}

type Notes []Note
type UserId struct{ Value string}
type NoteId struct{ Value string}
type Contents struct{ Value string}
type Image_url struct{ Value string}
type Created_at struct{ Value time.Time}
type Updated_at struct{ Value time.Time}
type Deleted_at struct{ Value time.Time}
