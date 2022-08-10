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

	"github.com/maja42/goval"
	"github.com/sangupta/berry"
	"github.com/sangupta/lhtml"
)

const PREFIX = "expr:"

//
// The Evaluator instance.
//
type Evaluator struct {
	builder   *strings.Builder
	processor *HtmlPageProcessor
}

//
// Write given string to internal string builder.
//
func (evaluator *Evaluator) WriteString(s string) (int, error) {
	return evaluator.builder.WriteString(s)
}

//
// Write byte array to internal string builder.
//
func (evaluator *Evaluator) Write(b []byte) (int, error) {
	return evaluator.builder.Write(b)
}

//
// Write rune to internal string builder.
//
func (evaluator *Evaluator) WriteRune(r rune) (int, error) {
	return evaluator.builder.WriteRune(r)
}

//
// Write single byte to internal string builder.
//
func (evaluator *Evaluator) WriteByte(b byte) error {
	return evaluator.builder.WriteByte(b)
}

func (evaluator *Evaluator) GetEvaluation() string {
	return evaluator.builder.String()
}

//
// Evaluate the given node against the model.
//
func (evaluator *Evaluator) EvaluateNode(node *lhtml.HtmlNode, model *Model) error {
	if node == nil {
		return nil
	}

	nodeName := node.NodeName()

	// custom tag, process it differently?
	customTag, exists := evaluator.processor.GetCustomTag(nodeName)
	if exists {
		return customTag(node, model, evaluator)
	}

	// process a normal tag
	evaluator.processNormalNode(node, model)
	return nil
}

//
// Evaluate multiple nodes against the model.
//
func (evaluator *Evaluator) EvaluateNodes(nodes []*lhtml.HtmlNode, model *Model) error {
	if nodes == nil || len(nodes) == 0 {
		return nil
	}

	for _, node := range nodes {
		err := evaluator.EvaluateNode(node, model)
		if err != nil {
			return err
		}
	}

	return nil
}

//
// Evaluate an expression against the model.
//
func (evaluator *Evaluator) EvaluateExpression(expr string, model *Model) (interface{}, error) {
	if expr == "" {
		return "", nil
	}

	eval := goval.NewEvaluator()
	return eval.Evaluate(expr, model._map, nil)
}

//
// Evaluate an expression against the model and return resulting
// value as a `string` using default platform conversion.
//
func (evaluator *Evaluator) EvaluateExpressionAsString(expr string, model *Model) (string, error) {
	value, err := evaluator.EvaluateExpression(expr, model)
	if err != nil {
		return "", err
	}

	return berry.ConvertToString(value), nil
}

//
// Get an attribute's value from the node. It also checks if an expression
// attribute with same name is present, and if yes, if will evalue the expression
// and return the value.
//
func (evaluator *Evaluator) GetAttributeValue(node *lhtml.HtmlNode, attributeName string, model *Model) (interface{}, error) {
	if attributeName == "" {
		return "", errors.New("Attribute name is required")
	}

	if node == nil {
		return "", errors.New("Node is required to read attribute from")
	}

	attr := node.GetAttribute(attributeName)
	if attr != nil {
		return attr.Value, nil
	}

	attr = node.GetAttribute("expr:" + attributeName)
	if attr == nil {
		return "", errors.New("Missing attribute 'value'")
	}

	// evaluate expression
	return evaluator.EvaluateExpression(attr.Value, model)
}

func (evaluator *Evaluator) GetAttributeValueAsString(node *lhtml.HtmlNode, attributeName string, model *Model) (string, error) {
	value, err := evaluator.GetAttributeValue(node, attributeName, model)
	if err != nil {
		return "", err
	}

	return berry.ConvertToString(value), nil
}

func (evaluator *Evaluator) processNormalNode(node *lhtml.HtmlNode, model *Model) error {
	if node == nil {
		return nil
	}

	// local reference to builder
	builder := evaluator.builder

	// doctype, text or comment?
	if node.NodeType == lhtml.DoctypeNode || node.NodeType == lhtml.TextNode || node.NodeType == lhtml.CommentNode {
		builder.WriteString(node.Data)
		return nil
	}

	// this is an element node
	// start building
	builder.WriteString("<")
	builder.WriteString(node.NodeName())

	// attributes
	if node.ContainsAttributes() {
		for _, attr := range node.Attributes {
			name := attr.Name
			value := attr.Value

			if strings.HasPrefix(name, PREFIX) {
				// evaluate expression
				updatedValue, err := evaluator.EvaluateExpressionAsString(value, model)
				if err != nil {
					return err
				}

				// update name and value
				name = strings.TrimPrefix(name, PREFIX)
				value = updatedValue
			}

			builder.WriteString(" ")
			builder.WriteString(name)
			builder.WriteString("=\"")
			builder.WriteString(value)
			builder.WriteString("\"")
		}
	}

	// self-closing?
	if !node.HasChildren() {
		builder.WriteString(" />")
	} else {
		builder.WriteString(">")

		// work on children
		if node.HasChildren() {
			for _, child := range node.Children() {
				evaluator.EvaluateNode(child, model)
			}
		}

		// close
		builder.WriteString("</")
		builder.WriteString(node.NodeName())
		builder.WriteString(">")
	}

	return nil
}
