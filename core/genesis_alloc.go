// Copyright 2017 The go-nilu Authors
// This file is part of the go-nilu library.
//
// The go-nilu library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-nilu library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-nilu library. If not, see <http://www.gnu.org/licenses/>.

package core

// Constants containing the genesis allocation of built-in genesis blocks.
// Their content is an RLP-encoded list of (address, balance) tuples.
// Use mkalloc.go to create/update them.

// nolint: misspell
const mainnetAllocData = "\xe1\xe0\x94\xc2\xd7\u03d5d]3\x00au\xb7\x89\x89\x03\\|\x90a\xd3\xf9\x8a\xbc:\xc3b}E\x00\xe3\x8e8"
const testnetAllocData = "\xe1\xe0\x94\xc2\xd7\u03d5d]3\x00au\xb7\x89\x89\x03\\|\x90a\xd3\xf9\x8a\xbc:\xc3b}E\x00\xe3\x8e8"
const rinkebyAllocData = "\xe1\xe0\x94\xc2\xd7\u03d5d]3\x00au\xb7\x89\x89\x03\\|\x90a\xd3\xf9\x8a\xbc:\xc3b}E\x00\xe3\x8e8"