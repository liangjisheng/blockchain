pragma solidity ^0.4.0;

// require 来检查输入条件，以及 assert 用于内部错误检查

contract Sharer {
    function sendHalf(address addr) public payable returns(uint balance) {
        require(msg.value % 2 == 0);    // 仅允许偶数
        uint balanceBeforeTransfer = this.balance;
        addr.transfer(msg.value / 2);   // 如果失败，抛出异常，下面的代码不执行
        assert(this.balance == balanceBeforeTransfer - msg.value / 2);
        return this.balance;
    }
}
