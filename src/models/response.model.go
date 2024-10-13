package models

type Response struct {
	Model                string  `json:"model"`
	Created_at           string  `json:"created_at"`
	Response             string  `json:"response"`
	Done                 bool    `json:"done"`
	Total_Duration       int64   `json:"total_duration"`
	Load_Duration        int64   `json:"load_duration"`
	Prompt_Eval_Count    int64   `json:"prompt_eval_count"`
	Prompt_Eval_Duration int64   `json:"prompt_eval_duration"`
	Eval_Count           int64   `json:"eval_count"`
	Eval_Duration        int64   `json:"eval_duration"`
	Context              []int64 `json:"context"`
}
