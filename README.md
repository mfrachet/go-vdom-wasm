_:warning: This is an experimental package. It's NOT production ready :blush:_

Webassembly VDOM to create web apps using Go. Widely inspired by https://github.com/mbasso/asm-dom

An example is available through github pages: https://mfrachet.github.io/go-vdom-wasm/

---

# Content

- [Installation](#installation)
- [Usage](#usage)

## Installation

You first need to follow a webassembly project structure for Go [like this one](https://github.com/golang/go/wiki/WebAssembly).

*You can also use https://github.com/mfrachet/go-wasm-cli to quickly start a go-webassembly project.*

Then simply run:

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

The expected result would be:

```html
<ul class="navbar">
    <li>First item</li>
    <li>Second item</li>
</ul>
```



### Handling events

The `Attrs` also allows to pass a function for a specific events.

In this example, we set a click event on of the the `li`:

```golang
func handleClick(args []js.Value){
	fmt.Println("I've been clicked!")
}

func main() {
    rootNode := vn.H("ul", &vn.Attrs{Props: &vn.Props{"class": "navbar"}}, vn.Children{
        vn.H("li", &vn.Attrs{Events: &vn.Ev{"click": handleClick}}, "First item"),
        vn.H("li", nil, "Second item"),
	})
	
	vn.patch("#app", rootNode)
}

```

This module binds the event using the [`addEventListener` API](https://developer.mozilla.org/en-US/docs/Web/API/EventTarget/addEventListener)

:warning: While using event handler, it's necessary to add en empty `select{}` at the end of the `main` function

