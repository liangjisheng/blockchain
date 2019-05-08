pragma solidity ^0.4.0;

contract AddrTest {
    event logdata(bytes data);
    function() payable {
        logdata(msg.data);
    }

    function getBalance() returns (uint) {
        return this.balance;
    }

    uint score = 0;
    function setScore(uint s) public {
        score = s;
    }

    function getScore() returns (uint) {
        return score;
    }
}

contract CallTest {
    function deposit() payable {

    }

    event logSendEvent(address to, uint value);
    function transferEther(address towho) payable {
        towho.transfer(10);
        logSendEvent(towho, 10);
    }

    function callNoFunc(address addr) returns (bool) {
        return addr.call("tinyxiong", 1234);
    }

    function callfunc(address addr) returns (bool) {
        bytes4 methodId = bytes4(keccak256("setScore(uint256)"));
        return addr.call(methodId, 100);
    }

    function getBalance() returns (uint) {
        return this.balance;
    }
}
