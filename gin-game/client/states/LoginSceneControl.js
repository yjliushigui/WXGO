var LoginSceneControl = function () {
    ConfigLayout.call(this);
}

LoginSceneControl.prototype = _.create(ConfigLayout.prototype, {
    "constructor": LoginSceneControl,
    "_super": ConfigLayout.prototype,


    init: function () {
        // load a config file
        this.addLayoutFile("LoginScene");
        this._super.init();
    },
    preload: function () {
        this._super.preload();
    },
    create: function () {
        this._super.create();

        PML.net.RegHandler("S2C_Login", this.onLogin);
        var btn = this.get("微信登录");
        btn.onInputUp.add(this.onWeiXin);
    },
    onWeiXin: function() {
        PML.net.Connect(function() {
            // 发送登录请求
            PML.net.Send("C2S_Login", { OpenID: "111", AccessToken: "###", Version: 1 });
        }, function() {
            console.log("fail");
        });
    },
    onLogin: function (m) {
        if (_.isUndefined(m)) {
            console.log("login:", m);
            return;
        }

        if (m.Code != 0) {
            console.log("login failed, code", m.Code);
            return;
        }

        console.log("登录成功！玩家ID", m.PlayerID);
        PML.game.state.add("HallSceneControl", HallSceneControl);
        PML.game.state.start("HallSceneControl");
    }

});
