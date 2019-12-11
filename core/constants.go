// Copyright (C) 2017-2019 go-nebulas authors
//
// This file is part of the go-nebulas library.
//
// the go-nebulas library is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// the go-nebulas library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with the go-nebulas library.  If not, see <http://www.gnu.org/licenses/>.
//

package core

// TODO: update if contract is ready
var (
	// AccessContract given to access config load contract.
	AccessContract, _ = AddressParse("n1wRERsLCoGsh2YZu7Qy74iFrraJwnV9gKX")

	// PoDContract given to pod consensus contract
	PoDContract, _ = AddressParse("n1xS3BoziPPidb5nXDmGfH9pb4RHQMfyBe9")

	// GovernanceContract given to governance contract
	GovernanceContract, _ = AddressParse("n1gTeXSSfvVNq4wHMURpafDYfRRzBA4JwKN")
)
