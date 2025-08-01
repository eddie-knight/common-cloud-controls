package cmd

import (
	"fmt"
	"io/ioutil"
	"strings"

	"gopkg.in/yaml.v3"
)

var finosMemberLogos = map[string]string{
	"FINOS":       "https://www.finos.org/hubfs/FINOS/finos-logo/FINOS_Icon_Wordmark_Name_horz_White.svg",
	"Citi":        "https://www.finos.org/hs-fs/hubfs/2-Jan-18-2025-03-02-33-3610-AM.png",
	"Sonatype":    "https://www.finos.org/hs-fs/hubfs/37.png",
	"Scott Logic": "https://www.finos.org/hs-fs/hubfs/69-1.png",
}

// Participant represents a single participant
type Participant struct {
	Name           string `yaml:"name"`
	EnrollmentDate string `yaml:"enrollment_date"`
	GithubID       string `yaml:"github_id"`
}

// ParticipantsData represents the structure of the participants.yaml file
type ParticipantsData struct {
	Participants map[string][]Participant `yaml:"participants"`
}

var companyParticipants ParticipantsData

// LoadParticipants loads the participants data from the YAML file
func LoadParticipants() (ParticipantsData, error) {
	if len(companyParticipants.Participants) > 0 {
		return companyParticipants, nil
	}
	data, err := ioutil.ReadFile("../participants.yaml")
	if err != nil {
		return companyParticipants, fmt.Errorf("failed to read participants file: %v", err)
	}

	err = yaml.Unmarshal(data, &companyParticipants)
	if err != nil {
		return companyParticipants, fmt.Errorf("failed to unmarshal participants data: %v", err)
	}

	return companyParticipants, nil
}

// GetCompanyByParticipantName returns the company name for a given participant name
func GetCompanyByParticipantName(participantName string) (string, error) {
	data, err := LoadParticipants()
	if err != nil {
		return "", err
	}

	for company, participants := range data.Participants {
		for _, participant := range participants {
			if strings.EqualFold(strings.TrimSpace(participant.Name), participantName) {
				return company, nil
			}
		}
	}

	return "", fmt.Errorf("participant '%s' not found", participantName)
}

// GetCompanyByGithubID returns the company name for a given GitHub ID
func GetCompanyByGithubID(githubID string) (string, error) {
	data, err := LoadParticipants()
	if err != nil {
		return "", err
	}
	for company, participants := range data.Participants {
		for _, participant := range participants {
			if participant.GithubID == githubID {
				return company, nil
			}
		}
	}

	return "", fmt.Errorf("GitHub ID '%s' not found", githubID)
}
