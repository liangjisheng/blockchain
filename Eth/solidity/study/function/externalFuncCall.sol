pragma solidity ^0.4.0;

// 外部调用，会创建 EVM 消息调用
// 表达式 this.g (8); 和 c.g (2)（这里的 c 是一个合约实例）是
// 外部调用函数的方式，它会发起一个消息调用，而不是 EVM 的指令跳转
// 需要注意的是，在合约的构造器中，不能使用 this 调用函数，
// 因为当前合约还没有创建完成

contract InfoFeed {
    // info() 函数，必须使用 payable 关键字，否则不能通过 value() 来接收以太币
    function info() public payable returns (uint) { return 42; }
}

contract Consumer {
    InfoFeed feed;

    function setFeed(address addr) public {
        feed = InfoFeed(addr)
    }

    function callFeed() public {
        // 当调用其它合约的函数时，可以通过选项.value()，和.gas()
        // 来分别指定要发送的以太币（以 wei 为单位）和 gas 值
        // 注意 feed.info.value(10).gas(800) 仅仅是对发送的以太币
        // 和 gas 值进行了设置，真正的调用是后面的括号 ()
        feed.info.value(10).gas(800)();
    }
}
