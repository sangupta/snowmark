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

//
// Represents a model of values
// that are used to merge with a page or a fragment.
//
type Model struct {
	_map map[string]interface{}
}

//
// Create new model instance.
//
func NewModel() *Model {
	return &Model{
		_map: make(map[string]interface{}),
	}
}

//
// Return the size of the model.
//
func (model *Model) Size() int {
	return len(model._map)
}

//
// Check if model is empty or not. Returns `true` if
// there are no keys in model, `false` otherwise.
//
func (model *Model) IsEmpty() bool {
	return model.Size() == 0
}

//
// Clear model and remove all keys
//
func (model *Model) Clear() {
	model._map = make(map[string]interface{})
}

//
// Get the value from the model.
//
func (model *Model) Get(key string) (interface{}, bool) {
	value, exists := model._map[key]
	return value, exists
}

//
// Put the value in model against given key.
//
func (model *Model) Put(key string, value interface{}) {
	model._map[key] = value
}

//
// Put the value in model if the key does not exist. Returns
// `true` if value was added to model, `false` otherwise.
//
func (model *Model) PutIfNotExists(key string, value interface{}) bool {
	_, exists := model._map[key]
	if exists {
		return false
	}

	model._map[key] = value
	return true
}

//
// Replace the current key with new value. If no key exists
// the function returns `false`. Returns `true` if value was
// replaced.
//
func (model *Model) Replace(key string, value interface{}) bool {
	_, exists := model._map[key]
	if !exists {
		return false
	}

	model._map[key] = value
	return true
}

//
// Remove the key from the model, if it exists.
//
func (model *Model) Remove(key string) {
	delete(model._map, key)
}
