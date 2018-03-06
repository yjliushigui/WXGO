// 定义扑克牌从网络协议转换为本地协议
// 图片资源顺序 "方块", "梅花", "红桃", "黑桃"

var CardDef = function() {};

CardDef.prototype = {
  NIU_NUMBERS: [
    9,
    10,
    11,
    12,
    13,
    14,
    15,
    16,
    17,
    18,
    0,
    1,
    2,
    3,
    4,
    5,
    6,
    7,
    8
  ],
  _PokerColor: ["方块", "梅花", "红桃", "黑桃"],
  _PokerPoint: [
    "A",
    "2",
    "3",
    "4",
    "5",
    "6",
    "7",
    "8",
    "9",
    "10",
    "J",
    "Q",
    "K"
  ],

  backFrame: function() {
    return this._PokerPoint.length * this._PokerColor.length;
  }
};
