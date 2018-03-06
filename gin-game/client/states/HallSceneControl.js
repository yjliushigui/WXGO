var HallSceneControl = function() {
  ConfigLayout.call(this);
};

HallSceneControl.prototype = _.create(ConfigLayout.prototype, {
  constructor: HallSceneControl,
  _super: ConfigLayout.prototype,

  init: function() {
    // load a config file
    this.addLayoutFile("HallScene");
    this._super.init();
  },
  preload: function() {
    this._super.preload();
  },
  create: function() {
    this._super.create();

    PML.net.RegHandler("S2C_CreateRoom", this.onCreateRoom, this);

    var btn = this.get("创建房间");
    btn.onInputUp.add(this.createRoom, this);
  },
  createRoom: function() {
    PML.net.Send("C2S_CreateRoom", {});
  },
  onCreateRoom: function(m) {
    if (_.isUndefined(m)) {
      console.log("login:", m);
      return;
    }

    if (m.Code != 0) {
      console.log("login failed, code", m.Code);
      return;
    }

    console.log("创建房间：", m);

    PML.game.state.add("NiuNiuSceneControl", NiuNiuSceneControl);
    PML.game.state.start("NiuNiuSceneControl");
  }
});
