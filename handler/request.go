package handler

import "fmt"

type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

// Helper function
func errParamIsValid(name, typ string) error {
	return fmt.Errorf("param: %s ( type %s ) is required", name, typ)
}
func (r *CreateOpeningRequest) Validate() error {

	if r.Company == "" {
		return errParamIsValid("company", "string")
	}
	if r.Location == "" {
		return errParamIsValid("location", "string")
	}
	if r.Remote == nil {
		return errParamIsValid("remote", "bool")
	}
	if r.Link == "" {
		return errParamIsValid("link", "string")
	}
	if r.Salary <= 0 {
		return errParamIsValid("salary", "int64")
	}
	return nil
}

type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (o *UpdateOpeningRequest) Validate() error {
	// if any field is provided, validation is truty
	if o.Role != "" || o.Company != "" || o.Link != "" || o.Location != "" || o.Remote != nil || o.Salary > 0 {
		return nil
	}
	// if none of the fields are provided, returns error
	return fmt.Errorf("at least one field on request must be provided")
}
