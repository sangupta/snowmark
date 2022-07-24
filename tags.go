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
	"strconv"

	"github.com/sangupta/lhtml"
)

//
// <custom:var value="name" />
//
func GetVariableTag(node *lhtml.HtmlNode, model *Model, evaluator *Evaluator) error {
	value, err := evaluator.GetAttributeValue(node, "var", model)
	if err != nil {
		return err
	}

	value, err = evaluator.EvaluateExpressionAsString(value, model)
	if err != nil {
		return err
	}

	evaluator.builder.WriteString(value)
	return nil
}

//
// You can use it in two ways. For a self-closing tag, the model will
// be modified immediately and any older value will not be preserved.
//
//  <custom:setVar var="hello" value="world" />
//
// Another way is to use children, in which case the older value (if any)
// will be restored at the closing of tag.
//
//  <custom:setVar>
//     ...
//  </custom:setVar>
//
func SetVariableTag(node *lhtml.HtmlNode, model *Model, evaluator *Evaluator) error {
	// get variable name
	variableName, err := node.GetAttributeValue("var")
	if err != nil {
		return err
	}

	if variableName == "" {
		return errors.New("Variable name cannot be empty")
	}

	// get new value
	newValue, err := evaluator.GetAttributeValue(node, "value", model)
	if err != nil {
		return err
	}

	// preserve old value
	olderValue, olderValueExists := model.Get(variableName)

	// set new value
	model.Put(variableName, newValue)

	if node.HasChildren() {
		// process all children
		evaluator.EvaluateNodes(node.Children(), model)

		// once we are done, recover to older value
		if olderValueExists {
			model.Put(variableName, olderValue)
		}
	}

	return nil
}

//
//  <custom:if condition="age > 10">
//     <custom:then>
//     </custom:then>
//     </custom:else>
//     </custom:else>
//  </custom:if>
func IfElseTag(node *lhtml.HtmlNode, model *Model, evaluator *Evaluator) error {
	condition, err := evaluator.GetAttributeValue(node, "condition", model)
	if err != nil {
		return err
	}

	conditionBool, err := strconv.ParseBool(condition)
	if err != nil {
		return err
	}

	if conditionBool {
		thenClause := node.GetChildByName("then")
		if thenClause == nil {
			return errors.New("If tag does not have a 'then' clause")
		}

		return evaluator.EvaluateNode(thenClause, model)
	}

	// do else part
	elseClause := node.GetChildByName("else")
	if elseClause == nil {
		return nil
	}

	return evaluator.EvaluateNode(elseClause, model)
}
