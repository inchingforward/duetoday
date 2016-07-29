package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("duetoday", func() {
		Title("Due Today")
		Description("A todo service")
		Scheme("http")
		Host("localhost:4030")
}
