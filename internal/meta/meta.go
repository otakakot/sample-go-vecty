package meta

import "syscall/js"

var (
	doc  = js.Global().Get("document")
	head = doc.Get("head")
)

func Set(name string, property string, content string) {
	meta := doc.Call("createElement", "meta")
	if name != "" {
		meta.Set("name", name)
	}
	if property != "" {
		meta.Set("property", property)
	}
	if content != "" {
		meta.Set("content", content)
	}
	head.Call("appendChild", meta)
}
