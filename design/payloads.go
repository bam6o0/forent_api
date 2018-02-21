package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// SigninPayload is Signin payload
var SigninPayload = Type("SigninPayload", func() {
	Attribute("email", String, func() {
		MinLength(6)
		MaxLength(400)
		Format("email")
		Example("jamesbond@gmail.com")
	})

	Attribute("password", String, func() {
		MinLength(8)
		MaxLength(100)
		Example("abcd1234")
	})
	Required("email", "password")
})

// SignupPayload is Signin payload
var SignupPayload = Type("SignupPayload", func() {
	Attribute("email", String, func() {
		MinLength(6)
		MaxLength(150)
		Format("email")
		Example("jamesbond@gmail.com")
	})

	Attribute("password", String, func() {
		MinLength(8)
		MaxLength(100)
		Example("abcd1234")
	})

	Required("email", "password")
})
