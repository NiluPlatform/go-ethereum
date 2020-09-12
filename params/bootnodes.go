// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Ethereum Foundation Go Bootnodes
	"enode://604b383ae13167e91abc99aaf557c8488419a6dc6b64660f38bab979a16778a994ee0b7796c82bdde1286e1bc93a4da900284bf210af2a0690049e6b0005e0c5@144.76.139.77:30303",
	"enode://ab52728a722fd4b0560caee3d6fcd341812309476f1fe95700d83e8f64069aed5a74e5f25fddcaeec60ccecea4e3e8cde9c09353f791a8c670530fca5c2ce9ff@209.250.234.227:30303",
	"enode://c18dbf793b646bbd1d94f834eb2fc1fa9385bf0b18f7b3149497c187667c135a21aa090eb7f7049ccb918ee24c8b21d69ca3720f38f23f2a67cff40ad3a0a120@45.77.177.233:30303",
	"enode://2aaaf66b5b4f2632780bc2160efaa4e0c3c8c5735d5f0cafcea96c48b5017b74efc7a68ba44b4d4a95bed46a794eba4bec73d4620d11f4cc93f7d88ca5518340@144.202.1.133:30303",
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
}
