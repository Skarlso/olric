// Copyright 2018-2025 The Olric Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package routingtable

import (
	"sync/atomic"
	"testing"
	"time"

	"github.com/olric-data/olric/internal/testutil"
)

func TestRoutingTable_Callback(t *testing.T) {
	c := testutil.NewConfig()
	rt := newRoutingTableForTest(c, testutil.NewServer(c))
	var num int32
	increase := func() {
		atomic.AddInt32(&num, 1)
	}
	rt.AddCallback(increase)
	rt.wg.Add(1)
	go rt.runCallbacks()
	<-time.After(100 * time.Millisecond)
	modified := atomic.LoadInt32(&num)
	if modified != 1 {
		t.Fatalf("Expected number: 1. Got: %v", modified)
	}
}
