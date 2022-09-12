package domain

type Task struct {
	ID      int    `json:"id"`
	Nombre  string `json:"nombre"`
	Content string `json:"content"`
}

type ResponseInfo struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}
