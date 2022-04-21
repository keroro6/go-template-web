package types

type Request struct {
	Name string `path:"name,options=you|me"`
}

type Response struct {
	Message string `json:"message"`
	ID      int64  `json:"id"`
	Name    string `json:"name"`
}
