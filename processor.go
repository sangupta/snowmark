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

func NewHtmlPageProcessor() *HtmlPageProcessor {
	return &HtmlPageProcessor{
		_tags: make(map[string]CustomTagProcessor),
	}
}

func (pageProcessor *HtmlPageProcessor) AddCustomTag(name string, tagProcessor CustomTagProcessor) (bool, error) {
	if name == "" {
		return false, errors.New("Name cannot be empty")
	}

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

func (pageProcessor *HtmlPageProcessor) RemoveCustomTag(name string) (bool, error) {
	if name == "" {
		return false, errors.New("Name cannot be empty")
	}

	delete(pageProcessor._tags, name)
	return true, nil
}

func (pageProcessor *HtmlPageProcessor) HasCustomTag(name string) bool {
	if name == "" {
		return false
	}

	_, exists := pageProcessor._tags[name]
	return exists
}

func (pageProcessor *HtmlPageProcessor) GetCustomTag(name string) (CustomTagProcessor, bool) {
	if name == "" {
		return nil, false
	}

	tag, exists := pageProcessor._tags[name]
	return tag, exists
}
