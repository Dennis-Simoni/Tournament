package Tournament

import (
	"encoding/csv"
	"fmt"
	"io"
	"sort"
	"strings"
	"text/tabwriter"
)

// Team is a football team in the League
type Team struct {
	Name           string
	MP, W, D, L, P int
}

// League represents the score board
type League struct {
	teams []*Team
}

// ParseData reads input data and delegates mapping to HandleResults
func (l *League) ParseData(data [][]string) error {
	for i := 0; i < len(data); i++ {
		split := strings.Split(data[i][0], ";")
		if len(split) != 3 {
			return fmt.Errorf("wrong number of fields")
		}
		err := l.HandleResults(split[2], split[0], split[1])
		if err != nil {
			return fmt.Errorf("could not handle result %v", err)
		}
	}
	return nil
}

// HandleResults updates scores for each team by creating or finding them.
func (l *League) HandleResults(res string, name1, name2 string) error {
	t := l.FindOrCreateTeam(name1)
	t2 := l.FindOrCreateTeam(name2)

	switch res {
	case "win":
		t.MP, t.W, t.P = t.MP+1, t.W+1, t.P+3
		t2.MP, t2.L = t2.MP+1, t2.L+1
	case "draw":
		t.MP, t.D, t.P = t.MP+1, t.D+1, t.P+1
		t2.MP, t2.D, t2.P = t2.MP+1, t2.D+1, t2.P+1
	case "loss":
		t.MP, t.L = t.MP+1, t.L+1
		t2.MP, t2.W, t2.P = t2.MP+1, t2.W+1, t2.P+3
	default:
		return fmt.Errorf("invalid match outcome %v", res)
	}
	return nil
}

// FindOrCreateTeam returns a team if is found otherwise creates one.
func (l *League) FindOrCreateTeam(name string) *Team {
	for _, v := range l.teams {
		if v.Name == name {
			return v
		}
	}

	nt := &Team{Name: name}
	l.AddTeam(nt)
	return nt
}

//AddTeam adds a team on the team list.
func (l *League) AddTeam(t *Team) {
	l.teams = append(l.teams, t)
}

// BuildTable is responsible for building the output.
func (l *League) BuildTable(w io.Writer) {
	// check for tie.
	l.CheckTie()

	tw := tabwriter.NewWriter(w, 31, 0, 0, ' ', tabwriter.TabIndent)
	fmt.Fprintln(tw, "Team\t| MP |  W |  D |  L |  P")

	for _, v := range l.teams {
		row := fmt.Sprintf("%s\t|  %d |  %d |  %d |  %d |  %d", v.Name, v.MP, v.W, v.D, v.L, v.P)
		fmt.Fprintln(tw, row)
	}

	tw.Flush()
}

// SortLeagueByPoints arranges the Teams in highest points order.
func (l *League) SortLeagueByPoints() {
	sort.Slice(l.teams, func(i, j int) bool {
		return l.teams[i].P > l.teams[j].P
	})
}

func (l *League) SortByName(team1, team2 *Team) {
	l.SortLeagueByPoints()
	sort.Slice(l.teams, func(i, j int) bool {
		return team1.Name > team2.Name
	})
}

func (l *League) CheckTie() {
	for i := 0; i < len(l.teams) -1; i++ {
		if l.teams[i].P == l.teams[i+1].P{
			l.SortByName(l.teams[i], l.teams[i + 1])
		}
	}
}

// Tally is reading the data input, converting it to a meaningful structure and producing an output
func Tally(in io.Reader, w io.Writer) error {
	league := &League{teams: []*Team{}}
	data, err := readData(in)

	if err != nil {
		return err
	}

	err = league.ParseData(data)
	if err != nil {
		return err
	}

	league.SortLeagueByPoints()
	league.BuildTable(w)
	return nil
}

func readData(input io.Reader) ([][]string, error) {
	reader := csv.NewReader(input)
	reader.Comment = '#'
	res, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return res, nil
}