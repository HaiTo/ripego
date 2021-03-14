package ripego

import (
	"strings"
)

func ApnicCheck(search, whoisServer string) (*WhoisInfo, error) {
	whoisData, err := getTcpContent(search, whoisServer)

	if err != nil {
		return nil, err
	}

	runes := []rune(whoisData)
	var reversed []rune

	for i := len(runes) - 1; i >= 0; i-- {
		reversed = append(reversed, runes[i])
	}

	// reversed Information related
	whoisForInformationRelate := strings.Split(string(reversed), "detaler noitamrofnI")[0]
	reversedSpecifiedRunes := []rune(whoisForInformationRelate)
	var reReversed []rune

	for i := len(reversedSpecifiedRunes) - 1; i >= 0; i-- {
		reReversed = append(reReversed, reversedSpecifiedRunes[i])
	}

	target := string(reReversed)

	wi := WhoisInfo{
		Inetnum:      parseRPSLValue(target, "inetnum", "inetnum"),
		Netname:      parseRPSLValue(target, "inetnum", "netname"),
		AdminC:       parseRPSLValue(target, "inetnum", "admin-c"),
		Country:      parseRPSLValue(target, "inetnum", "country"),
		Descr:        parseRPSLValue(target, "inetnum", "descr"),
		LastModified: parseRPSLValue(target, "inetnum", "changed"),
		MntBy:        parseRPSLValue(target, "inetnum", "mnt-by"),
		MntLower:     parseRPSLValue(target, "inetnum", "mnt-lower"),
		MntRoutes:    parseRPSLValue(target, "inetnum", "mnt-routes"),
		Source:       parseRPSLValue(target, "inetnum", "source"),
		TechC:        parseRPSLValue(target, "inetnum", "tech-c"),
		Organization: parseRPSLValue(target, "irt", "irt"),
		RawData: string(reversed),

		Person: WhoisPerson{
			Name:         parseRPSLValue(target, "role", "role"),
			AbuseMailbox: parseRPSLValue(target, "irt", "abuse-mailbox"),
			Address:      parseRPSLValue(target, "role", "address"),
			LastModified: parseRPSLValue(target, "role", "changed"),
			MntBy:        parseRPSLValue(target, "role", "mnt-by"),
			NicHdl:       parseRPSLValue(target, "role", "nic-hdl"),
			Phone:        parseRPSLValue(target, "role", "phone"),
			Source:       parseRPSLValue(target, "role", "source"),
		},
	}

	return &wi, err
}

func ApnicCheck6(search, whoisServer string) (*WhoisInfo, error) {
	whoisData, err := getTcpContent(search, whoisServer)

	if err != nil {
		return nil, err
	}

	wi := WhoisInfo{
		Inetnum:      parseRPSLValue(whoisData, "inet6num", "inet6num"),
		Netname:      parseRPSLValue(whoisData, "inet6num", "netname"),
		AdminC:       parseRPSLValue(whoisData, "inet6num", "admin-c"),
		Country:      parseRPSLValue(whoisData, "inet6num", "country"),
		Descr:        parseRPSLValue(whoisData, "inet6num", "descr"),
		LastModified: parseRPSLValue(whoisData, "inet6num", "changed"),
		MntBy:        parseRPSLValue(whoisData, "inet6num", "mnt-by"),
		MntLower:     parseRPSLValue(whoisData, "inet6num", "mnt-lower"),
		MntRoutes:    parseRPSLValue(whoisData, "inet6num", "mnt-routes"),
		Source:       parseRPSLValue(whoisData, "inet6num", "source"),
		TechC:        parseRPSLValue(whoisData, "inet6num", "tech-c"),
		Organization: parseRPSLValue(whoisData, "irt", "irt"),

		Person: WhoisPerson{
			Name:         parseRPSLValue(whoisData, "role", "role"),
			AbuseMailbox: parseRPSLValue(whoisData, "irt", "abuse-mailbox"),
			Address:      parseRPSLValue(whoisData, "role", "address"),
			LastModified: parseRPSLValue(whoisData, "role", "changed"),
			MntBy:        parseRPSLValue(whoisData, "role", "mnt-by"),
			NicHdl:       parseRPSLValue(whoisData, "role", "nic-hdl"),
			Phone:        parseRPSLValue(whoisData, "role", "phone"),
			Source:       parseRPSLValue(whoisData, "role", "source"),
		},
	}

	return &wi, err
}
