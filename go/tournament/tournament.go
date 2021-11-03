package tournament

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"
)

// TournamentResults holds an array of Teams and the table for the tournament
type TournamentResults struct {
	Table string
	Teams []Team
}

// Team is the data structure for a team containing various statistics.
type Team struct {
	Name          string
	MacthesPlayed int
	Wins          int
	Draws         int
	Losses        int
	Points        int
}

func Tally(r io.Reader, w io.Writer) error {
	// Read all data from the io.Reader
	byteData, err := ioutil.ReadAll(r)

	if err != nil {
		return err
	}

	// Convert the data into a string to work with.
	data := string(byteData)

	// Sanitized data removign comments
	re := regexp.MustCompile(`(?m)^#.*`)
	data = re.ReplaceAllString(data, "")

	// Trim any newlines from the top and bottom of the input
	// then split each line into a slice entry.
	matches := strings.Split(strings.Trim(data, "\n"), "\n")

	// Create a new instance of TournamentResults with an header for the table and
	// a slice of teams.
	tr := TournamentResults{
		Table: fmt.Sprintf("%-31s|%3s |%3s |%3s |%3s |%3s\n", "Team", "MP", "W", "D", "L", "P"),
		Teams: []Team{},
	}

	// Used to ensure the team entry is valid.
	re = regexp.MustCompile(`^[A-Z][a-z].+? [A-Z][a-z].+`)

	// For each match
	for _, v := range matches {
		// if match is a newline or empty line then skip it.
		if v == "" {
			continue
		}

		// For each entry (seperated by ;)
		matchDetails := strings.Split(v, ";")
		for i, v := range matchDetails {

			// index 0, 1 will always be a team name while index 2
			// will contain the match result. The match result will
			// determine how to set the statistics for each team
			switch i {
			case 0, 1:
				// if the team entry doesn't match the regex then report an error.
				if !re.MatchString(v) {
					return errors.New("invalid entry")
				}

				// if a team doesn't exist in the slice then add the team
				// to the tournament results
				if !tr.TeamPlayed(v) {
					tr.Teams = append(tr.Teams, Team{Name: v})
				}

				// get the index of the team by name
				ti, err := tr.TeamIndex(v)
				if err != nil {
					return err
				}

				// increase the matches played for the specified team
				tr.Teams[ti].MacthesPlayed++
			case 2:
				// get the index of the first team in the matchDetails
				homeTeamIndex, err := tr.TeamIndex(matchDetails[i-2])
				if err != nil {
					return err
				}

				// get the index of the second team in the matchDetails
				opposingTeamIndex, err := tr.TeamIndex(matchDetails[i-1])
				if err != nil {
					return err
				}

				// increment the stats of the teams that played in the current match
				switch v {
				case "win":
					tr.Teams[homeTeamIndex].Wins++
					tr.Teams[opposingTeamIndex].Losses++
				case "loss":
					tr.Teams[homeTeamIndex].Losses++
					tr.Teams[opposingTeamIndex].Wins++
				case "draw":
					tr.Teams[homeTeamIndex].Draws++
					tr.Teams[opposingTeamIndex].Draws++
				default:
					return errors.New("invalid entry")
				}
			}
		}
	}

	// tally the total score for each team that played
	for i := range tr.Teams {
		tr.Teams[i].Points = tr.Teams[i].Wins*3 + tr.Teams[i].Draws
	}

	// Sort the teams in desending order by points
	// if two teams have the same points they are sorted by name
	sort.Slice(tr.Teams, func(i, j int) bool {
		if tr.Teams[i].Points != tr.Teams[j].Points {
			return tr.Teams[i].Points > tr.Teams[j].Points
		} else {
			return tr.Teams[i].Name < tr.Teams[j].Name
		}
	})

	// Creat the table containing the tournament results.
	for _, v := range tr.Teams {
		tr.Table += fmt.Sprintf("%-31s|%3d |%3d |%3d |%3d |%3d\n", v.Name, v.MacthesPlayed, v.Wins, v.Draws, v.Losses, v.Points)
	}

	// write the table to the io.Writer.
	w.Write([]byte(tr.Table))

	return nil
}

// TeamPlayed searches for a team by name and returns whether it exists or not
// in the slice of teams of the tournament.
func (t *TournamentResults) TeamPlayed(s string) bool {
	for _, v := range t.Teams {
		if v.Name == s {
			return true
		}
	}

	return false
}

// TeamIndex searchs for a team by name and returns the index of the team.
func (t *TournamentResults) TeamIndex(s string) (int, error) {
	for i, v := range t.Teams {
		if s == v.Name {
			return i, nil
		}
	}

	return 0, errors.New("no team found matching search string")
}
