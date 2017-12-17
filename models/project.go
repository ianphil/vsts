package models

// Project represents json from api
type Project struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	URL        string `json:"url"`
	State      string `json:"state"`
	Revision   int    `json:"revision"`
	Visibility string `json:"visibility"`
}

// ProjectList represents res.body from api call
type ProjectList struct {
	Count int       `json:"count"`
	Value []Project `json:"value"`
}
