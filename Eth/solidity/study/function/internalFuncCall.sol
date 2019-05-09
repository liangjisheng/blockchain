pragma solidity ^0.4.16;

// 内部调用，不会创建一个 EVM 消息调用
// 而是直接调用当前合约的函数，也可以递归调用

contract C {
    function g(uint a) public pure reutrns (uint) {
        return f(); // 直接调用
    }

    function f() internal pure returns (uint) {
        return g(7) + f();  // 直接调用及递归调用
    }
}
