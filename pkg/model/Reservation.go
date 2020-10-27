package model

type Reservation struct {
	ID        int    `json:"id" id:"auto"`
	Something string `json:"something"`
}

func (r Reservation) Table() string {
	return "reservation"
}

func (r Reservation) Validate() (string, bool) {
	if r.Something == "" {
		return "Sommething is required", false
	}
	return "", true
}
