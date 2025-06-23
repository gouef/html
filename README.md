<img align=right width="168" src="docs/gouef_logo.png">

# Html
Go code generate html

[![GoDoc](https://pkg.go.dev/badge/github.com/gouef/html.svg)](https://pkg.go.dev/github.com/gouef/html)
[![GitHub stars](https://img.shields.io/github/stars/gouef/html?style=social)](https://github.com/gouef/html/stargazers)
[![Go Report Card](https://goreportcard.com/badge/github.com/gouef/html)](https://goreportcard.com/report/github.com/gouef/html)
[![codecov](https://codecov.io/github/gouef/html/branch/main/graph/badge.svg?token=YUG8EMH6Q8)](https://codecov.io/github/gouef/html)

## Versions
![Stable Version](https://img.shields.io/github/v/release/gouef/html?label=Stable&labelColor=green)
![GitHub Release](https://img.shields.io/github/v/release/gouef/html?label=RC&include_prereleases&filter=*rc*&logoSize=diago)
![GitHub Release](https://img.shields.io/github/v/release/gouef/html?label=Beta&include_prereleases&filter=*beta*&logoSize=diago)

## Instalation

```shell
go get -u github.com/gouef/html
```

## Usages

```go
package main

import "github.com/gouef/html"

func getHtml() string {
	el := html.El("div")
	el.AddAttribute("class", "container-fluid")
	child := html.El("p").AddString("some text")
	el.AddHtml(child)
	
	return el.Render()
}
```

## Commit rules
Commit message should looks like
```
[TYPE] some message
```

### Types
 - Add
 - Fix
 - Update
 - Remove
 - Refactor
 - Docs
 - Test
 - Improve

## Contributors

<div>
<span>
  <a href="https://github.com/actions-user"><img src="https://raw.githubusercontent.com/gouef/html/refs/heads/contributors-svg/.github/contributors/actions-user.svg" alt="actions-user" /></a>
</span>
<span>
  <a href="https://github.com/JanGalek"><img src="https://raw.githubusercontent.com/gouef/html/refs/heads/contributors-svg/.github/contributors/JanGalek.svg" alt="JanGalek" /></a>
</span>
</div>

