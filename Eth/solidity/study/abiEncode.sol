pragma solidity ^0.4.24;

// abi.encode(...) returns (bytes)：计算参数的ABI编码。
// abi.encodePacked(...) returns (bytes)：计算参数的紧密打包编码
// abi. encodeWithSelector(bytes4 selector, ...) returns (bytes)： 计算函数选择器和参数的ABI编码
// abi.encodeWithSignature(string signature, ...) returns (bytes): 等价于* abi.encodeWithSelector(bytes4(keccak256(signature), ...)

// 通过 ABI 编码函数可以在不用调用函数的情况下，获得 ABI 编码值
contract testABI {
    function abiEncode() public constant returns (bytes) {
        abi.encode(1);  // 计算 1 的ABI编码
        //计算函数set(uint256) 及参数1 的ABI 编码
        return abi.encodeWithSignature("set(uint256)", 1); 
    }
}