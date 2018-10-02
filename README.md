_:warning: This is an experimental package. Don't use it in production app for now :blush:_

Webassembly VDOM to create web application using Golang

---

# Content

- Installation
- Usage

## Installation

```shell
$ go get https://github.com/mfrachet/go-wasm-vdom
```

## Usage

### A first (v)node

_index.html_

```html
<div id="app"></div>
```

_main.go_

```golang
rootNode := vn.H("div", nil, "Hello world")
vn.Patch("#app", rootNode)
```

### Using props

The second argument of `vn.H` is a reference to an `Attrs` structure. It allows to manage `HTML attributes` and events.

In this example, we set a `class="navbar"` to the underlying `ul` element

```golang
vn.H("ul", &vn.Attrs{Props: &vn.Props{"class": "navbar"}}, vn.Children{
	vn.H("li", nil, "First item"),
	vn.H("li", nil, "Second item"),
}),
```

### Handling events

The `Attrs` also allows to pass a function for a specific events.

In this example, we set a click event on of the the `li`:

```golang
func handleClick(args []js.Value){
	fmt.Println("I've been called")
}

/* ... */

vn.H("ul", &vn.Attrs{Props: &vn.Props{"class": "navbar"}}, vn.Children{
	vn.H("li", &vn.Attrs{Events: &vn.Ev{"click": handleClick}}, "First item"),
	vn.H("li", nil, "Second item"),
}),
```

:warning: While using event handler, it's necessary to add en empty `select{}` at the end of the `main` function
