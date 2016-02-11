/*
The MIT License (MIT)

Copyright (c) 2016 gzook

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package cntr

import (
    "sync"
)

// Int64 is a threadsafe counter of type int64
type Int64 struct {
	mu sync.Mutex
	v  int64
}

// NewInt64 creates a new Int64 counter
func NewInt64() *Int64 {
    return &Int64{}
}

// PlusOne increments the counter by 1
func (c *Int64) PlusOne() {
	c.mu.Lock()
	c.v++
	c.mu.Unlock()
}

// PlusX increments the counter by X
func (c *Int64) PlusX(x int64) {
	c.mu.Lock()
	c.v += x
	c.mu.Unlock()
}

// Value returns the value of the counter
func (c *Int64) Value() (v int64) {
	c.mu.Lock()
	v = c.v
	c.mu.Unlock()
	return
}

// ToZero sets the value of the counter to 0
func (c *Int64) ToZero() {
	c.mu.Lock()
	c.v = 0
	c.mu.Unlock()
}

// ToValue sets the value of the counter to a given value
func (c *Int64) ToValue(v int64) {
	c.mu.Lock()
	c.v = v
	c.mu.Unlock()
}

// Int is a threadsafe counter of type int - i.e. will be int64 on 64 bit OS, int32 on 32 bit
type Int struct {
	mu sync.Mutex
	v  int
}

// NewInt creates a new Int counter
func NewInt() *Int {
    return &Int{}
}

// PlusOne increments the counter by 1
func (c *Int) PlusOne() {
	c.mu.Lock()
	c.v++
	c.mu.Unlock()
}

// PlusX increments the counter by X
func (c *Int) PlusX(x int) {
	c.mu.Lock()
	c.v += x
	c.mu.Unlock()
}

// Value returns the value of the counter
func (c *Int) Value() (v int) {
	c.mu.Lock()
	v = c.v
	c.mu.Unlock()
	return
}

// ToZero sets the value of the counter to 0
func (c *Int) ToZero() {
	c.mu.Lock()
	c.v = 0
	c.mu.Unlock()
}

// ToValue sets the value of the counter to a given value
func (c *Int) ToValue(v int) {
	c.mu.Lock()
	c.v = v
	c.mu.Unlock()
}