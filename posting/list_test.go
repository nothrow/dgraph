/*
 * Copyright 2015 Manish R Jain <manishrjain@gmail.com>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 		http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package posting

import (
	"testing"
	"time"

	"github.com/manishrjain/dgraph/posting/types"
	"github.com/manishrjain/dgraph/x"
)

var uids = [...]uint64{
	9, 49, 81,
}

func TestAddTriple(t *testing.T) {
	var l List
	l.Init()

	triple := x.Triple{
		ValueId:   9,
		Source:    "testing",
		Timestamp: time.Now(),
	}
	l.AddTriple(triple)

	if l.TList.PostingsLength() != 1 {
		t.Error("Unable to find added elements in posting list")
	}
	var p types.Posting
	if ok := l.TList.Postings(&p, 0); !ok {
		t.Error("Unable to retrieve posting at 1st iter")
		t.Fail()
	}
	if p.Uid() != 9 {
		t.Errorf("Expected 9. Got: %v", p.Uid)
	}
	if string(p.Source()) != "testing" {
		t.Errorf("Expected testing. Got: %v", string(p.Source()))
	}

	// Add another triple now.
	triple.ValueId = 81
	l.AddTriple(triple)
	if l.TList.PostingsLength() != 2 {
		t.Errorf("Length: %d", l.TList.PostingsLength())
		t.Fail()
	}

	var uid uint64
	uid = 1
	for i := 0; i < l.TList.PostingsLength(); i++ {
		if ok := l.TList.Postings(&p, i); !ok {
			t.Error("Unable to retrieve posting at 2nd iter")
		}
		uid *= 9
		if p.Uid() != uid {
			t.Errorf("Expected: %v. Got: %v", uid, p.Uid())
		}
	}

	// Add another triple, in between the two above.
	triple.ValueId = 49
	l.AddTriple(triple)
	if l.TList.PostingsLength() != 3 {
		t.Errorf("Length: %d", l.TList.PostingsLength())
		t.Fail()
	}
	for i := 0; i < len(uids); i++ {
		if ok := l.TList.Postings(&p, i); !ok {
			t.Error("Unable to retrieve posting at 2nd iter")
		}
		if p.Uid() != uids[i] {
			t.Errorf("Expected: %v. Got: %v", uids[i], p.Uid())
		}
	}
}
