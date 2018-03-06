package msg

// C2S_Register 注册
type C2S_Register struct {
	OpenID      string
	AccessToken string
	Version     int
}

type S2C_Register struct {
	// success     = 0
	// failed      = 1
	// notRegister = 2
	// ban         = 3
	// versionLow  = 4
	Code int
	// 注册成功后得到的玩家ID
	PlayerID int
}

// C2S_Login 请求登录
type C2S_Login struct {
	OpenID      string
	AccessToken string
	Version     int
}

type S2C_Login struct {
	// success     = 0
	// failed      = 1
	// notRegister = 2
	// ban         = 3
	// versionLow  = 4
	Code     int
	PlayerID int
}

// C2S_PlayerInfo 请求用户信息
type C2S_PlayerInfo struct {
}

// S2C_PlayerInfo 用户信息
type S2C_PlayerInfo struct {
	PlayerID int
	Avatar   string
	Nickname string
	// 1-男 2-女
	Sex int

	// 房卡数量
	RoomCard int

	// 0-没在房间中, >0 房间ID
	RoomKey int
}

type C2S_CreateRoom struct {
}

// S2C_CreateRoom 创建房间
type S2C_CreateRoom struct {
	// 成功6位数房间号，失败返回-0
	RoomKey int
	// 房间类型：
	Type int
	// 0-成功 1-房卡数量不足
	Code int
}

// S2C_DingHao 通知玩家顶号
type S2C_DingHao struct {
}

// S2C_Offline 通知玩家掉线
type S2C_Offline struct {
	PlayerID int
}

// S2C_Online 通知玩家上线
type S2C_Online struct {
	PlayerID int
}

// S2C_StartGame 开始游戏（发送手牌）
type S2C_StartGame struct {
	// 庄家ID
	Dealer int
	// 玩家手牌
	HandCard []int
	// 当前局数
	Round int
}

// C2S_StartGame 开始游戏
type C2S_StartGame struct {
}
