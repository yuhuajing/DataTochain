// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.8.2 <0.9.0;

contract UserPassStore {
    mapping(string => UserPassInfo) public tracedata;

    struct UserPassInfo {
        string username;
        string passwordhash;
        string name;
        string phone;
        string description;
    }

    constructor() {}

    function addorupdateData(
        string memory username,
        string memory passwordhash,
        string memory name,
        string memory phone,
        string memory description
    ) external {
        tracedata[username].username = username;
        tracedata[username].passwordhash = passwordhash;
        tracedata[username].name = name;
        tracedata[username].phone = phone;
        tracedata[username].description = description;
    }

    function tracedataByName(string memory name)
        external
        view
        returns (UserPassInfo memory)
    {
        return tracedata[name];
    }
}
