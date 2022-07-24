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
	"strings"

	"github.com/sangupta/lhtml"
)

//
// Merge given HTML string with the given model.
//
func (pageProcessor *HtmlPageProcessor) MergeHtml(html string, model *Model) (string, error) {
	elements, err := lhtml.ParseHtmlString(html)
	if err != nil {
		return "", err
	}

	return pageProcessor.Merge(elements, model)
}

//
// Merge given parsed HTML document with the given model.
//
func (pageProcessor *HtmlPageProcessor) Merge(elements *lhtml.HtmlElements, model *Model) (string, error) {
	if elements.IsEmpty() {
		return "", nil
	}

	// create evaluator
	builder := strings.Builder{}
	evaluator := &Evaluator{
		builder:   &builder,
		processor: pageProcessor,
	}

	evaluator.EvaluateNodes(elements.Nodes(), model)

	return evaluator.builder.String(), nil
}
