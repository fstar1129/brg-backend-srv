package app

// ErrorMsg ...
type ErrorMsg struct {
	Type  string `json:"type"`
	Error string `json:"error"`
}

func createNewError(err, msg string) ErrorMsg {
	return ErrorMsg{
		Type:  err,
		Error: msg,
	}
}
