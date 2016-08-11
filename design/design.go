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
})

var _ = Resource("todo", func() {
	BasePath("/todos")
	DefaultMedia(TodoMedia)

	Action("show", func() {
		Description("Get todo by id")
		Routing(GET("/:todoID"))
		Params(func() {
			Param("todoID", Integer, "Todo ID")
			Response(OK)
			Response(NotFound)
		})
	})
})

// TodoMedia defines the media type used to render todos.
var TodoMedia = MediaType("application/vnd.duetoday.todo+json", func() {
	Description("A To Do")
	Attributes(func() {
		Attribute("id", Integer, "Unique Todo ID")
		Attribute("title", String, "Todo title")
		Attribute("details", String, "Additional notes")
		Required("id", "title")
	})
	View("default", func() {
		Attribute("id")
		Attribute("title")
		Attribute("details")
	})
})
