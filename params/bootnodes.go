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
	//"enode://075aa1d81891b410b873d7aa33d20e44cc701daee0dc57eef3995a801a27a0004eccb3e90ba4c04717858284e7a0b27998dabebcf6b708e0c2668f241d0b7d45@164.132.141.202:30303",
	"enode://58224a2145aac14a63b499e79b8c10b59cb97d7d3e1ed424334714fbbee39cf2d4f3df7f46726f19842fd545d2c4c8562a323b94a6bb8b6db88c6c5bb9623085@147.135.243.251:30303",
	"enode://ab52728a722fd4b0560caee3d6fcd341812309476f1fe95700d83e8f64069aed5a74e5f25fddcaeec60ccecea4e3e8cde9c09353f791a8c670530fca5c2ce9ff@209.250.234.227:30303",
	"enode://c18dbf793b646bbd1d94f834eb2fc1fa9385bf0b18f7b3149497c187667c135a21aa090eb7f7049ccb918ee24c8b21d69ca3720f38f23f2a67cff40ad3a0a120@45.77.177.233:30303",
	"enode://2aaaf66b5exitb4f2632780bc2160efaa4e0c3c8c5735d5f0cafcea96c48b5017b74efc7a68ba44b4d4a95bed46a794eba4bec73d4620d11f4cc93f7d88ca5518340@144.202.1.133:30303",
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
