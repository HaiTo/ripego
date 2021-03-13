package ripego

import "strings"

func ApnicCheck(search, whoisServer string) (*WhoisInfo, error) {
	whoisData, err := getTcpContent(search, whoisServer)

	if err != nil {
		return nil, err
	}

	whoisForInformationRelate := strings.Split(whoisData, "Information related")[1]

	wi := WhoisInfo{
		Inetnum:      parseRPSLValue(whoisForInformationRelate, "inetnum", "inetnum"),
		Netname:      parseRPSLValue(whoisForInformationRelate, "inetnum", "netname"),
		AdminC:       parseRPSLValue(whoisForInformationRelate, "inetnum", "admin-c"),
		Country:      parseRPSLValue(whoisForInformationRelate, "inetnum", "country"),
		Descr:        parseRPSLValue(whoisForInformationRelate, "inetnum", "descr"),
		LastModified: parseRPSLValue(whoisForInformationRelate, "inetnum", "changed"),
		MntBy:        parseRPSLValue(whoisForInformationRelate, "inetnum", "mnt-by"),
		MntLower:     parseRPSLValue(whoisForInformationRelate, "inetnum", "mnt-lower"),
		MntRoutes:    parseRPSLValue(whoisForInformationRelate, "inetnum", "mnt-routes"),
		Source:       parseRPSLValue(whoisForInformationRelate, "inetnum", "source"),
		TechC:        parseRPSLValue(whoisForInformationRelate, "inetnum", "tech-c"),
		Organization: parseRPSLValue(whoisForInformationRelate, "irt", "irt"),

		Person: WhoisPerson{
			Name:         parseRPSLValue(whoisForInformationRelate, "role", "role"),
			AbuseMailbox: parseRPSLValue(whoisForInformationRelate, "irt", "abuse-mailbox"),
			Address:      parseRPSLValue(whoisForInformationRelate, "role", "address"),
			LastModified: parseRPSLValue(whoisForInformationRelate, "role", "changed"),
			MntBy:        parseRPSLValue(whoisForInformationRelate, "role", "mnt-by"),
			NicHdl:       parseRPSLValue(whoisForInformationRelate, "role", "nic-hdl"),
			Phone:        parseRPSLValue(whoisForInformationRelate, "role", "phone"),
			Source:       parseRPSLValue(whoisForInformationRelate, "role", "source"),
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
