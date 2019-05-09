pragma solidity ^0.4.16;

// solidity 还可以返回任意数量的参数作为输出

contract Simple {
    function arithmetics(uint _a, uint _b)
        public
        pure
        returns (uint o_sum, uint o_product)
    {
        o_sum = _a + _b;
        o_product = _a * _b;
    }

    function f() public pure returns (uint, bool, uint) {
        // 使用元组返回多个值
        return (7, true, 2);
    }

    function callf() public {
        uint x;
        bool y;
        uint z;
        // 使用元组给多个变量赋值
        (x, y , z)  = f();
    }
}
