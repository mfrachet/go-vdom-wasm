![go-vdom-wasm](https://i.ibb.co/kmpdhXW/govdom.png)

_:warning: This is an experimental package. It's NOT production ready :blush:_

Webassembly VDOM to create web apps using Go. Widely inspired by https://github.com/mbasso/asm-dom

A TODO application on gh-pages: https://mfrachet.github.io/go-vdom-wasm/ . For the sources, [check this branch](https://github.com/mfrachet/go-vdom-wasm/tree/gh-pages).

---

# Content

- [Installation](#installation)
- [Usage](#usage)
- [Todos](#todos)

## Installation

### Configuring Go for Webassembly

You first need to follow a webassembly project structure for Go [like this one](https://github.com/golang/go/wiki/WebAssembly).

_:information_source: You can also use https://github.com/mfrachet/go-wasm-cli to quickly start a go-webassembly project._

### In your favorite terminal

Try running:

```shell
$ go get github.com/mfrachet/go-vdom-wasm
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

## Todos

- [ ] Add more tests :woman_facepalming: :man_facepalming:
- [ ] Find a better solution [to compare two virtual nodes](https://github.com/mfrachet/go-vdom-wasm/blob/bddbb032b6c048cf6ee58368241f4b3d3c427691/vnode.go#L24)
- [ ] Ensure consistency and avoid unecessary rerendering that sometimes occur for the same DOMNodes :question:
- [ ] Make a better abstraction concerning the reconciler (it's a really first step)
- [ ] Rely on Go good practices (reference passing, better algorithms and so forth...)
