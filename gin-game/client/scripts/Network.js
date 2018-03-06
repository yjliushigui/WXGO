Network = function () { };



Network.prototype = {
    ServerUrl: "ws://127.0.0.1:15000/ws",
    _handlers: {},

    // The connection is not yet open.
    CONNECTING: 0,
    // The connection is open and ready to communicate.
    OPEN: 1,
    // The connection is in the process of closing.
    CLOSING: 2,
    // The connection is closed or couldn't be opened.
    CLOSED: 3,

    Connect: function (cbOk, cbFail) {

        if (!_.isUndefined(this._ws) && this._ws.readyState != this.CLOSED) {
            return;
        }

        this._ws = new WebSocket(this.ServerUrl);
        // this._ws.binaryType = "arraybuffer";
        this._ws.onopen = this.OnConnect;
        this._ws.onerror = this.OnError;
        this._ws.onclose = this.OnDisconnect;
        this._ws.cbOk = cbOk;
        this._ws.cbFail = cbFail;
        this._ws.onmessage = this.OnMessage;
        // console.log(this._ws);
    },
    RegHandler: function (name, handler, ctx) {
        this._handlers[name] = { h: handler, c: ctx };
        console.log("注册消息处理【", name, "】");
    },
    OnConnect: function () {
        console.log("服务器连接成功");
        if (!_.isFunction(this.cbOk)) {
            return;
        }

        this.cbFail = undefined;
        this.cbOk();
    },
    OnError: function (event) {
        // console.log("error: ", event.type);
    },
    OnDisconnect: function () {
        console.log("已经同服务器断开", this);
        if (_.isFunction(this.cbFail)) {
            this.cbFail();
        }
        this.cbFail = undefined;
        this.cbOk = undefined;
    },
    OnMessage: function (event) {
        console.log("recv: ", event.data);

        var ob = JSON.parse(event.data);
        _.forIn(ob, function (val, key) {
            var handler = PML.net._handlers[key];
            if (!_.isFunction(handler.h)) {
                return;
            }

            try {
                // console.log(handler.h);
                if (_.isUndefined(handler.c)) {
                handler.h(val);
                } else {
                    handler.h.apply(handler.c, [val]);  
                }                
            } catch (e) {
                console.error(e);
            }
        });
    },
    Send: function (name, m) {
        var ob = {};
        ob[name] = m;
        var txt = JSON.stringify(ob);
        this._ws.send(txt);
        console.log("send: ", txt);
    },
};
