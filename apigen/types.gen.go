// Package apigen provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package apigen

// TestObject defines model for TestObject.
type TestObject struct {
	Id   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Q    *string `json:"q,omitempty"`
}

// Q defines model for Q.
type Q = string

// GetTestParams defines parameters for GetTest.
type GetTestParams struct {
	// Q test query param
	Q Q `form:"q" json:"q"`
}
