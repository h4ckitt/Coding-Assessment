package presenter

type Car struct {
	Name       string   `json:"name"`
	Type       string   `json:"type"`
	Color      string   `json:"color"`
	SpeedRange int      `json:"speed_range"`
	Features   []string `json:"features"`
}
