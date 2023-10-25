package configs

type User struct {
	Id       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty" validate:"required"`
	Location string `json:"location,omitempty" validate:"required"`
	Title    string `json:"title,omitempty" validate:"required"`
}
