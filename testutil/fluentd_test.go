// Copyright 2018 pixiv Technologies Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package testutil_test

import (
	"context"
	"testing"
	"time"

	"github.com/pixiv/go-fluent-quicksilver/testutil"
)

func TestSpawnFluentd(t *testing.T) {
	config := ``
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	fluentd, err := testutil.SpawnFluentd(ctx, config, nil)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(time.Second)

	fluentd.Kill()
	<-fluentd.Exit
}