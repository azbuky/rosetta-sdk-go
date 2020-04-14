// Copyright 2020 Coinbase, Inc.
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

package types

// ConstructPartialBlockIdentifier constructs a *PartialBlockIdentifier
// from a *BlockIdentifier.
//
// It is useful to have this helper when making block requests
// with the fetcher.
func ConstructPartialBlockIdentifier(
	blockIdentifier *BlockIdentifier,
) *PartialBlockIdentifier {
	return &PartialBlockIdentifier{
		Hash:  &blockIdentifier.Hash,
		Index: &blockIdentifier.Index,
	}
}