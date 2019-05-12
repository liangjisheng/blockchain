pragma solidity ^0.5.8;

contract Store {
  string public version;
  constructor(string memory _version) public {
    version = _version;
  }

  event ItemSet(bytes32 key, bytes32 value);
  mapping (bytes32 => bytes32) public items;

  function setItem(bytes32 key, bytes32 value) external {
    items[key] = value;
    emit ItemSet(key, value);
  }
}