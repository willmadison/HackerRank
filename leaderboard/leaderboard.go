package leaderboard

type scoreboard struct {
	scores          []int
	rankingsByScore map[int]int
}

func newScoreboard(scores []int) *scoreboard {
	s := scoreboard{
		scores: scores,
	}
	s.rank()
	return &s
}

func (s *scoreboard) rank() {
	currentMax := s.scores[0]
	currentRank := 1
	s.rankingsByScore = map[int]int{
		currentMax: 1,
	}

	for _, score := range s.scores {
		if score < currentMax {
			currentMax = score
			currentRank++
			s.rankingsByScore[currentMax] = currentRank
		}
	}
}

func (s *scoreboard) addScore(score int) int {
	var updatedScores []int
	before, after := s.findInsertionPoint(score)
	updatedScores = append(updatedScores, before...)
	updatedScores = append(updatedScores, score)
	s.scores = append(updatedScores, after...)

	s.rank()

	return s.rankingsByScore[score]
}

func (s *scoreboard) findInsertionPoint(score int) ([]int, []int) {
	var insertionPoint int

	for insertionPoint < len(s.scores) && score < s.scores[insertionPoint] {
		insertionPoint++
	}

	return s.scores[0:insertionPoint], s.scores[insertionPoint:len(s.scores)]
}

func Climb(scores []int, aliceScores []int) []int {
	var rankings []int

	scoreboard := newScoreboard(scores)

	for _, score := range aliceScores {
		s := scoreboard.addScore(score)
		rankings = append(rankings, s)
	}

	return rankings
}
