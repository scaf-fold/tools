package courier

import (
	"net/http"
)

type MethodGet struct {
	Method
}

func (m MethodGet) String() string {
	return http.MethodGet
}

type MethodHead struct {
	Method
}

func (m MethodHead) String() string {
	return http.MethodHead
}

type MethodPost struct {
	Method
}

func (m MethodPost) String() string {
	return http.MethodPost
}

type MethodPut struct {
	Method
}

func (m MethodPut) String() string {
	return http.MethodPut
}

type MethodPatch struct {
	Method
}

func (m MethodPatch) String() string {
	return http.MethodPatch
}

type MethodDelete struct {
	Method
}

func (m MethodDelete) String() string {
	return http.MethodDelete
}

type MethodConnect struct {
	Method
}

func (m MethodConnect) String() string {
	return http.MethodConnect
}

type MethodOptions struct {
	Method
}

func (m MethodOptions) String() string {
	return http.MethodOptions
}

type MethodTrace struct {
	Method
}

func (m MethodTrace) String() string {
	return http.MethodTrace
}
