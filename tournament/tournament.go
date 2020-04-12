package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

const (
	LOSS_PTS = iota
	DRAW_PTS = iota
	WIN_PTS  = iota + 1
)

type Team struct {
	MP, W, D, L, P int
}

func (t *Team) Win() {
	t.W++
	t.MP++
	t.P += WIN_PTS
}

func (t *Team) Loose() {
	t.L++
	t.MP++
}

func (t *Team) Draw() {
	t.D++
	t.MP++
	t.P += DRAW_PTS
}

func readData(scanner *bufio.Scanner, scoreboard map[string]Team) error {
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		splitted := strings.Split(line, ";")

		if len(splitted) != 3 {
			return errors.New("Line must contain three elements.")
		}

		away := scoreboard[splitted[0]]
		home := scoreboard[splitted[1]]
		gameResult := splitted[2]

		switch gameResult {
		case "win":
			away.Win()
			home.Loose()
			break
		case "draw":
			away.Draw()
			home.Draw()
			break
		case "loss":
			away.Loose()
			home.Win()
			break
		default:
			return errors.New("Invalid game status.")
		}

		scoreboard[splitted[0]] = away
		scoreboard[splitted[1]] = home
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func writeData(writer io.Writer, scoreboard map[string]Team) error {
	if len(scoreboard) == 0 {
		return errors.New("Empty scoreboard.")
	}

	teamNames := make([]string, 0, len(scoreboard))
	for teamName := range scoreboard {
		teamNames = append(teamNames, teamName)
	}

	sort.Slice(teamNames, func(i, j int) bool {
		first, second := scoreboard[teamNames[i]], scoreboard[teamNames[j]]
		if first.P == second.P {
			return teamNames[i] < teamNames[j]
		}
		return first.P > second.P
	})

	writer.Write([]byte(fmt.Sprintf("Team                           | MP |  W |  D |  L |  P\n")))
	for _, teamName := range teamNames {
		team := scoreboard[teamName]
		writer.Write([]byte(fmt.Sprintf("%-30s | %2d | %2d | %2d | %2d | %2d\n", teamName, team.MP, team.W, team.D, team.L, team.P)))
	}

	return nil
}

func Tally(reader io.Reader, writer io.Writer) (err error) {
	scoreboard := make(map[string]Team)
	scanner := bufio.NewScanner(reader)

	err = readData(scanner, scoreboard)
	if err != nil {
		return err
	}

	err = writeData(writer, scoreboard)
	if err != nil {
		return err
	}

	return nil
}
