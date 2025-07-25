/*
 * Copyright 2025 The Go-Spring Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package assert_test

import (
	"testing"

	"github.com/go-spring/assert"
	"github.com/go-spring/assert/internal"
)

func TestMap_Length(t *testing.T) {
	m := new(internal.MockTestingT)
	testMap := map[string]int{"a": 1}

	m.Reset()
	assert.ThatMap(m, testMap).Length(1)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, testMap).Length(0)
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected map to have length 0, but it has length 1
  actual: {"a":1}`)

	m.Reset()
	assert.ThatMap(m, testMap).Must().Length(0, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected map to have length 0, but it has length 1
  actual: {"a":1}
 message: "index is 0"`)
}

func TestMap_Nil(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.ThatMap(m, map[string]int(nil)).Nil()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, map[string]int{"a": 1}).Nil()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected map to be nil, but it is not
  actual: {"a":1}`)

	m.Reset()
	assert.ThatMap(m, map[string]int{"a": 1}).Must().Nil("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected map to be nil, but it is not
  actual: {"a":1}
 message: "index is 0"`)
}

func TestMap_NotNil(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.ThatMap(m, map[string]int{"a": 1}).NotNil()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, map[string]int(nil)).NotNil()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected map not to be nil, but it is
  actual: null`)

	m.Reset()
	assert.ThatMap(m, map[string]int(nil)).Must().NotNil("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected map not to be nil, but it is
  actual: null
 message: "index is 0"`)
}

func TestMap_IsEmpty(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.ThatMap(m, map[string]int(nil)).Empty()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, map[string]int{"a": 1}).Empty()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected map to be empty, but it is not
  actual: {"a":1}`)

	m.Reset()
	assert.ThatMap(m, map[string]int{"a": 1}).Must().Empty("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected map to be empty, but it is not
  actual: {"a":1}
 message: "index is 0"`)
}

func TestMap_IsNotEmpty(t *testing.T) {
	m := new(internal.MockTestingT)
	testMap := map[string]int{"a": 1}
	assert.ThatMap(m, testMap).NotEmpty()
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, map[string]int(nil)).NotEmpty()
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected map to be non-empty, but it is empty
  actual: null`)

	m.Reset()
	assert.ThatMap(m, map[string]int{}).Must().NotEmpty("index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected map to be non-empty, but it is empty
  actual: {}
 message: "index is 0"`)
}

func TestMap_Equal(t *testing.T) {
	m := new(internal.MockTestingT)
	testMap := map[string]int{"a": 1}

	m.Reset()
	assert.ThatMap(m, testMap).Equal(map[string]int{"a": 1})
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, testMap).Equal(nil)
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected maps to be equal, but their lengths are different
  actual: {"a":1}
expected: null`)

	m.Reset()
	assert.ThatMap(m, testMap).Equal(map[string]int{"b": 2})
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected maps to be equal, but key 'a' is missing
  actual: {"a":1}
expected: {"b":2}`)

	m.Reset()
	assert.ThatMap(m, testMap).Must().Equal(map[string]int{"a": 2}, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected maps to be equal, but values for key 'a' are different
  actual: {"a":1}
expected: {"a":2}
 message: "index is 0"`)
}

func TestMap_NotEqual(t *testing.T) {
	m := new(internal.MockTestingT)
	testMap := map[string]int{"a": 1}

	m.Reset()
	assert.ThatMap(m, testMap).NotEqual(map[string]int{"b": 2})
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, testMap).NotEqual(testMap)
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected maps to be different, but they are equal
  actual: {"a":1}`)

	m.Reset()
	assert.ThatMap(m, testMap).Must().NotEqual(testMap, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected maps to be different, but they are equal
  actual: {"a":1}
 message: "index is 0"`)
}

func TestMap_ContainsKey(t *testing.T) {
	m := new(internal.MockTestingT)
	testMap := map[string]int{"a": 1}

	m.Reset()
	assert.ThatMap(m, testMap).ContainsKey("a")
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, testMap).ContainsKey("b")
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected map to contain key 'b', but it is missing
  actual: {"a":1}`)

	m.Reset()
	assert.ThatMap(m, testMap).Must().ContainsKey("b", "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected map to contain key 'b', but it is missing
  actual: {"a":1}
 message: "index is 0"`)
}

func TestMap_NotContainsKey(t *testing.T) {
	m := new(internal.MockTestingT)
	testMap := map[string]int{"a": 1}

	m.Reset()
	assert.ThatMap(m, testMap).NotContainsKey("b")
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, testMap).NotContainsKey("a")
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected map not to contain key 'a', but it is found
  actual: {"a":1}`)

	m.Reset()
	assert.ThatMap(m, testMap).Must().NotContainsKey("a", "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected map not to contain key 'a', but it is found
  actual: {"a":1}
 message: "index is 0"`)
}

func TestMap_ContainsValue(t *testing.T) {
	m := new(internal.MockTestingT)
	testMap := map[string]int{"a": 1}

	m.Reset()
	assert.ThatMap(m, testMap).ContainsValue(1)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, testMap).ContainsValue(2)
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected map to contain value 2, but it is missing
  actual: {"a":1}`)

	m.Reset()
	assert.ThatMap(m, testMap).Must().ContainsValue(2, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected map to contain value 2, but it is missing
  actual: {"a":1}
 message: "index is 0"`)
}

func TestMap_NotContainsValue(t *testing.T) {
	m := new(internal.MockTestingT)
	testMap := map[string]int{"a": 1}

	m.Reset()
	assert.ThatMap(m, testMap).NotContainsValue(2)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, testMap).NotContainsValue(1)
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected map not to contain value 1, but it is found
  actual: {"a":1}`)

	m.Reset()
	assert.ThatMap(m, testMap).Must().NotContainsValue(1, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected map not to contain value 1, but it is found
  actual: {"a":1}
 message: "index is 0"`)
}

func TestMap_ContainsKeyValue(t *testing.T) {
	m := new(internal.MockTestingT)
	testMap := map[string]int{"a": 1}

	m.Reset()
	assert.ThatMap(m, testMap).ContainsKeyValue("a", 1)
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, testMap).ContainsKeyValue("b", 2)
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected map to contain key 'b', but it is missing
  actual: {"a":1}`)

	m.Reset()
	assert.ThatMap(m, testMap).Must().ContainsKeyValue("a", 2, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected value 2 for key 'a', but got 1 instead
  actual: {"a":1}
 message: "index is 0"`)
}

func TestMap_ContainsKeys(t *testing.T) {
	m := new(internal.MockTestingT)
	testMap := map[string]int{"a": 1, "b": 2}

	m.Reset()
	assert.ThatMap(m, testMap).ContainsKeys([]string{"a", "b"})
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, testMap).ContainsKeys([]string{"c"})
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected map to contain key 'c', but it is missing
  actual: {"a":1,"b":2}`)

	m.Reset()
	assert.ThatMap(m, testMap).Must().ContainsKeys([]string{"c"}, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected map to contain key 'c', but it is missing
  actual: {"a":1,"b":2}
 message: "index is 0"`)
}

func TestMap_NotContainsKeys(t *testing.T) {
	m := new(internal.MockTestingT)
	testMap := map[string]int{"a": 1, "b": 2}

	m.Reset()
	assert.ThatMap(m, testMap).NotContainsKeys([]string{"c"})
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, testMap).NotContainsKeys([]string{"a"})
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected map not to contain key 'a', but it is found
  actual: {"a":1,"b":2}`)

	m.Reset()
	assert.ThatMap(m, testMap).Must().NotContainsKeys([]string{"a"}, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected map not to contain key 'a', but it is found
  actual: {"a":1,"b":2}
 message: "index is 0"`)
}

func TestMap_ContainsValues(t *testing.T) {
	m := new(internal.MockTestingT)
	testMap := map[string]int{"a": 1, "b": 2}

	m.Reset()
	assert.ThatMap(m, testMap).ContainsValues([]int{1, 2})
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, testMap).ContainsValues([]int{3})
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected map to contain value 3, but it is missing
  actual: {"a":1,"b":2}`)

	m.Reset()
	assert.ThatMap(m, testMap).Must().ContainsValues([]int{3}, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected map to contain value 3, but it is missing
  actual: {"a":1,"b":2}
 message: "index is 0"`)
}

func TestMap_NotContainsValues(t *testing.T) {
	m := new(internal.MockTestingT)
	testMap := map[string]int{"a": 1, "b": 2}

	m.Reset()
	assert.ThatMap(m, testMap).NotContainsValues([]int{3})
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, testMap).NotContainsValues([]int{1})
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected map not to contain value 1, but it is found
  actual: {"a":1,"b":2}`)

	m.Reset()
	assert.ThatMap(m, testMap).Must().NotContainsValues([]int{1}, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected map not to contain value 1, but it is found
  actual: {"a":1,"b":2}
 message: "index is 0"`)
}

func TestMap_SubsetOf(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.ThatMap(m, map[string]int{"a": 1}).SubsetOf(map[string]int{"a": 1, "b": 2})
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, map[string]int{"a": 1, "b": 2}).SubsetOf(map[string]int{"a": 1})
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected map to be a subset, but unexpected key 'b' is found
  actual: {"a":1,"b":2}
expected: {"a":1}`)

	m.Reset()
	assert.ThatMap(m, map[string]int{"a": 1}).Must().SubsetOf(map[string]int{"a": 2}, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected map to be a subset, but values for key 'a' are different
  actual: {"a":1}
expected: {"a":2}
 message: "index is 0"`)
}

func TestMap_SupersetOf(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.ThatMap(m, map[string]int{"a": 1, "b": 2}).SupersetOf(map[string]int{"a": 1})
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, map[string]int{"a": 1}).SupersetOf(map[string]int{"a": 1, "b": 2})
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected map to be a superset, but key 'b' is missing
  actual: {"a":1}
expected: {"a":1,"b":2}`)

	m.Reset()
	assert.ThatMap(m, map[string]int{"a": 1}).Must().SupersetOf(map[string]int{"a": 2}, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected map to be a superset, but values for key 'a' are different
  actual: {"a":1}
expected: {"a":2}
 message: "index is 0"`)
}

func TestMap_HasSameKeys(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.ThatMap(m, map[string]int{"a": 1, "b": 2}).HasSameKeys(map[string]int{"b": 3, "a": 4})
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, map[string]int{"a": 1, "b": 2}).HasSameKeys(map[string]int{"c": 3})
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected maps to have the same keys, but their lengths are different
  actual: {"a":1,"b":2}
expected: {"c":3}`)

	m.Reset()
	assert.ThatMap(m, map[string]int{"a": 1, "b": 2}).Must().HasSameKeys(map[string]int{"b": 2, "c": 3}, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected maps to have the same keys, but key 'a' is missing
  actual: {"a":1,"b":2}
expected: {"b":2,"c":3}
 message: "index is 0"`)
}

func TestMap_HasSameValues(t *testing.T) {
	m := new(internal.MockTestingT)

	m.Reset()
	assert.ThatMap(m, map[string]int{"a": 1, "b": 2}).HasSameValues(map[string]int{"x": 1, "y": 2})
	assert.ThatString(t, m.String()).Equal("")

	m.Reset()
	assert.ThatMap(m, map[string]int{"a": 1, "b": 2}).HasSameValues(map[string]int{"c": 3})
	assert.ThatString(t, m.String()).Equal(`error# Assertion failed: expected maps to have the same values, but their lengths are different
  actual: {"a":1,"b":2}
expected: {"c":3}`)

	m.Reset()
	assert.ThatMap(m, map[string]int{"a": 1, "b": 2}).Must().HasSameValues(map[string]int{"b": 2, "c": 3}, "index is 0")
	assert.ThatString(t, m.String()).Equal(`fatal# Assertion failed: expected maps to have the same values, but their values are different
  actual: {"a":1,"b":2}
expected: {"b":2,"c":3}
 message: "index is 0"`)
}
