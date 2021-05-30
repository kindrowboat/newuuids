package uuidbump

import (
	"regexp"
	"strings"

	"github.com/google/uuid"
)


type fnUuidGenerator func() string

// UuidRotator can rotate the uuids in one or more strings
type UuidRotator struct {
	generateUuid fnUuidGenerator
	rotatedUuids map[string]string
}

// New makes a new UuidRotator with a default random UUID generator
func New() *UuidRotator{
	return &UuidRotator {
		generateUuid: generateRandomUuid,
		rotatedUuids: map[string]string{},
	}
}

// NewWithMockedGenerator makes a new UuidRotator with a custom UUID generator (for testing)
func NewWithMockedGenerator(generator fnUuidGenerator) *UuidRotator {
	return &UuidRotator{
		generateUuid: generator,
		rotatedUuids: map[string]string{},
	}
}

// BumpUuids searches the string for any UUIDs and replaces them with new ones
// if the receiver has seen any given UUID before, it will replace it with the same UUID it replaced it with previously
func (rotator *UuidRotator) BumpUuids(input string) string {
	re, err := regexp.Compile("[0-9a-fA-F]{8}(-[0-9a-fA-F]{4}){3}-[0-9a-fA-F]{12}")
	if err != nil {
		panic(err);
	}

	output := re.ReplaceAllStringFunc(input, rotator.bumpUuid)

	return output
}

func (rotator *UuidRotator) bumpUuid(oldUuid string) string {
	previouslyUsedNewUuid, seenBefore := rotator.rotatedUuids[oldUuid]
	if seenBefore {
		return previouslyUsedNewUuid
	} else {
		newUuid := strings.ToUpper(rotator.generateUuid())
		rotator.rotatedUuids[oldUuid] = newUuid
		return newUuid

	}
}

func generateRandomUuid() string {
	return uuid.NewString()
}