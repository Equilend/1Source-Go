package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/EquiLend/1Source-Go/api"
	"github.com/EquiLend/1Source-Go/models"
	"github.com/EquiLend/1Source-Go/utils"
	"github.com/Nerzal/gocloak/v13"
)

var (
	LogFile   = "1source-go.log"
	fileName  string
	token     *gocloak.JWT
	appConfig *models.AppConfig
)

func main() {
	// Open the log file
	var logFile, err = os.OpenFile(LogFile, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		log.Panic("Error opening the Log file: ", err)
	}

	defer func(logFile *os.File) {
		err := logFile.Close()
		if err != nil {
			fmt.Printf("Error closing the log file	 '%s'", LogFile)
		}
	}(logFile)

	// Set the log to output the Log file
	log.SetOutput(logFile)

	// Begin parsing the command line arguments
	if len(os.Args) == 1 {
		utils.DisplayHelp()

		// Graceful exit after displaying help
		os.Exit(0)
	}

	argsWithoutProg := os.Args[1:]

	// Command line of length 1 usually means help or version info requested
	if len(argsWithoutProg) == 1 {
		switch argsWithoutProg[0] {
		case "--help", "help", "-h":
			utils.DisplayHelp()
		case "--version", "-v":
			utils.DisplayVersion()
		default:
			utils.DisplayHelp()
		}

		// Graceful exit after displaying help
		os.Exit(0)
	}

	// Command line of length 2 means either help or version
	if len(argsWithoutProg) == 2 {
		// Command line of length 2 means -t TOML file
		if argsWithoutProg[0] == "-t" {
			fileName = argsWithoutProg[1]

			// Read and parse configuration TOML file
			appConfig, err = utils.ReadTOML(fileName)

			if err != nil {
				log.Println("Error reading and parsing configuration TOML file: ", err)
				os.Exit(10)
			}
		} else {
			log.Println("Unknown command line flag combination")
			os.Exit(20)
		}

		// Graceful exit after reading and parsing configuration TOML file
		os.Exit(0)
	}

	// Command line of length 3 is not supported
	if len(argsWithoutProg) == 3 {
		log.Println("Unknown command line flag combination")
		os.Exit(30)
	}

	// Command line of length 4 contains the actual command to execute
	if len(argsWithoutProg) == 4 {
		fileName = argsWithoutProg[1]

		// Read and parse configuration TOML file
		appConfig, err = utils.ReadTOML(fileName)

		if err != nil {
			log.Println("Error reading and parsing configuration TOML file: ", err)
			os.Exit(15)
		}

		// Get Auth Token using credentials from config file
		token, err = api.GetAuthToken(appConfig)
		var bearer string

		if err != nil {
			log.Panic("Error retrieving Auth Token: ", err)
		} else {
			bearer = `Bearer ` + token.AccessToken
		}

		// Get the 3rd and 4th command line parameters
		// The 3rd parameter will be a switch, the 4th parameter will be the entity
		param := argsWithoutProg[2]
		entity := argsWithoutProg[3]

		switch param {
		// Get all of a particular type from the API
		case "-g":
			switch entity {
			case "events":
				header := "1Source Events"
				events, err := api.GetEntity(appConfig.Endpoints.Events, bearer, header)
				utils.PrintResults(err, events, "Error retrieving 1Source Events: ", header)

			case "parties":
				header := "1Source Parties"
				parties, err := api.GetEntity(appConfig.Endpoints.Parties, bearer, header)
				utils.PrintResults(err, parties, "Error retrieving 1Source Parties: ", header)

			case "agreements":
				header := "1Source Trade Agreements"
				tas, err := api.GetEntity(appConfig.Endpoints.Agreements, bearer, header)
				utils.PrintResults(err, tas, "Error retrieving 1Source Trade Agreements: ", header)

			case "loans":
				header := "1Source Loans"
				loans, err := api.GetEntity(appConfig.Endpoints.Loans, bearer, header)
				utils.PrintResults(err, loans, "Error retrieving 1Source Loans: ", header)

			case "rerates":
				header := "1Source Rerates"
				rerates, err := api.GetEntity(appConfig.Endpoints.Rerates, bearer, header)
				utils.PrintResults(err, rerates, "Error retrieving 1Source Rerates: ", header)

			case "returns":
				header := "1Source Returns"
				returns, err := api.GetEntity(appConfig.Endpoints.Returns, bearer, header)
				utils.PrintResults(err, returns, "Error retrieving 1Source Returns: ", header)

			case "recalls":
				header := "1Source Recalls"
				recalls, err := api.GetEntity(appConfig.Endpoints.Recalls, bearer, header)
				utils.PrintResults(err, recalls, "Error retrieving 1Source Recalls: ", header)

			case "buyins":
				header := "1Source Buyins"
				buyins, err := api.GetEntity(appConfig.Endpoints.Buyins, bearer, header)
				utils.PrintResults(err, buyins, "Error retrieving 1Source Buyins: ", header)

			default:
				log.Println("Unknown command-line entity entered: ", entity)
				fmt.Println("Unknown command-line entity entered: ", entity)
			}

		// Get trade agreement by agreement_id
		case "-a":
			header := "1Source Trade Agreement"
			prompt := fmt.Sprintf("Error retrieving Trade Agreement with agreement_id = [%s]: ", entity)
			agreement, err := api.GetEntityById(appConfig.Endpoints.Agreements, entity, bearer, header)
			utils.PrintResults(err, agreement, prompt, header)

		// Get event agreement by event_id
		case "-e":
			header := "1Source Event"
			prompt := fmt.Sprintf("Error retrieving Event with event_id = [%s]: ", entity)
			event, err := api.GetEntityById(appConfig.Endpoints.Events, entity, bearer, header)
			utils.PrintResults(err, event, prompt, header)

		// Get loan by loan_id
		case "-l":
			header := "1Source Loan"
			prompt := fmt.Sprintf("Error retrieving Loan with loan_id = [%s]: ", entity)
			loan, err := api.GetEntityById(appConfig.Endpoints.Loans, entity, bearer, header)
			utils.PrintResults(err, loan, prompt, header)

		// Get loan history by loan_id
		case "-lh":
			header := "1Source Loan History"
			prompt := fmt.Sprintf("Error retrieving Loan History with loan_id = [%s]: ", entity)
			endPoint := appConfig.Endpoints.Loans + "/" + entity + "/history"
			history, err := api.GetEntity(endPoint, bearer, header)
			utils.PrintResults(err, history, prompt, header)

		// Get party by party_id
		case "-p":
			header := "1Source Party"
			prompt := fmt.Sprintf("Error retrieving 1Source with party_id = [%s]: ", entity)
			party, err := api.GetEntityById(appConfig.Endpoints.Parties, entity, bearer, "Party")
			utils.PrintResults(err, party, prompt, header)

		// Propose loan
		case "-lp":
			// Read on JSON file specified on the command line as bytes
			body, err := os.ReadFile(entity)
			if err != nil {
				fmt.Printf("Error JSON reading file [%s]: %s\n", entity, err)
				log.Printf("Error JSON reading file [%s]: %s\n", entity, err)
			}

			// Do HTTP PostProposeLoan to initiate the loan
			resp, err := api.PostProposeLoan(appConfig.Endpoints.Loans, bearer, body)

			if err == nil {
				fmt.Println("Success: ", resp)
			} else {
				fmt.Println("Error proposing loan: ", err)
			}

		// Cancel a proposed loan
		case "-lc":
			// Get the Loan by loan_id - check that it is in the proposed state
			loan, err := api.GetEntityById(appConfig.Endpoints.Loans, entity, bearer, "1Source Loan")

			if err != nil {
				log.Printf("Error GET %s by id [%s]: %s", "Loan", entity, err)
			} else {
				// Check the state of the loan
				if strings.Contains(loan, "PROPOSED") {
					// Do HTTP PostCanceLoan to cancel the loan
					endPoint := appConfig.Endpoints.Loans + "/" + entity + "/cancel"
					resp, err := api.PostCancelLoan(endPoint, bearer)

					if err == nil {
						fmt.Println("Successful: ", resp)
					} else {
						fmt.Println("Error canceling loan: ", err)
					}

				} else {
					fmt.Printf("Loan with id [%s] is not in PROPOSED state and cannot be canceled\n", entity)
				}
			}

		// Decline a proposed loan
		case "-ld":
			// Decline the Loan by loan_id - check that it is in the proposed state
			loan, err := api.GetEntityById(appConfig.Endpoints.Loans, entity, bearer, "1Source Loan")

			if err != nil {
				log.Printf("Error GET %s by id [%s]: %s", "Loan", entity, err)
			} else {
				// Check the state of the loan
				if strings.Contains(loan, "PROPOSED") {
					// Do HTTP PostCancelLoan to cancel the loan
					endPoint := appConfig.Endpoints.Loans + "/" + entity + "/decline"
					resp, err := api.PostDeclineLoan(endPoint, bearer)

					if err == nil {
						fmt.Println("Successful: ", resp)
					} else {
						fmt.Println("Error canceling loan: ", err)
					}

				} else {
					fmt.Printf("Loan with id [%s] is not in PROPOSED state and cannot be canceled\n", entity)
				}
			}

		default:
			log.Println("Unknown command-line switch entered: ", argsWithoutProg)
		}
	}
}
