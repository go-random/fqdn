package fqdn

// FQDN generates a random Fully Qualified Domain Name (FQDN)
func (tr *Randomizer) FQDN() string {
	randomDomain := domains[tr.Rand.Intn(len(domains))]
	randomTLD := tlds[tr.Rand.Intn(len(tlds))]
	return randomDomain + "." + randomTLD
}

func (tr *Randomizer) FQDNWithSubDomain() string {
	randomDomain := domains[tr.Rand.Intn(len(domains))]
	randomSubDomain := subdomains[tr.Rand.Intn(len(subdomains))]
	randomTLD := tlds[tr.Rand.Intn(len(tlds))]
	fqdn := randomSubDomain + "." + randomDomain + "." + randomTLD
	return fqdn
}
