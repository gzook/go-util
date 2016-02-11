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

package cntr_test

import (
    "math"
    "sync"
    "testing"
    "github.com/gzook/go-util/cntr"
)

func TestCounterInt64(t *testing.T) {
    
    // Start thread-safe testing of plus etc
    var wg sync.WaitGroup
    iMax := 1000
    jMax := 100000

    wg.Add(iMax)
    cPlusOne := cntr.NewInt64()
	for i := 0; i < iMax; i++ {
		go func(no int) {
            defer wg.Done()

			for j := 0; j < jMax; j++ {
				cPlusOne.PlusOne()
			}
		}(i)
	}

    wg.Add(iMax)
    cPlusX := cntr.NewInt64()
	for i := 0; i < iMax; i++ {
		go func(no int) {
            defer wg.Done()

			for j := 0; j < jMax; j++ {
				cPlusX.PlusX(2)
			}
		}(i)
	}

	wg.Wait()
    if cPlusOne.Value() != int64(iMax * jMax) {
        t.Errorf("Int64 counter PlusOne value incorrect, expected %v, got %v", iMax * jMax, cPlusOne.Value())
    }
    if cPlusX.Value() != int64(iMax * jMax * 2) {
        t.Errorf("Int64 counter PlusX value incorrect, expected %v, got %v", iMax * jMax, cPlusOne.Value())
    }

    // Other basic tests
    cPlusOne.ToZero()
    if cPlusOne.Value() != 0 {
        t.Error("Int64 counter ToZero() did not zero the value")
    }

    cPlusOne.ToValue(math.MaxInt64)
    if cPlusOne.Value() != math.MaxInt64 {
        t.Error("Int64 counter ToValue() did not set the value")
    }
    
    cPlusOne.ToValue(math.MinInt64)
    if cPlusOne.Value() != math.MinInt64 {
        t.Error("Int64 counter ToValue() did not set the value")
    }
}

func TestCounterInt(t *testing.T) {
    
    // Start thread-safe testing of plus etc
    var wg sync.WaitGroup
    iMax := 1000
    jMax := 100000

    wg.Add(iMax)
    cPlusOne := cntr.NewInt()
	for i := 0; i < iMax; i++ {
		go func(no int) {
            defer wg.Done()

			for j := 0; j < jMax; j++ {
				cPlusOne.PlusOne()
			}
		}(i)
	}

    wg.Add(iMax)
    cPlusX := cntr.NewInt()
	for i := 0; i < iMax; i++ {
		go func(no int) {
            defer wg.Done()

			for j := 0; j < jMax; j++ {
				cPlusX.PlusX(2)
			}
		}(i)
	}

	wg.Wait()
    if cPlusOne.Value() != int(iMax * jMax) {
        t.Errorf("Int counter PlusOne value incorrect, expected %v, got %v", iMax * jMax, cPlusOne.Value())
    }
    if cPlusX.Value() != int(iMax * jMax * 2) {
        t.Errorf("Int counter PlusX value incorrect, expected %v, got %v", iMax * jMax, cPlusOne.Value())
    }

    // Other basic tests
    cPlusOne.ToZero()
    if cPlusOne.Value() != 0 {
        t.Error("Int counter ToZero() did not zero the value")
    }

    cPlusOne.ToValue(23)
    if cPlusOne.Value() != 23 {
        t.Error("Int counter ToValue() did not set the value")
    }

}