//    Copyright 2018 Google Inc.
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package sst

// Iterator is an iterator over keys and values.
// Iterators MUST be closed when done.
type Iterator interface {
	// Key returns the current key.
	Key() string

	// Value returns the current value. Callers should not modify the returned value.
	// Returned value is only valid until the next call to Next.
	Value() []byte

	// Next advances the iterator. Returns whether there is another key/value.
	Next() bool

	// Close closes the iterator.
	Close()
}
