package fqdn_test

import (
	"strings"
	"testing"
	"time"

	fqdn "github.com/go-random/fqdn"
	"github.com/stretchr/testify/assert"
)

func TestFQDN(t *testing.T) {
	// Arrange
	randomizer := fqdn.NewRandomizer()

	// Act
	randomFQDN := randomizer.FQDN()
	parts := strings.Split(randomFQDN, ".")

	// Assert
	expectedParts := 2
	assert.Equal(t, expectedParts, len(parts), "An FQDN should have 2 parts, the domain and the TLD (example.com)")
}

func TestFQDN_WithSeed(t *testing.T) {
	// Arrange
	randomizer := fqdn.NewRandomizerWithSeed(time.Now().UnixNano())

	// Act
	randomFQDN := randomizer.FQDN()
	parts := strings.Split(randomFQDN, ".")

	// Assert
	expectedParts := 2
	assert.Equal(t, expectedParts, len(parts), "An FQDN should have 2 parts, the domain and the TLD (example.com)")
}

func TestFQDN_WhenCalledMultipleTimes_ShouldReturnDifferentFQDNs(t *testing.T) {
	// Arrange
	randomizer := fqdn.NewRandomizer()

	// Act
	randomFQDN1 := randomizer.FQDN()
	randomFQDN2 := randomizer.FQDN()

	// Assert
	assert.NotEqual(t, randomFQDN1, randomFQDN2, "FQDNs should be different")
}

func TestFQDN_WhenCalledMultipleTimes_ShouldReturnDifferentFQDNs_WithSeed(t *testing.T) {
	// Arrange
	randomizer := fqdn.NewRandomizerWithSeed(time.Now().UnixNano())

	// Act
	randomFQDN1 := randomizer.FQDN()
	randomFQDN2 := randomizer.FQDN()

	// Assert
	assert.NotEqual(t, randomFQDN1, randomFQDN2, "FQDNs should be different")
}

func TestFQDNWithSubDomain(t *testing.T) {
	// Arrange
	randomizer := fqdn.NewRandomizer()

	// Act
	randomFQDN := randomizer.FQDNWithSubDomain()
	parts := strings.Split(randomFQDN, ".")

	// Assert
	expectedParts := 3
	assert.Equal(t, expectedParts, len(parts), "An FQDN including a subdomain should have 3 parts, the subdomain, domain and the TLD (api.example.com)")
}

func TestFQDNWithSubDomain_WithSeed(t *testing.T) {
	// Arrange
	randomizer := fqdn.NewRandomizerWithSeed(time.Now().UnixNano())

	// Act
	randomFQDN := randomizer.FQDNWithSubDomain()
	parts := strings.Split(randomFQDN, ".")

	// Assert
	expectedParts := 3
	assert.Equal(t, expectedParts, len(parts), "An FQDN including a subdomain should have 3 parts, the subdomain, domain and the TLD (api.example.com)")
}

func TestFQDNWithSubDomain_WhenCalledMultipleTimes_ShouldReturnDifferentFQDNs(t *testing.T) {
	// Arrange
	randomizer := fqdn.NewRandomizer()

	// Act
	randomFQDN1 := randomizer.FQDNWithSubDomain()
	randomFQDN2 := randomizer.FQDNWithSubDomain()

	// Assert
	assert.NotEqual(t, randomFQDN1, randomFQDN2, "FQDNs should be different")
}

func TestFQDNWithSubDomain_WhenCalledMultipleTimes_ShouldReturnDifferentFQDNs_WithSeed(t *testing.T) {
	// Arrange
	randomizer := fqdn.NewRandomizerWithSeed(time.Now().UnixNano())

	// Act
	randomFQDN1 := randomizer.FQDNWithSubDomain()
	randomFQDN2 := randomizer.FQDNWithSubDomain()

	// Assert
	assert.NotEqual(t, randomFQDN1, randomFQDN2, "FQDNs should be different")
}

func TestFQDNWithSubDomain_ShouldReturnDifferentThanFQDN(t *testing.T) {
	// Arrange
	randomizer := fqdn.NewRandomizer()

	// Act
	randomFQDN1 := randomizer.FQDN()
	randomFQDN2 := randomizer.FQDNWithSubDomain()

	// Assert
	assert.NotEqual(t, randomFQDN1, randomFQDN2, "FQDNs should be different")
}

func TestFQDNWithSubDomain_ShouldReturnDifferentThanFQDN_WithSeed(t *testing.T) {
	// Arrange
	randomizer := fqdn.NewRandomizerWithSeed(time.Now().UnixNano())

	// Act
	randomFQDN1 := randomizer.FQDN()
	randomFQDN2 := randomizer.FQDNWithSubDomain()

	// Assert
	assert.NotEqual(t, randomFQDN1, randomFQDN2, "FQDNs should be different")
}
