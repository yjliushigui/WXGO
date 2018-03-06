// `Phaser Mobile Layout` global domain
var PML = {};

(function () {
    // "window.screen.width/height 比 window.innerWidth/innerHeight(不包含通知栏) 更准确
    // var w = window.innerWidth * window.devicePixelRatio;
    // var h = window.innerHeight * window.devicePixelRatio;
    var w = 1280, h = 720;

    var game = new Phaser.Game(w, h, Phaser.CANVAS, "game");
    Boot = function () { }

    Boot.prototype = {
        init: function () {
            // 保存到全局域中
            PML.game = game;
        },
        preload: function () {
            game.load.script("Network", "scripts/Network.js");
            game.load.script("Shape", "scripts/Shape.js");
            game.load.script("UICreator", "scripts/UICreator.js");
            game.load.script("ConfigLayout", "scripts/ConfigLayout.js");
            game.load.script("CardDef", "scripts/CardDef.js");

            // load layout config file
            game.load.text("LoginScene", "/static/assets/layout/LoginScene.yaml");
            game.load.text("HallScene", "/static/assets/layout/HallScene.yaml");
            game.load.text("NiuNiuScene", "/static/assets/layout/NiuNiuScene.yaml");
            // inherit ConfigLayout
            game.load.script("LoginSceneControl", "states/LoginSceneControl.js");
            game.load.script("HallSceneControl", "states/HallSceneControl.js");
            game.load.script("NiuNiuSceneControl", "states/NiuNiuSceneControl.js");
            game.load.script("Preload", "states/Preload.js");
        },

        create: function () {
            PML.net = new Network();
            PML.card = new CardDef();
            
            game.input.maxPointers = 1;
            game.scale.pageAlignVertically = true;
            game.scale.pageAlignHorizontally = true;
            // game.scale.scaleMode = Phaser.ScaleManager.SHOW_ALL;
            // 强制使用横屏游戏
            // game.scale.forceOrientation(true, false);
            game.scale.updateLayout(true);
            game.scale.setShowAll();
            game.scale.refresh();

            // *important: calc scale ratio by resource raw coordinate
            // this.calcCoordinate(900, 468);
            this.calcCoordinate(w, h);

            PML.game.state.add("Preload", Preload);
            PML.game.state.start("Preload");
        },

        calcCoordinate: function (rawWidth, rawHeight) {
            PML.coordinate = {};
            // 计算出调整的比率
            PML.coordinate.horizontal = this.game.width / rawWidth;
            PML.coordinate.vertical = this.game.height / (rawHeight * PML.coordinate.horizontal) * PML.coordinate.horizontal;
        },
    };

    function startGame() {
        console.log("Boot state");
        game.state.add("Boot", Boot);
        game.state.start("Boot");
    }

    if (true || game.device.desktop) {
        console.log("PC");
        startGame();
    } else {
        console.log("Mobile");
        document.addEventListener("deviceready", startGame, false);
    }


}());

