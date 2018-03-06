Preload = function () { };

Preload.prototype = {
    preload: function () {
        PML.game.load.script("UIConfig", "scripts/UIConfig.js");
     },
    create: function () {
        console.log("Preload");
        // and to global domain
        PML.uiConfig = new UIConfig();

        PML.game.state.add("LoginSceneControl", LoginSceneControl);
        PML.game.state.start("LoginSceneControl");
        // PML.game.state.add("HallSceneControl", HallSceneControl);
        // PML.game.state.start("HallSceneControl");
        // PML.game.state.add("NiuNiuSceneControl", NiuNiuSceneControl);
        // PML.game.state.start("NiuNiuSceneControl");
    },

    update: function () { },
};
