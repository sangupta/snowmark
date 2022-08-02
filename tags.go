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
	"reflect"

	"github.com/sangupta/berry"
	"github.com/sangupta/lhtml"
)

//
// A simple tag to get the value of any variable inside the model.
//
//   <get value="name" />
//
func GetVariableTag(node *lhtml.HtmlNode, model *Model, evaluator *Evaluator) error {
	value, err := evaluator.GetAttributeValueAsString(node, "var", model)
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
// For example, the variable `hello` is set to `world` for the rest of
// processing.
//
//  <setVar var="hello" value="world" />
//
// Another way is to use children, in which case the older value (if any)
// will be restored at the closing of tag. In the following example, the
// value of variable `hello` is only available to the children of the
// `setVar` node.
//
//  <setVar var="hello" value="world">
//     ...
//  </setVar>
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
	newValue, err := evaluator.GetAttributeValueAsString(node, "value", model)
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
// A simple if-else tag.
//
//  <custom:if condition="age > 10">
//     <custom:then>
//     </custom:then>
//     </custom:else>
//     </custom:else>
//  </custom:if>
func IfElseTag(node *lhtml.HtmlNode, model *Model, evaluator *Evaluator) error {
	condition, err := node.GetAttributeValue("condition")
	if err != nil {
		return err
	}

	conditionValue, err := evaluator.EvaluateExpression(condition, model)
	if err != nil {
		return err
	}

	conditionBool, err := berry.ConvertToBool(conditionValue)
	if err != nil {
		return err
	}

	if conditionBool {
		thenClause := node.GetChildByName("then")
		if thenClause == nil {
			return errors.New("If tag does not have a 'then' clause")
		}

		return evaluator.EvaluateNodes(thenClause.Children(), model)
	}

	// do else part
	elseClause := node.GetChildByName("else")
	if elseClause == nil {
		return nil
	}

	return evaluator.EvaluateNodes(elseClause.Children(), model)
}

//
// A for-each tag that takes in a `collection` and then adds it as the
// given variable name. If the collection is a `slice` the variable gets
// one collection object at a time.
//
//  <foreach collection="mySlice" item="sliceItem">
//     <get var="sliceItem" />
//  </foreach>
//
// If the collection is a `map` the variable gets an object that contains
// two properties: `key` and `value` representing the key/value pair for
// each entry in the map.
//
//  <foreach collection="myMap" item="mapItem">
//     <get var="mapItem.key" />
//     <get var="mapItem.value" />
//  </foreach>
//
func ForEachTag(node *lhtml.HtmlNode, model *Model, evaluator *Evaluator) error {
	// get expression which represents collection
	expression, err := node.GetAttributeValue("collection")
	if err != nil {
		return err
	}

	// get variable name that we want to add
	variableName, err := node.GetAttributeValue("var")
	if err != nil {
		return err
	}

	// find collection
	collection, err := evaluator.EvaluateExpression(expression, model)
	if err != nil {
		return err
	}

	// if not object in collection, do nothing
	if collection == nil {
		return nil
	}

	// preserve old value
	olderValue, olderValueExists := model.Get(variableName)

	// start checking what we are iterating over
	switch reflect.TypeOf(collection).Kind() {
	case reflect.Slice:
		slice := reflect.ValueOf(collection)
		for index := 0; index < slice.Len(); index++ {
			item := slice.Index(index)

			// now run the nodes with this value
			model.Put(variableName, item)

			// evaluate all child nodes
			evaluator.EvaluateNodes(node.Children(), model)
		}

	case reflect.Map:
		mapp := reflect.ValueOf(collection)
		for _, key := range mapp.MapKeys() {
			value := mapp.MapIndex(key)

			// put this in model
			pair := map[string]interface{}{
				"key":   key,
				"value": value,
			}
			model.Put(variableName, pair)

			// evaluate all child nodes
			evaluator.EvaluateNodes(node.Children(), model)
		}
	}

	// we are done iterating
	if olderValueExists {
		model.Put(variableName, olderValue)
	} else {
		model.Remove(variableName)
	}

	// all done
	return nil
}
