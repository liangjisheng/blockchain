#!/bin/bash

eth.accounts
eth.getBalance(eth.accounts[0])

personal.newAccount("ljs")  # 0xbb02b04bf65a7a63cffd174de349c7202fa3ac29
eth.sendTransaction({from: '0x3298bce87334544a7f5a14b90413ae9deb4b7cd0', to: '0xbb02b04bf65a7a63cffd174de349c7202fa3ac29', value: web3.toWei(1, "ether")})
eth.getBalance(eth.accounts[1])
personal.unlockAccount(eth.accounts[1],"ljs");

personal.newAccount("ljs1") # 0x4aa03cbee2eb575b8fd57537816e8c179d79eeaa
eth.sendTransaction({from: '0x3298bce87334544a7f5a14b90413ae9deb4b7cd0', to: '0x4aa03cbee2eb575b8fd57537816e8c179d79eeaa', value: web3.toWei(1, "ether")})
eth.getTransactionReceipt("0x1f9beceb69c10b701c603a118322d5994ed13246c29cee37ecbc932680f3b208")
eth.getBalance(eth.accounts[2])
personal.unlockAccount(eth.accounts[2],"ljs1");

personal.newAccount("ljs2") # 0x408ac3d91f1b17ad0d57c19b408d2c105bca7afb
eth.sendTransaction({from: '0x3298bce87334544a7f5a14b90413ae9deb4b7cd0', to: '0x408ac3d91f1b17ad0d57c19b408d2c105bca7afb', value: web3.toWei(1, "ether")})
eth.getTransactionReceipt("0x096bb0978a6644d2a98fdee208c835d9b28b55b4f6ad5be0566e3b88b57705ce")
eth.getBalance(eth.accounts[3])
personal.unlockAccount(eth.accounts[3],"ljs2");

personal.newAccount("ljs3") # 0x8a2c828a77b024eabcc30c844ab7d36c71121c21
eth.sendTransaction({from: '0x3298bce87334544a7f5a14b90413ae9deb4b7cd0', to: '0x8a2c828a77b024eabcc30c844ab7d36c71121c21', value: web3.toWei(1, "ether")})
eth.getTransactionReceipt("0x401e0d033599faea37fce073c64652862efd6151f4bc4960406eae5160307fc3")
eth.getBalance(eth.accounts[4])
personal.unlockAccount(eth.accounts[4],"ljs3");

# contract address 
0x6fb78ca7f06b992e9197a4b8d01bb8e36a85d197
eth.sendTransaction({from: '0x3298bce87334544a7f5a14b90413ae9deb4b7cd0', to: '0x6fb78ca7f06b992e9197a4b8d01bb8e36a85d197', value: web3.toWei(1, "ether")})
eth.getTransactionReceipt("0x901c38772788dd3e5e08b6b2e6b71c0b4976cab3c6ded1b458646a1b633f2cfb")

eth.sendTransaction({from: '0x3298bce87334544a7f5a14b90413ae9deb4b7cd0', to: '0x6fb78ca7f06b992e9197a4b8d01bb8e36a85d197', value: web3.toWei(1, "ether")})
eth.getTransactionReceipt("0xd8e8dd6e28c2077b20da668ff75b826b533d7c8701f222dba45f0d7ecb4d7dfa")

multisigwallet.walletBalance()
multisigwallet.getPendingTransactions()

multisigwallet.addManager.call("0x4aa03cbee2eb575b8fd57537816e8c179d79eeaa")
multisigwallet.addManager.call("0x408ac3d91f1b17ad0d57c19b408d2c105bca7afb")
multisigwallet.addManager.call("0x8a2c828a77b024eabcc30c844ab7d36c71121c21")

multisigwallet.transferTo.call("0x3298bce87334544a7f5a14b90413ae9deb4b7cd0", web3.toWei(0.5, "ether"))

multisigwallet.signTransaction(0)
