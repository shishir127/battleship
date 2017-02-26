package parser

type BattleShipParser struct {
	input []string
}

func NewBattleShipParser(stringInput []string) *BattleShipParser {
	return &BattleShipParser{input: stringInput}
}

func (parser *BattleShipParser) Parse() error {
	return nil
}
