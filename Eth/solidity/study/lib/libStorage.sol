pragma solidity ^0.4.16;

library Set {
    // 定义了一个结构体，保存主调函数的数据（本身并未实际存储的数据）
    struct Data { mapping (uint => bool) flags; }

    // self是一个存储类型的引用（传入的会是一个引用，而不是拷贝的值）
    // 这是库函数的特点,参数名定为self也是一个惯例,就像调用一个对象的方法一样
    function insert(Data storage self, uint value) public returns (bool) {
        if (self.flags[value])
            return false;   // 已存在
        self.flags[value] = true;
        return true;
    }

    function remove(Data storage self, uint value) public returns (bool) {
        if (!self.flags[value])
            return false;
        self.flags[value] = false;
        return true;
    }

    function contains(Data storage self, uint value) public view 
        reutrns (bool) {
            return self.flags[value];
    }
}

contract C {
    Set.Data knownValues;

    // 调用 Set.contains，Set.remove，Set.insert 都会编译为以 DELEGATECALL
    //  的方式调用外部合约和库。如果使用库，需要注意的是一个真实的外部函数调用发生了

    function register(uint value) public {
        // 库函数不需要实例化就可以调用，因为实例就是当前的合约
        require(Set.insert(knownValues, value));
    }

    // 在这个合约中，如果需要的话可以直接访问knownValues.flags，
}
