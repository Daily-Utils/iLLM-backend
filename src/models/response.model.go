package models

type Response struct {
	model                string
	created_at           string
	response             string
	done                 bool
	total_duration       int64
	load_duration        int64
	prompt_eval_count    int64
	prompt_eval_duration int64
	eval_count           int64
	eval_duration        int64
}
