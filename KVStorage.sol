// SPDX-License-Identifier: MIT
pragma solidity >=0.6.12 <0.9.0;

contract KVStorage {
    mapping(address => mapping(string => string)) private data;

    function get(address origin, string calldata key) public view returns (string memory) {
        return data[origin][key];
    }

    function set(string calldata key, string calldata value) public {
        data[msg.sender][key] = value;
    }
}
