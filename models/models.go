package models

// Location : Lat, Long
type Location struct {
	Latitude  float32 `json:"lat"`
	Longitude float32 `json:"long"`
}

// User : Stores user info
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name" binding:"required"`
	GoogleID string `json:"GoogleID" binding:"required"`
	Score    int    `json:"score"`
}

// Challenge challenge
type Challenge struct {
	ID         int
	Icon       string
	IsActive   bool
	Location   Location
	PictureURL string
	Title      string
}

// CreateChallenge create challenge
type CreateChallenge struct {
	ChallengerID  int      `json:"challenger_id" binding:"required"`
	Title         string   `json:"title" binding:"required"`
	Location      Location `json:"location"`
	PictureURL    string   `json:"picture_url"`
	Icon          string   `json:"icon"`
	ChallengedIDs []int    `json:"challenged_ids"`
}

// ResponseStatus response status
type ResponseStatus int64

const (
	// Open : response status which has not been updated
	Open ResponseStatus = iota

	// Accepted : response which one the challenge
	Accepted

	// Declined : response that has beed declined
	Declined
)

func (e ResponseStatus) String() string {
	switch e {
	case Open:
		return "open"
	case Accepted:
		return "accepted"
	case Declined:
		return "declined"
	}
	// TODO: Might just want to throw error
	return ""
}

// ResponseStringToEnum Translate a response string to the status enum
func ResponseStringToEnum(str string) ResponseStatus {
	switch str {
	case "open":
		return Open
	case "accepted":
		return Accepted
	case "declined":
		return Declined
	}
	// TODO: Might just want to throw error
	return Open
}

// Response response
type Response struct {
	ID          int
	ChallengeID int `json:"challenge_id" binding:"required"`
	UserID      int `json:"user_id" binding:"required"`
	Status      ResponseStatus
	PictureURL  string `json:"picture_url" binding:"required"`
}

// UserChallenge userchallenge
type UserChallenge struct {
	ChallengeID  int
	ChallengerID int
	ChallengedID int
}

type ImageCreatedResponse struct {
	Location string `json:"location" binding:"required"`
}
