// Copyright 2019 The go-aroeum Authors
// This file is part of the go-aroeum library.
//
// The go-aroeum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-aroeum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-aroeum library. If not, see <http://www.gnu.org/licenses/>.

var OpCodes = artifacts.require("./OpCodes.sol");

module.exports = function(deployer) {
  deployer.deploy(OpCodes);
};
