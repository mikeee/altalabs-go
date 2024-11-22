package altalabs

type Site struct {
	ID       string        `json:"id"`
	Name     string        `json:"name"`
	Icon     string        `json:"icon"`
	Devices  []interface{} `json:"devices"` // TODO: implement devices struct
	Online   int           `json:"online"`
	Emails   []string      `json:"emails"`
	Identity struct {
		Email             string `json:"email"`
		Admin             bool   `json:"admin"`
		AllPasswords      bool   `json:"allPasswords"`
		UnlockedPasswords bool   `json:"unlockedPasswords"`
	} `json:"perms"`
}

type Sites map[string]Site

// TODO: unmarshal sites into a map by string rather than ID
