package pokemon


type Pokemon struct {
	Id int               `json:"id"`
	Name string          `json:"name"`
	Type string          `json:"type"`
	Description string   `json:"description"`
}
