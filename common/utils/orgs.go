package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

type org struct {
	Type       string
	Address    string
	PrivateKey string
}

func InitOrgsForDevEnv() error {

	orgsBytes, err := os.ReadFile("../../nft-deploy/orgs.json")
	if err != nil {
		return err
	}

	var dltOrgs = struct {
		Addresses   map[string]string `json:"addresses"`
		PrivateKeys map[string]string `json:"private_keys"`
	}{}

	err = json.Unmarshal(orgsBytes, &dltOrgs)
	if err != nil {
		return err
	}

	if len(dltOrgs.Addresses) != 4 {
		return fmt.Errorf("found only %v orgs wallets", len(dltOrgs.Addresses))
	}

	var orgs = map[string]org{}

	var orgsTypes = []string{"deployer", "initiator", "transactioner", "auditor"}

	var orgIndex = 0
	for k := range dltOrgs.Addresses {
		orgs[orgsTypes[orgIndex]+"01"] = org{orgsTypes[orgIndex], k, dltOrgs.PrivateKeys[k]}
		orgIndex++
	}

	payload, err := json.Marshal(orgs)
	if err != nil {
		return err
	}
	err = os.WriteFile("../common/config/orgs.json", payload, 0777)
	if err != nil {
		return err
	}

	return nil
}

func GetOrgKeys(orgName string) (address, privateKey, publicKey string, err error) {

	orgsBytes, err := os.ReadFile("../common/config/orgs.json")
	if err != nil {
		return "", "", "", err
	}

	var orgs = map[string]org{}

	err = json.Unmarshal(orgsBytes, &orgs)
	if err != nil {
		return "", "", "", err
	}

	org, isExits := orgs[orgName]
	if !isExits {
		return "", "", "", fmt.Errorf("org not exists")
	}

	return org.Address, org.PrivateKey, "", nil
}
