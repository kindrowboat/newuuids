package uuidbump

import (
	"fmt"
	"testing"
)

const testString = `Hello D64F6BB5-518B-4BDF-A12F-BB3A698224CA
Hello CB90F03D-7B45-4645-B033-FF315363C8EE
Hello again D64F6BB5-518B-4BDF-A12F-BB3A698224CA
`

const testString2 = "Why Hello again D64F6BB5-518B-4BDF-A12F-BB3A698224CA"

const expectedOutput = `Hello 00000000-0000-0000-0000-00000000000A
Hello 00000000-0000-0000-0000-00000000000B
Hello again 00000000-0000-0000-0000-00000000000A
`

const expectedOutput2 = "Why Hello again 00000000-0000-0000-0000-00000000000A"

func TestBumpUuids(t *testing.T) {
	t.Run("it do", func(t *testing.T) {
		mockedUuidGenerator := mockUuidGenerator{count: 0xa} //start at 10 so we get some letters
		rotator := NewWithMockedGenerator(mockedUuidGenerator.generateSequentialUuid)
		actualOutput := rotator.BumpUuids(testString)
		if(actualOutput != expectedOutput) {
			t.Fatalf("expected %s to equal %s", actualOutput, expectedOutput)
		}

		actualOutput2 := rotator.BumpUuids(testString2)
		if(actualOutput2 != expectedOutput2) {
			t.Fatalf("expected %s to equal %s", actualOutput2, expectedOutput2)
		}
	})
}

type mockUuidGenerator struct {
	count int
}
// 8-4-4-4-12
func (m *mockUuidGenerator) generateSequentialUuid() string {
	if m.count > 0xffffffffffff {
		panic("cant generate a uuid over 0xffffffffffff")
	}
	suffix := m.count
	m.count++
	uuidStart := "00000000-0000-0000-0000-"

	return fmt.Sprintf("%s%012x", uuidStart, suffix);
}