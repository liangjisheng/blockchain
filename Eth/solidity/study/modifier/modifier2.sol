pragma solidity ^0.4.11;

contract modifysample {
    uint a = 10;

    modifier mf1 (uint b) {
        uint c = b;
        _;
        c = a;
        a = 11;
    }

    modifier mf2 () {
        uint c = a;
        _;
    }

    modifier mf3() {
        a = 12;
        return;
        _;
        a = 13;
    }

    // test1扩展之后是这样的
    // uint c = b;
    //     uint c = a;
    //         a = 12;
    //         return;  // 这里会返回，所以下面的两行不执行
    //         _;
    //         a = 13;
    // c = a;
    // a = 11;
    // test1()运行后，状态变量a的值为11
    function test1() mf1(a) mf2 mf3 public {
        a = 1;
    }

    function test2() public constant returns (uint) {
        return a;
    }
}
