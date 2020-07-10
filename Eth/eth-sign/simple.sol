pragma solidity >=0.4.22 <0.6.0;
contract Test {
    uint256 internal a;
    event SetA(address indexed _from, uint256 _value);
    
    function setA(uint256 _a) public {
        a = _a;
        emit SetA(msg.sender, _a);
    }
    
    function getA() public view returns (uint256) {
        return a;
    }
}