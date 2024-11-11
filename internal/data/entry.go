package data

type Entry struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

type EntryPayload struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}
