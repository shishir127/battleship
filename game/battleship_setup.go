package game

type BattleshipSetup struct {
	GridSize             int
	NumberOfShips        int
	Player1ShipPositions []*Location
	Player2ShipPositions []*Location
	NumberOfMissiles     int
	Player1sMissiles     []*Missile
	Player2sMissiles     []*Missile
}
