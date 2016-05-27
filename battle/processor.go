package battle

import (
	"bufio"
	"fmt"
	"os"

	"github.com/matthewmpalen/mochi/core"
)

type (
	Processor struct {
		playerParty *core.Party
		enemyParty  *core.Party
	}
)

func (p *Processor) ReadInput() {
	reader := bufio.NewScanner(os.Stdin)
	for _, player := range p.playerParty.Members() {
		text := reader.Text()
		fmt.Printf("Player: %s, text=%s", player, text)
	}
}
