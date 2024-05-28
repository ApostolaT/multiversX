package argsParser

import (
	"errors"
	"strconv"
)

type IssueNFTArgs struct {
	JsonFile         string
	Password         string
	Value            uint64
	GasLimit         uint64
	CollectionName   string
	CollectionTicker string
}

const EGLD = 1000000000000000000

func ParseIssueNFTArgs(args []string) (IssueNFTArgs, error) {
	var issueNFTArgs IssueNFTArgs

	if len(args) != 6 {
		return issueNFTArgs, errors.New("invalid number of arguments")
	}

	issueNFTArgs.JsonFile = args[0]
	if len(issueNFTArgs.JsonFile) < 6 {
		return issueNFTArgs, errors.New("invalid length of jsonFile argument")
	}

	issueNFTArgs.Password = args[1]
	if len(issueNFTArgs.Password) < 8 {
		return issueNFTArgs, errors.New("password too short")
	}

	value := args[2]
	v, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return issueNFTArgs, errors.New("invalid value provided for NFT token")
	}
	issueNFTArgs.Value = uint64(v * EGLD)

	gl, err := strconv.Atoi(args[3])
	if err != nil {
		return issueNFTArgs, errors.New("invalid gasLimit provided for NFT token")
	}
	issueNFTArgs.GasLimit = uint64(gl)

	issueNFTArgs.CollectionName = args[4]
	if len(issueNFTArgs.CollectionName) < 3 {
		return issueNFTArgs, errors.New("collection name too short")
	}

	issueNFTArgs.CollectionTicker = args[5]
	if len(issueNFTArgs.CollectionTicker) != 3 {
		return issueNFTArgs, errors.New("collection ticker too short, must be at least 3 characters")
	}

	return issueNFTArgs, nil
}
