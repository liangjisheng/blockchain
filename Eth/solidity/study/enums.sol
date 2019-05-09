pragma solidity ^0.4.0;

// 枚举可以用来自定义类型。它可以显示的与整数进行转换，但不能进行隐式转换

contract test {
    enum ActionChoices { GoLeft, GoRight, GoStraight, SitStill }
    ActionChoices choice;
    ActionChoices constant defaultChoice = ActionChoices.GoStraight;

    function setGoStraight() {
        choice = ActionChoices.GoStraight;
    }

    function getChoice() constant returns (ActionChoices) {
        return choice;
    }

    function getDefaultChoice() constant returns (uint) {
        return uint(defaultChoice);
    }
}
