package game

import (
	"gin-game/server/msg"
)

func (c *Game) handleFirstDraw(playerID int, data []byte) {
	// c.syscard.Init()
	c.syscard.Shuffle()

	for _, playerIDn := range c.players {
		playerCard := c.playerCard[playerIDn]
		playerCard.Init(playerIDn, c.syscard.Deal(5))

		msg.WriteMsg(playerIDn, &msg.S2C_StartGame{
			Dealer:   c.dealer,
			HandCard: playerCard.getHandCards(),
			Round:    c.nRound,
		})
	}
}
