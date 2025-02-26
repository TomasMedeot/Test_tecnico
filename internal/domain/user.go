package domain

type user struct {
	ID       interface{} `json:"id"`       // Unique identifier for the user
	Name     string      `json:"name"`     // Name of the user
	Email    string      `json:"email"`    // Email address of the user
	Password string      `json:"password"` // Hashed password for the user
}
