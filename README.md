<p align="center">
<img src="https://i.ibb.co/J5dgL4v/Webp-net-resizeimage.png">
</p>
<h1 align="center" style="text-decoration:none">go-vdom-wasm</h1>

This package allows to build frontend applications in Go targeting [Webassembly](https://webassembly.org/). It provides a conveniant syntax close to the [hyperscript](https://github.com/hyperhype/hyperscript) one in the JavaScript lands and aims to provide _efficient_ way to deal with the [DOM](https://developer.mozilla.org/en-US/docs/Web/API/Document_Object_Model). If you're not familiar the Virtual DOM notion, I suggest you take a look at [this document]().

---

:warning: This package is not production ready. Use at your own risks. A [Todos](#todos) section is available if you want to take part of the projet and try to make it better. You can also check the [Go + WASM wiki](https://github.com/golang/go/wiki/WebAssembly) to get a better overview of the Webassembly support for the Go language.

This module has been inspired by the awesome work of [@mbasso](https://github.com/mbasso) on https://github.com/mbasso/asm-dom.

[Want to see go-vdom-wasm in action?](https://mfrachet.github.io/go-vdom-wasm/)

# Ready to start?

- [Before starting](#before-starting)
- [Install the module](#install-the-module)
- [Usage](#usage)
- [Todos](#todos)

## Before starting

You have to make some little configuration that are well explained in the [Go + WASM wiki - Getting started](https://github.com/golang/go/wiki/WebAssembly#getting-started) guide.

If you prefer a _quickest_ solution, you can also rely on the [go-wasm-cli](https://github.com/mfrachet/go-wasm-cli) utility. It allows to create a WASM application in one command line and to have hot reload while coding.

## Install the module

In your favorite terminal

```shell
$ go get github.com/mfrachet/go-vdom-wasm
```

## Usage

Let's define the root of our `go-vdom-wasm` application. Everything that the library will work on is **inside** that `div`:

```html
<!-- index.html-->

<div id="app"></div>
```

Let's now `Patch` that node with a virtual one to make something happen:

```go
// main.go

rootNode := vn.H("div", nil, "Hello world")
vn.Patch("#app", rootNode)
```

### Passing attributes

The second argument of `vn.H` is a reference to an `Attrs` structure. It allows to manage `HTML attributes` and events.

In this example, we set a `class="navbar"` to the underlying `ul` element

```go
vn.H("ul", &vn.Attrs{Props: &vn.Props{"class": "navbar"}}, vn.Children{
	vn.H("li", nil, "First item"),
	vn.H("li", nil, "Second item"),
}),
```

The [DOM](https://developer.mozilla.org/en-US/docs/Web/API/Document_Object_Model) representation of the previous snippet is:

```html
<ul class="navbar">
  <li>First item</li>
  <li>Second item</li>
</ul>
```

### Handling events

The `Attrs` also allows to pass functions to handle specific events.

In this example, we set a click event on of the the `li`:

```go
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
- [ ] Generating the API doc and adding decent comment in code
- [ ] Find a way to abstract the `[]js.Value` passed in every event handler
