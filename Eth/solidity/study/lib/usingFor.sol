pragma solidity ^0.4.16;

// 指令 using A for B; 用来把库函数 (从库 A) 关联到类型 B。这些函数将会把
// 调用函数的实例作为第一个参数。语法类似，python 中的 self 变量一样。例如：
// A 库有函数 add(B b1, B b2)，则使用 Using A for B 指令后，如果有 B b1 
// 就可以使用 b1.add(b2)

// 库合约代码和上一节一样
library Set {
  struct Data { mapping(uint => bool) flags; }

  function insert(Data storage self, uint value)
      public
      returns (bool)
  {
      if (self.flags[value])
        return false; // already there
      self.flags[value] = true;
      return true;
  }

  function remove(Data storage self, uint value)
      public
      returns (bool)
  {
      if (!self.flags[value])
          return false; // not there
      self.flags[value] = false;
      return true;
  }

  function contains(Data storage self, uint value)
      public
      view
      returns (bool)
  {
      return self.flags[value];
  }
}

contract C {
    using Set for Set.Data; // 这是一个关键的变化
    Set.Data knownValues;

    function register(uint value) public {
        // 现在 Set.Data都对应的成员方法
        // 效果和Set.insert(knownValues, value)相同
        require(knownValues.insert(value));
    }
}