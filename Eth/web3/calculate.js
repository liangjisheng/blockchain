var http = require('http');
let Web3 = require('web3');
var Tx = require('ethereumjs-tx');
let web3;

if (typeof web3 != 'undefined') {
    web3 = new Web3(web3.currentProvider);
} else {
    // set the provider you want from Web3.providers
    // web3 = new Web3(new Web3.providers.HttpProvider("http://localhost:7545"));
    // web3 = new Web3(new Web3.providers.HttpProvider("https://ropsten.etherscan.io/"));
    // ropsten.infura.io/v3/2fa35fe0ec5e42cebd9f4b37f9ba29dd
    web3 = new Web3(new Web3.providers.HttpProvider("ropsten.infura.io/v3/2fa35fe0ec5e42cebd9f4b37f9ba29dd"));
    console.log("set myself provider");
}

// var version = web3.version.ethereum;
// console.log('version:' + version)

// var connected = web3.isConnected();
// if (!connected) {
//     console.log("node not connected!")
// } else {
//     console.log("node connected.");
// }

// var balanceWei = web3.eth.getBalance("0xB72582e120727D468545C33501086FA870F3498E");
// console.log(balanceWei)
// 从wei转换成ether
// var balance = web3.fromWei(balanceWei, 'ether');

// compile contract
// let source = "pragma solidity ^0.4.0;contract Calc{ /*区块链存储*/ uint count;  /*执行会写入数据，所以需要`transaction`的方式执行。*/ function add(uint a, uint b) public returns(uint){ count++; return a + b;}  /*执行不会写入数据，所以允许`call`的方式执行。*/ function getCount() constant public returns (uint){ return count;}}";
// let calcCompiled = web3.eth.compile.solidity(source);

// console.log(calcCompiled)
// console.log("ABI definition:");
// console.log(calcCompiled["info"]["abiDefinition"]);