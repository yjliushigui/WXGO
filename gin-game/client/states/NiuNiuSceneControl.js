var NiuNiuSceneControl = function () {
    ConfigLayout.call(this);
}

NiuNiuSceneControl.prototype = _.create(ConfigLayout.prototype, {
    "constructor": NiuNiuSceneControl,
    "_super": ConfigLayout.prototype,


    init: function () {
        // load a config file
        this.addLayoutFile("NiuNiuScene");
        this._super.init();
    },
    preload: function () {
        this._super.preload();
    },
    create: function () {
        this._super.create();

        this.get("开始按钮").alignIn(this.get("背景"), Phaser.CENTER, 0, 152);
        this.get("等待提示").alignIn(this.get("背景"), Phaser.CENTER, 0, 88);
        this.get("邀请好友按钮").alignIn(this.get("背景"), Phaser.CENTER, 20);
        this.get("牛牛标题").alignIn(this.get("背景"), Phaser.CENTER, 0, -162);
        this.get("牛牛顶栏").alignIn(this.get("背景"), Phaser.TOP_CENTER);
        this.get("牛牛房间号").alignIn(this.get("背景"), Phaser.TOP_CENTER, 0, -13);
        this.get("牛牛局数").alignIn(this.get("背景"), Phaser.TOP_CENTER, 230, -13);    
        
        this.faceDownPokers("自己的扑克区");
        this.faceDownPokers("下左扑克区");
        this.faceDownPokers("下右扑克区");
        this.faceDownPokers("上左扑克区");
        this.faceDownPokers("上中扑克区");
        this.faceDownPokers("上右扑克区");
    },

    faceDownPokers: function(key) {
        var selfPoker = this.get(key);
        selfPoker.forEach(function(e){
            e.frame = PML.card.backFrame();
        });
    },

});
