// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by "mdtogo"; DO NOT EDIT.
package overview

var ReferenceShort = `Overview of kpt commands`
var ReferenceLong = `
kpt functionality is subdivided into command groups, each of which operates on
a particular set of entities, with a consistent command syntax and pattern of
inputs and outputs.
`
var ReferenceExamples = `
  # get a package
  $ kpt pkg get https://github.com/GoogleContainerTools/kpt.git/package-examples/helloworld-set@v0.5.0 helloworld
  fetching package /package-examples/helloworld-set from \
    https://github.com/GoogleContainerTools/kpt to helloworld

  # list setters and set a value
  $ kpt cfg list-setters helloworld
  NAME            DESCRIPTION         VALUE    TYPE     COUNT   SETBY
  http-port   'helloworld port'         80      integer   3
  image-tag   'hello-world image tag'   v0.3.0  string    1
  replicas    'helloworld replicas'     5       integer   1
  
  $ kpt cfg set helloworld replicas 3 --set-by pwittrock  --description 'reason'
  set 1 fields

  # apply the package to a clsuter
  $ kpt live apply --wait-for-reconcile helloworld
  ...
  all resources has reached the Current status
`
