pragma solidity ^0.4.11;

// 修改器是一种可被继承合约属性，同时还可被继承的合约重写(override)

contract owned {
    function owned() public { owner = msg.sender; }
    address owner;

    // 定义了一个函数修饰器，可被继承，修饰时，函数体被插入到 "_" 处
    // 不符合条件时，将抛出异常
    // onlyOwner就是定义的一个函数修改器，当用这个修改器区修饰一个函数时，
    // 则函数必须满足onlyOwner的条件才能运行，这里的条件是：
    // 必须是合约的创建这才能调用函数，否则抛出异常
    modifier onlyOwner {
        require(msg.sender == owner);
        _;
    }
}

contract mortal is owned {
    // 使用继承的onlyOwner
    function close() public onlyOwner {
        selfdestruct(owner);
    }
}

contract priced {
    // 函数修饰器可接收参数
    modifier costs(uint price) {
        if (msg.value >= price) {
            _;
        }
    }
}

contract Register is priced, owned {
    mapping (address => bool) registeredAddresses;
    uint price;

    function Register(uint initialPrice) public { price = initialPrice; }

    // 需要提供paybale, 以接受以太币
    function register() public payable costs(price) {
        registeredAddresses[msg.sender] = true;
    }

    function changePrice(uint _price) public onlyOwner {
        price = _price;
    }
}
