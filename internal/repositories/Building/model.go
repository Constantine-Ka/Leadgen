package Building

type Building struct {
	ID    int64  `json:"id,omitempty" db:"id" form:"id,omitempty"`
	Title string `db:"name" json:"name,omitempty" form:"name,omitempty"`
	City  string `db:"city" json:"city,omitempty" form:"city,omitempty"`
	Year  int    `db:"year" json:"year,omitempty" form:"year,omitempty"`
	Level int    `db:"level" json:"level,omitempty" form:"level,omitempty"`
}
