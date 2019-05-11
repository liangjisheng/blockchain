pragma solidity ^0.4.0;

contract Calc{
  /*区块链存储*/
  uint count;

  /*执行会写入数据，所以需要`transaction`的方式执行。*/
  function add(uint a, uint b) public returns(uint){
    count++;
    return a + b;
  }

  /*执行不会写入数据，所以允许`call`的方式执行。*/
  function getCount() constant public returns (uint){
    return count;
  }
}