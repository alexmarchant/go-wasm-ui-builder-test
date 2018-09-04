package main

import (
  "syscall/js"
  "strconv"
)

var i = 0
var document js.Value
var root js.Value

type Node interface {
  JsEl() js.Value
}

type Text struct {
  Value string
}

func (t *Text) JsEl() js.Value {
  return document.Call("createTextNode", t.Value)
}

type Element struct {
  Name string
  Children []Node
  OnClick func()
}

func (e *Element) JsEl() js.Value {
  el := document.Call("createElement", e.Name)

  if (e.OnClick != nil) {
    cb := js.NewCallback(func(args []js.Value) {
      e.OnClick()
      updateDom()
    })
    el.Call("addEventListener", "click", cb)
  }

  for _, child := range e.Children {
    el.Call("appendChild", child.JsEl())
  }

  return el
}

func e(name string, children []Node, onClick func()) *Element {
  return &Element{name, children, onClick}
}

func t(value string) *Text {
  return &Text{value}
}

func c(children ...Node) []Node {
  return children
}

func inc() {
  i += 1
}

func render() js.Value {
  return e("div", c(
    e("h1", c(t("h1")), nil),
    e("h2", c(t("h2")), nil),
    e("h3", c(t("h3")), nil),
    e("p", c(t("p")), nil),
    e("button", c(t("button")), inc),
    e("p", c(t("Counter: "), t(strconv.Itoa(i))), nil),
  ), nil).JsEl()
}

func updateDom() {
  root.Set("innerHTML", "")
  root.Call("appendChild", render())
}

func main() {
  c := make(chan struct{}, 0)
  document = js.Global().Get("document")
  root = document.Call("getElementById", "app")
  updateDom()
  <-c
}

