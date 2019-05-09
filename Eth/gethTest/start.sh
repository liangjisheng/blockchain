#!/bin/bash

# https://learnblockchain.cn/2017/11/24/init-env/

./geth --datadir testNet --dev console 2>> test.log

# eth.accounts    // 查看账户
# personal.listAccounts   // 查看账户
# eth.getBalance(eth.accounts[0]) // 看一下账户里的余额
# personal.newAccount("ljs")  // 创建账户
# eth.getBalance(eth.accounts[1])
# eth.sendTransaction({from: '0x3298bce87334544a7f5a14b90413ae9deb4b7cd0', to: '0xbb02b04bf65a7a63cffd174de349c7202fa3ac29', value: web3.toWei(1, "ether")})
# personal.unlockAccount(eth.accounts[1],"ljs");  // 解锁账户