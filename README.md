# snowmark - HTML templates for Go.

[![Build Status](https://github.com/sangupta/snowmark/actions/workflows/unittest.yml/badge.svg?branch=main)](https://github.com/sangupta/snowmark/actions)
[![Code Coverage](https://codecov.io/gh/sangupta/snowmark/branch/main/graphs/badge.svg?branch=main)](https://codecov.io/gh/sangupta/snowmark)
[![go.mod version](https://img.shields.io/github/go-mod/go-version/sangupta/snowmark.svg)](https://github.com/sangupta/snowmark)
![GitHub](https://img.shields.io/github/license/sangupta/snowmark)

`snowmark` is a library for HTML templates that uses HTML
custom tags and custom attributes instead of confusing markup
intertwined within HTML. It is very similar to Java Server
Pages but uses `string` manipulation to merge templates. In
some ways `snowmark` is also similar to `Velocity` templates
except using custom tags.

# Table of contents

- [Features](#features)
- [API](#api)
- [Usage Example](#usage-example)
- [Hacking](#hacking)
- [Changelog](#changelog)
- [License](#license)

# Features

* Merge HTML templates with custom model
* Bring your own custom tags
* Attribute expressions
* Standard tag library includes:
  - Get variable
  - Set variable (global or in block)
  - If-then

# API

# Usage Example

Let's use the following example HTML template that we want to
merge with our own data model:

```html
<html>
    <head>
        <title>
            <test:get var="pageTitle" />
        </title>
    </head>
</html>
```

The data model to be used represented as JSON is:

```json
{
    "pageTitle" : "Hello World"
}
```

The following code allows us to do the same:

```go
// parse HTML doc
template := "<html><head><title><get var='pageTitle' /></title></head></html>"

// create the model
model := snowmark.NewModel()
model.Put("pageTitle", "Hello World")

// create a page processor
processor := snowmark.NewHtmlPageProcessor()

// add all your custom tags
// you have the choice to name each tag differently
processor.AddCustomTag("test:get", snowmark.GetVariableTag)

// call merge
html, _ := processor.MergeHtml(template, model)
fmt.Println(html)
```

# Hacking

* To build the Go docs locally:
  - `$ godoc -http=:6060`
  - Open http://localhost:6060/pkg/github.com/sangupta/snowmark

* To run all tests along with code coverage report
  - `$ go test ./... -v -coverprofile coverage.out`
  - `$ go tool cover -html=coverage.out`

# Changelog

* **Version 0.1.0**
  - Initial release

# License

MIT License. Copyright (C) 2022, Sandeep Gupta.
