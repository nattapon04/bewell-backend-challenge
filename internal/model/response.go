package model

type ValidateMessage struct {
	StatusText string         `json:"status_text,omitempty" bson:"status_text,omitempty" example:"Validation Failed"`
	Error      *ValidateError `json:"error,omitempty"`
}

type ValidateError struct {
	ErrorCode string `json:"error_code,omitempty" example:"REQUIRED"`
	Field     string `json:"field,omitempty" example:"name"`
	Message   string `json:"message,omitempty" example:"Validation failed on field 'name', condition: required"`
}

type Message struct {
	StatusText  string      `json:"status_text,omitempty" example:"Unprocessable Entity"`
	Error       *Error      `json:"error,omitempty"`
	Description interface{} `json:"description,omitempty"`
}

type Error struct {
	ErrorCode string `json:"error_code,omitempty" example:"UNPROCESSABLE_ENTITY"`
	Message   string `json:"message,omitempty"`
}
