package model

type (
	User struct {
		Base
		HasPassword
		HasTimestamp
		Username string `json:"username"`
	}
)
