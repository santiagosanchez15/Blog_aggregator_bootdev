package config

type Config struct {
	DBURL           string `json:"db_url"`            // tags work with json to mathc the exact values
	CurrentUserName string `json:"current_user_name"` // add tag
}
