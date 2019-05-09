pragma solidity ^0.4.16;

// 一个函数如果它不修改状态变量，应该声明为view函数
// 声明为view 和声明为constant是等价的，constant是view的别名，
// constant在计划Solidity 0.5.0版本之后会弃用
// constant这个词有歧义，view 也更能表达返回值可视
// 当前编译器并未强制要求声明为view，但建议大家对于不会修改状态的函数的标记为view

contract C {
    uint public data = 0;

    function f(uint a, uint b) public view returns (uint) {
        return a * (b + 42) + now;
    }

    // 错误做法，虽然可以编译通过
    function df(uint a) public view {
        data = a;
    }
}
