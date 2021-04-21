// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: customer.proto

package customer

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	regexp "regexp"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

var _regex_Customer_Id = regexp.MustCompile(`^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$`)
var _regex_Customer_FirstName = regexp.MustCompile(`^[A-Z][a-zA-Z]+$`)
var _regex_Customer_LastName = regexp.MustCompile(`^[A-Z][a-zA-Z]+$`)
var _regex_Customer_PhoneNumber = regexp.MustCompile(`^01\d{9}$`)
var _regex_Customer_Email = regexp.MustCompile(`.+\@.+\..+`)
var _regex_Customer_Password = regexp.MustCompile(`^\d{4}$`)

func (this *Customer) Validate() error {
	if !_regex_Customer_Id.MatchString(this.Id) {
		return github_com_mwitkow_go_proto_validators.FieldError("Id", fmt.Errorf(`value '%v' must be a string conforming to regex "^([a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[4][a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12})?$"`, this.Id))
	}
	if !_regex_Customer_FirstName.MatchString(this.FirstName) {
		return github_com_mwitkow_go_proto_validators.FieldError("FirstName", fmt.Errorf(`value '%v' must be a string conforming to regex "^[A-Z][a-zA-Z]+$"`, this.FirstName))
	}
	if this.FirstName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("FirstName", fmt.Errorf(`value '%v' must not be an empty string`, this.FirstName))
	}
	if !_regex_Customer_LastName.MatchString(this.LastName) {
		return github_com_mwitkow_go_proto_validators.FieldError("LastName", fmt.Errorf(`value '%v' must be a string conforming to regex "^[A-Z][a-zA-Z]+$"`, this.LastName))
	}
	if this.LastName == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("LastName", fmt.Errorf(`value '%v' must not be an empty string`, this.LastName))
	}
	if !_regex_Customer_PhoneNumber.MatchString(this.PhoneNumber) {
		return github_com_mwitkow_go_proto_validators.FieldError("PhoneNumber", fmt.Errorf(`value '%v' must be a string conforming to regex "^01\\d{9}$"`, this.PhoneNumber))
	}
	if this.PhoneNumber == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PhoneNumber", fmt.Errorf(`value '%v' must not be an empty string`, this.PhoneNumber))
	}
	if !_regex_Customer_Email.MatchString(this.Email) {
		return github_com_mwitkow_go_proto_validators.FieldError("Email", fmt.Errorf(`value '%v' must be a string conforming to regex ".+\\@.+\\..+"`, this.Email))
	}
	if this.Email == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Email", fmt.Errorf(`value '%v' must not be an empty string`, this.Email))
	}
	if !_regex_Customer_Password.MatchString(this.Password) {
		return github_com_mwitkow_go_proto_validators.FieldError("Password", fmt.Errorf(`value '%v' must be a string conforming to regex "^\\d{4}$"`, this.Password))
	}
	if this.Password == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Password", fmt.Errorf(`value '%v' must not be an empty string`, this.Password))
	}
	if this.CreatedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CreatedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CreatedAt", err)
		}
	}
	if this.UpdatedAt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.UpdatedAt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("UpdatedAt", err)
		}
	}
	return nil
}
func (this *ReadAllRequest) Validate() error {
	return nil
}
func (this *ReadAllResponse) Validate() error {
	for _, item := range this.Customer {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Customer", err)
			}
		}
	}
	return nil
}
func (this *CreateRequest) Validate() error {
	if this.Customer != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Customer); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Customer", err)
		}
	}
	return nil
}
func (this *CreateResponse) Validate() error {
	if this.Customer != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Customer); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Customer", err)
		}
	}
	return nil
}
func (this *ReadRequest) Validate() error {
	return nil
}
func (this *ReadResponse) Validate() error {
	if this.Customer != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Customer); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Customer", err)
		}
	}
	return nil
}
func (this *UpdateRequest) Validate() error {
	if this.Customer != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Customer); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Customer", err)
		}
	}
	return nil
}
func (this *UpdateResponse) Validate() error {
	if this.Customer != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Customer); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Customer", err)
		}
	}
	return nil
}
func (this *DeleteRequest) Validate() error {
	return nil
}
func (this *DeleteResponse) Validate() error {
	return nil
}
