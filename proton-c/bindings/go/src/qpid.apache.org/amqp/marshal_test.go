/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

package amqp

import (
	"testing"
)

func TestSymbolKey(t *testing.T) {
	bytes, err := Marshal(SymbolKey("foo"), nil)
	if err != nil {
		t.Fatal(err)
	}
	var k Key
	if _, err := Unmarshal(bytes, &k); err != nil {
		t.Error(err)
	}
	if err := checkEqual("foo", string(k.Get().(Symbol))); err != nil {
		t.Error(err)
	}
	var sym Symbol
	if _, err := Unmarshal(bytes, &sym); err != nil {
		t.Error(err)
	}
	if err := checkEqual("foo", sym.String()); err != nil {
		t.Error(err)
	}

}

func TestStringKey(t *testing.T) {
	bytes, err := Marshal(StringKey("foo"), nil)
	if err != nil {
		t.Fatal(err)
	}
	var k Key
	if _, err := Unmarshal(bytes, &k); err != nil {
		t.Error(err)
	}
	if err := checkEqual("foo", string(k.Get().(Symbol))); err != nil {
		t.Error(err)
	}
	var s string
	if _, err := Unmarshal(bytes, &s); err != nil {
		t.Error(err)
	}
	if err := checkEqual("foo", s); err != nil {
		t.Error(err)
	}

}

func TestIntKey(t *testing.T) {
	bytes, err := Marshal(IntKey(12345), nil)
	if err != nil {
		t.Fatal(err)
	}
	var k Key
	if _, err := Unmarshal(bytes, &k); err != nil {
		t.Error(err)
	}
	if 12345 != k.Get().(uint64) {
		t.Errorf("(%T)%v != (%T)%v", 12345, k.Get().(uint64))
	}
	var n uint64
	if _, err := Unmarshal(bytes, &n); err != nil {
		t.Error(err)
	}
	if 12345 != n {
		t.Errorf("%v != %v", 12345, k.Get().(uint64))
	}

}
