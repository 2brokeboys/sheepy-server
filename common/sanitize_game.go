package common

// Sanitize checks whether the struct is valid, returns a list of occured problems
func (g *Game) Sanitize() []string {
	ret := make([]string, 0)

	// Check player constellation
	if g.Player < 0 {
		ret = append(ret, "No player given.")
	}
	if g.Playmate == g.Player {
		ret = append(ret, "Player and playmate must not match.")
	}
	playmatePart := false
	playerPart := false
	for _, n := range g.Participants {
		if n == g.Player {
			playerPart = true
		}
		if n == g.Playmate {
			playmatePart = true
		}
	}
	if !playerPart {
		ret = append(ret, "The player has to participate in the game.")
	}
	if g.Playmate >= 0 && !playmatePart {
		ret = append(ret, "The playmate has to participate in the game.")
	}

	// Check point annomaly
	if g.Schwarz && g.Points != 0 && g.Points != 120 {
		ret = append(ret, "Schwarz can only occur in combination with 0 or 120 points.")
	}
	if g.Points > 120 {
		ret = append(ret, "More than 120 points is impossible.")
	}
	if g.Points < 0 {
		ret = append(ret, "Less than 0 points is impossible.")
	}

	if g.GameType != Ramsch && g.Virgins > 0 {
		ret = append(ret, "Virgins only occur in ramsch.")
	}

	switch g.GameType {
	case SauEichel:
		fallthrough
	case SauGras:
		fallthrough
	case SauSchell:
		if g.Playmate < 0 {
			ret = append(ret, "Sau has to be with playmate.")
		}
		if g.Runners > 14 {
			ret = append(ret, "More than 14 runners are impossible.")
		}
	case Wenz:
		if g.Playmate >= 0 {
			ret = append(ret, "Wenz can not be with playmate.")
		}
		if g.Runners > 4 {
			ret = append(ret, "More than 4 runners are impossible.")
		}
	case SoloEichel:
		fallthrough
	case SoloGras:
		fallthrough
	case SoloHerz:
		fallthrough
	case SoloSchell:
		if g.Playmate >= 0 {
			ret = append(ret, "Solo can not be with playmate.")
		}
		if g.Runners > 8 {
			ret = append(ret, "More than 8 runner is impossible.")
		}
	case Ramsch:
		if g.Playmate >= 0 {
			ret = append(ret, "Only one person can loose ramsch.")
		}
		if g.Virgins > 3 {
			ret = append(ret, "Not more than 3 people can be virgin.")
		}
		if g.Points < 120/(4-g.Virgins) {
			ret = append(ret, "Something is wrong with the combination of points and virgins.")
		}
	}

	return ret
}
