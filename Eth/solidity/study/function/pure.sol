pragma solidity ^0.4.16;

// 函数可以声明为pure，表示它即不读取状态变量，也不修改状态变量
// 尽管view 和 pure 修饰符编译器并未强制要求使用，view 和 pure
//  修饰也不会带来gas 消耗的改变，但是更好的编码习惯让我们跟容易
// 发现智能合约中的错误

contract C {
    function f(uint a, uint b) public pure returns (uint) {
        return a * (b + 42);
    }
}
