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
	"errors"
	"strings"

	"github.com/sangupta/lhtml"
)

type CustomTagProcessor func(node *lhtml.HtmlNode, model *Model, evaluator *Evaluator) error

//
// The main interface that allows you to add custom
// tag processors and then use this instance to merge
// HTML templates.
//
type HtmlPageProcessor struct {
	_tags map[string]CustomTagProcessor
}

//
// Create a new instance of HTML page processor.
//
func NewHtmlPageProcessor() *HtmlPageProcessor {
	return &HtmlPageProcessor{
		_tags: make(map[string]CustomTagProcessor),
	}
}

//
// Add a custom tag to the processor which will make use of the provided
// `func` to work upon. The tag name is case insensitive. If a tag with
// the same name already exists, an error is returned.
//
func (pageProcessor *HtmlPageProcessor) AddCustomTag(name string, tagProcessor CustomTagProcessor) (bool, error) {
	if name == "" {
		return false, errors.New("Name cannot be empty")
	}

	name = strings.TrimSpace(name)
	if name == "" {
		return false, errors.New("Name cannot be blank")
	}

	name = strings.ToLower(name)
	if tagProcessor == nil {
		return false, errors.New("Custom tag processor cannot be nil")
	}

	// check if we already have a tag with same name
	_, exists := pageProcessor._tags[name]
	if exists {
		return false, errors.New("Tag already exists")
	}

	pageProcessor._tags[name] = tagProcessor
	return true, nil
}

//
// Remove a custom tag with given name. An error is returned if the name
// is empty, or blank. No error is returned if no custom tag processor
// can be found attached to this name.
//
func (pageProcessor *HtmlPageProcessor) RemoveCustomTag(name string) (bool, error) {
	if name == "" {
		return false, errors.New("Name cannot be empty")
	}

	name = strings.TrimSpace(name)
	if name == "" {
		return false, errors.New("Name cannot be blank")
	}

	name = strings.ToLower(name)
	delete(pageProcessor._tags, name)
	return true, nil
}

//
// Check if a custom tag processor is attached for the given name.
//
func (pageProcessor *HtmlPageProcessor) HasCustomTag(name string) bool {
	if name == "" {
		return false
	}

	name = strings.TrimSpace(name)
	if name == "" {
		return false
	}

	name = strings.ToLower(name)
	_, exists := pageProcessor._tags[name]
	return exists
}

//
// Return the custom tag processor attached for the given name. If the name
// is empty or blank, a `nil` is returned.
//
func (pageProcessor *HtmlPageProcessor) GetCustomTag(name string) (CustomTagProcessor, bool) {
	if name == "" {
		return nil, false
	}

	name = strings.TrimSpace(name)
	if name == "" {
		return nil, false
	}

	name = strings.ToLower(name)
	tag, exists := pageProcessor._tags[name]
	return tag, exists
}
