package altalabs

type AP struct {
	ID     string   `json:"id"`
	Chan2  int      `json:"chan2,omitempty"`
	Width2 int      `json:"width2,omitempty"`
	Chan5  int      `json:"chan5,omitempty"`
	Width5 int      `json:"width5,omitempty"`
	Txp2   string   `json:"txp2,omitempty"`
	Txp5   string   `json:"txp5,omitempty"`
	Colors []string `json:"colors,omitempty"`
}
