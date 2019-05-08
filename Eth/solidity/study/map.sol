pragma solidity ^0.4.0;

contract MappingExample {
    // 可以通过将映射标记为public，来让Solidity创建一个访问器
    // 通过提供一个键值做为参数来访问它，将返回对应的值    
    mapping(address => uint) public balances;

    function update(uint newBalance) public {
        balances[msg.sender] = newBalance;
    }
}

contract MappingUser {
    function f() public returns (uint) {
        MappingExample m = new MappingExample();
        m.update(100);
        // this 表示当前合约，可以显式的转换为Address
        return m.balances(this);
    }
}
