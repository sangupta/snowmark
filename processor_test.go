/**
 * snowmark - HTML templates for Go.
 *
 * MIT License.
 * Copyright (c) 2022, Sandeep Gupta.
 * https://github.com/sangupta/snowmark
 *
 * Use of this source code is governed by a MIT style license
 * that can be found in LICENSE file in the code repository:
 */

package snowmark

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessorEmptyTemplate(t *testing.T) {
	processor := NewHtmlPageProcessor()

	processor.AddCustomTag("if", IfElseTag)
	processor.AddCustomTag("get", GetVariableTag)
	processor.AddCustomTag("set", SetVariableTag)
	processor.AddCustomTag("for", ForEachTag)

	template := ""
	model := NewModel()

	html, err := processor.MergeHtml(template, model)
	assert.NoError(t, err)
	assert.Equal(t, "", html)
}

func TestProcessorNoTag(t *testing.T) {
	processor := NewHtmlPageProcessor()
	template := "<html></html>"
	model := NewModel()

	html, err := processor.MergeHtml(template, model)
	assert.NoError(t, err)
	assert.Equal(t, "<html />", html)
}

func TestProcessorGetTagNoValue(t *testing.T) {
	processor := NewHtmlPageProcessor()
	processor.AddCustomTag("if", IfElseTag)
	processor.AddCustomTag("get", GetVariableTag)
	processor.AddCustomTag("set", SetVariableTag)
	processor.AddCustomTag("for", ForEachTag)

	template := "<html><get var='hello' /></html>"
	model := NewModel()

	html, err := processor.MergeHtml(template, model)
	assert.NoError(t, err)
	assert.Equal(t, "<html></html>", html)
}

func TestProcessorGetTagWithValue(t *testing.T) {
	processor := NewHtmlPageProcessor()
	processor.AddCustomTag("if", IfElseTag)
	processor.AddCustomTag("get", GetVariableTag)
	processor.AddCustomTag("set", SetVariableTag)
	processor.AddCustomTag("for", ForEachTag)

	template := "<html><get var='hello' /></html>"
	model := NewModel()
	model.Put("hello", "world")

	html, err := processor.MergeHtml(template, model)
	assert.NoError(t, err)
	assert.Equal(t, "<html>world</html>", html)
}
