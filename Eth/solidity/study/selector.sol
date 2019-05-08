pragma solidity ^0.4.16;

// public (或 external) 函数有一个特殊的成员 selector, 它对应一个 ABI 函数选择器

contract Selector {
    function f() public view returns (bytes4) {
        return this.f.selector;
    }
}
