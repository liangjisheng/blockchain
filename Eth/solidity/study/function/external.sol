pragma solidity ^0.4.11;

// 外部(external)函数由地址和函数方法签名两部分组成，可作为外部函数调用的
// 参数，或返回值

contract Oracle {
    struct Request {
        bytes data;
        function(bytes memory) external callback;
    }

    Request[] requests;
    event NewRequest(uint);

    function query(bytes data, function(bytes memory) external callback)
        public {
        requests.push(Request(data, callback));
        NewRequest(requests.length - 1);
    }

    function reply(uint requestID, bytes response) public {
        // Here goes the check that the reply comes from a trusted source
        requests[requestID].callback(response);
    }
}

contract OracleUser {
    Oracle constant oracle = Oracle(0x1234567);  // known contract

    function buySomething() {
        oracle.query("USD", this.oracleResponse);
    }

    function oracleResponse(bytes response) public {
        require(msg.sender == address(oracle));
        // Use the data
    }
}
