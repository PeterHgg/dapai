package mahjong

import (
	"math/rand"
	"time"
)

// CardTypes 长沙麻将牌型：万、条、饼 (只有 1-9)
// 10-18: 1-9万, 20-28: 1-9条, 30-38: 1-9饼
type Card int

// CSGameLogic 长沙麻将游戏逻辑
type CSGameLogic struct {
	Cards       []Card
	PlayersHand map[string][]Card // 玩家 ID -> 手牌
	Status      int               // 0: 等待, 1: 起手胡判定, 2: 轮转中
}

func NewCSGameLogic() *CSGameLogic {
	return &CSGameLogic{
		Cards:       make([]Card, 0),
		PlayersHand: make(map[string][]Card),
	}
}

// InitCards 初始化长沙麻将 108 张牌
func (g *CSGameLogic) InitCards() {
	g.Cards = make([]Card, 0)
	types := []int{1, 2, 3} // 万条饼
	for _, t := range types {
		for v := 1; v <= 9; v++ {
			card := Card(t*10 + v)
			for i := 0; i < 4; i++ { // 每种 4 张
				g.Cards = append(g.Cards, card)
			}
		}
	}
}

// Shuffle 洗牌
func (g *CSGameLogic) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(g.Cards), func(i, j int) {
		g.Cards[i], g.Cards[j] = g.Cards[j], g.Cards[i]
	})
}

// CheckQiShouHu 检查起手胡 (板板胡、四喜、六六顺、三同)
func (g *CSGameLogic) CheckQiShouHu(hand []Card) []string {
	results := make([]string, 0)
	counts := make(map[Card]int)
	for _, c := range hand {
		counts[c]++
	}

	hasJiang := false
	for _, count := range counts {
		if count == 4 {
			results = append(results, "四喜")
		}
		if count == 3 {
			// 六六顺检查通常需要两组刻子
		}
		// 这里后续扩展详细逻辑
	}

	if !hasJiang {
		// 检查板板胡 (没有 2, 5, 8)
		isBanBan := true
		for _, c := range hand {
			val := int(c) % 10
			if val == 2 || val == 5 || val == 8 {
				isBanBan = false
				break
			}
		}
		if isBanBan {
			results = append(results, "板板胡")
		}
	}

	return results
}
