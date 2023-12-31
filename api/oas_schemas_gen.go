// Code generated by ogen, DO NOT EDIT.

package api

import (
	"fmt"
)

func (s *MessageStatusCode) Error() string {
	return fmt.Sprintf("code %d: %+v", s.StatusCode, s.Response)
}

type BasicAuth struct {
	Username string
	Password string
}

// GetUsername returns the value of Username.
func (s *BasicAuth) GetUsername() string {
	return s.Username
}

// GetPassword returns the value of Password.
func (s *BasicAuth) GetPassword() string {
	return s.Password
}

// SetUsername sets the value of Username.
func (s *BasicAuth) SetUsername(val string) {
	s.Username = val
}

// SetPassword sets the value of Password.
func (s *BasicAuth) SetPassword(val string) {
	s.Password = val
}

type CookieAuth struct {
	APIKey string
}

// GetAPIKey returns the value of APIKey.
func (s *CookieAuth) GetAPIKey() string {
	return s.APIKey
}

// SetAPIKey sets the value of APIKey.
func (s *CookieAuth) SetAPIKey(val string) {
	s.APIKey = val
}

// Ref: #/components/schemas/Message
type Message struct {
	// Message from the server.
	Message string `json:"message"`
	// Error reason.
	Error OptString `json:"error"`
}

// GetMessage returns the value of Message.
func (s *Message) GetMessage() string {
	return s.Message
}

// GetError returns the value of Error.
func (s *Message) GetError() OptString {
	return s.Error
}

// SetMessage sets the value of Message.
func (s *Message) SetMessage(val string) {
	s.Message = val
}

// SetError sets the value of Error.
func (s *Message) SetError(val OptString) {
	s.Error = val
}

// MessageStatusCode wraps Message with StatusCode.
type MessageStatusCode struct {
	StatusCode int
	Response   Message
}

// GetStatusCode returns the value of StatusCode.
func (s *MessageStatusCode) GetStatusCode() int {
	return s.StatusCode
}

// GetResponse returns the value of Response.
func (s *MessageStatusCode) GetResponse() Message {
	return s.Response
}

// SetStatusCode sets the value of StatusCode.
func (s *MessageStatusCode) SetStatusCode(val int) {
	s.StatusCode = val
}

// SetResponse sets the value of Response.
func (s *MessageStatusCode) SetResponse(val Message) {
	s.Response = val
}

// NewOptInt returns new OptInt with value set to v.
func NewOptInt(v int) OptInt {
	return OptInt{
		Value: v,
		Set:   true,
	}
}

// OptInt is optional int.
type OptInt struct {
	Value int
	Set   bool
}

// IsSet returns true if OptInt was set.
func (o OptInt) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt) SetTo(v int) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt) Get() (v int, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Ref: #/components/schemas/Pagination
type Pagination struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

// GetLimit returns the value of Limit.
func (s *Pagination) GetLimit() int {
	return s.Limit
}

// GetOffset returns the value of Offset.
func (s *Pagination) GetOffset() int {
	return s.Offset
}

// GetTotal returns the value of Total.
func (s *Pagination) GetTotal() int {
	return s.Total
}

// SetLimit sets the value of Limit.
func (s *Pagination) SetLimit(val int) {
	s.Limit = val
}

// SetOffset sets the value of Offset.
func (s *Pagination) SetOffset(val int) {
	s.Offset = val
}

// SetTotal sets the value of Total.
func (s *Pagination) SetTotal(val int) {
	s.Total = val
}

// Ref: #/components/schemas/User
type User struct {
	UID  string `json:"uid"`
	Name string `json:"name"`
	// Groups the user belongs to in RFC 1123 format.
	Group       []string  `json:"group"`
	Description OptString `json:"description"`
}

// GetUID returns the value of UID.
func (s *User) GetUID() string {
	return s.UID
}

// GetName returns the value of Name.
func (s *User) GetName() string {
	return s.Name
}

// GetGroup returns the value of Group.
func (s *User) GetGroup() []string {
	return s.Group
}

// GetDescription returns the value of Description.
func (s *User) GetDescription() OptString {
	return s.Description
}

// SetUID sets the value of UID.
func (s *User) SetUID(val string) {
	s.UID = val
}

// SetName sets the value of Name.
func (s *User) SetName(val string) {
	s.Name = val
}

// SetGroup sets the value of Group.
func (s *User) SetGroup(val []string) {
	s.Group = val
}

// SetDescription sets the value of Description.
func (s *User) SetDescription(val OptString) {
	s.Description = val
}

// Ref: #/components/schemas/UserCreateRequest
type UserCreateRequest struct {
	Name string `json:"name"`
	// Groups the user belongs to in RFC 1123 format.
	Group       []string  `json:"group"`
	Password    string    `json:"password"`
	Description OptString `json:"description"`
}

// GetName returns the value of Name.
func (s *UserCreateRequest) GetName() string {
	return s.Name
}

// GetGroup returns the value of Group.
func (s *UserCreateRequest) GetGroup() []string {
	return s.Group
}

// GetPassword returns the value of Password.
func (s *UserCreateRequest) GetPassword() string {
	return s.Password
}

// GetDescription returns the value of Description.
func (s *UserCreateRequest) GetDescription() OptString {
	return s.Description
}

// SetName sets the value of Name.
func (s *UserCreateRequest) SetName(val string) {
	s.Name = val
}

// SetGroup sets the value of Group.
func (s *UserCreateRequest) SetGroup(val []string) {
	s.Group = val
}

// SetPassword sets the value of Password.
func (s *UserCreateRequest) SetPassword(val string) {
	s.Password = val
}

// SetDescription sets the value of Description.
func (s *UserCreateRequest) SetDescription(val OptString) {
	s.Description = val
}

// Ref: #/components/schemas/UserList
type UserList struct {
	Items      []User     `json:"items"`
	Pagination Pagination `json:"pagination"`
}

// GetItems returns the value of Items.
func (s *UserList) GetItems() []User {
	return s.Items
}

// GetPagination returns the value of Pagination.
func (s *UserList) GetPagination() Pagination {
	return s.Pagination
}

// SetItems sets the value of Items.
func (s *UserList) SetItems(val []User) {
	s.Items = val
}

// SetPagination sets the value of Pagination.
func (s *UserList) SetPagination(val Pagination) {
	s.Pagination = val
}
