package gnum

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

const (
	red testColor = iota
	blue
	yellow
	green testColor = -1
)

type (
	testColor = Enum[struct {
		RedSuffix,
		BlueSuffix,
		YellowSuffix color
		GreenSuffix color `gnum:"value=-1"`
	}]
	color int
)

type enumMetadataTestSuite struct {
	suite.Suite
}

func (s *enumMetadataTestSuite) SetupTest() {
	SetOptions(
		ParseCallback(func(value string) string {
			return value + "Suffix"
		}),
		StringCallback(func(enumName string) string {
			return "Prefix" + enumName
		}))
}

func (s *enumMetadataTestSuite) TearDownSuite() {
	SetOptions(
		ParseCallback(nil),
		StringCallback(nil))
}

func TestEnumMetadataTestSuite(t *testing.T) {
	suite.Run(t, new(enumMetadataTestSuite))
}

func (s *enumMetadataTestSuite) TestEnumMetadata_OnDuplicateEnumValues_ThenPanic() {
	// Arrange
	type (
		color_ int
		enum   = Enum[struct {
			Red    color_
			Blue   color_
			Green  color_ `gnum:"value=0"`
			Yellow color_
		}]
	)

	red_ := enum(0)

	// Act
	// Assert
	assert.Panics(s.T(), func() {
		red_.Enums()
	})
}

func (s *enumMetadataTestSuite) TestEnumMetadata_OnDuplicateEnumNames_ThenPanic() {
	// Arrange
	type (
		color_ int
		enum   = Enum[struct {
			Red    color_
			Blue   color_
			Green  color_ `gnum:"name=Red"`
			Yellow color_
		}]
	)

	red_ := enum(0)

	// Act
	// Assert
	assert.Panics(s.T(), func() {
		red_.Enums()
	})
}

func (s *enumMetadataTestSuite) TestEnumMetadata_OnInvalidEnumTagDefinition_ThenPanic() {
	// Arrange
	type (
		color_ int
		enum   = Enum[struct {
			Red    color_ `gnum:"0"`
			Blue   color_ `gnum:"0"`
			Green  color_
			Yellow color_
		}]
	)

	red_ := enum(0)

	// Act
	// Assert
	assert.Panics(s.T(), func() {
		red_.Enums()
	})
}

func (s *enumMetadataTestSuite) TestReceiverString_OnCustomEnumMetadata_ThenReturnString() {
	// Arrange
	// Act
	actualString := red.String()

	// Assert
	assert.Equal(s.T(), "PrefixRedSuffix", actualString)
}

func (s *enumMetadataTestSuite) TestReceiverStrings_OnCustomEnumMetadata_ThenReturnStrings() {
	// Arrange
	// Act
	actualStrings := red.Strings()

	// Assert
	assert.Equal(s.T(), []string{"PrefixGreenSuffix", "PrefixRedSuffix", "PrefixBlueSuffix", "PrefixYellowSuffix"}, actualStrings)
}

func (s *enumMetadataTestSuite) TestReceiverName_OnCustomEnumMetadata_ThenReturnName() {
	// Arrange
	// Act
	actualName := red.Name()

	// Assert
	assert.Equal(s.T(), "RedSuffix", actualName)
}

func (s *enumMetadataTestSuite) TestReceiverNames_OnCustomEnumMetadataAndMultipleEnums_ThenReturnDifferentNames() {
	// Arrange
	// Act
	actualNames := red.Names()

	// Assert
	assert.Equal(s.T(), []string{"GreenSuffix", "RedSuffix", "BlueSuffix", "YellowSuffix"}, actualNames)
}

func (s *enumMetadataTestSuite) TestReceiverParse_OnCustomEnumMetadataAndExistingEnumName_ThenReturnEnum() {
	// Arrange
	// Act
	actualEnum, err := red.Parse("Blue")
	require.NoError(s.T(), err)

	// Assert
	assert.Equal(s.T(), blue, actualEnum)
}

func (s *enumMetadataTestSuite) TestReceiverParse_OnCustomEnumMetadataAndNonExistingEnumName_ThenReturnError() {
	// Arrange
	// Act
	_, err := red.Parse("blue")

	// Assert
	assert.Error(s.T(), err)
}

func (s *enumMetadataTestSuite) TestReceiverParse_OnCaseInsensitiveAndExists_ThenReturnEnum() {
	// Arrange
	SetOptions(CaseInsensitive(true))
	defer SetOptions(CaseInsensitive(false))

	// Act
	actualEnum, err := red.Parse("BLUE")
	require.NoError(s.T(), err)

	// Assert
	assert.Equal(s.T(), blue, actualEnum)
}

func (s *enumMetadataTestSuite) TestReceiverParse_OnCaseInsensitiveAndMissing_ThenReturnError() {
	// Arrange
	SetOptions(CaseInsensitive(true))
	defer SetOptions(CaseInsensitive(false))

	// Act
	_, err := red.Parse("NOT_COLOR")

	// Assert
	assert.Error(s.T(), err)
}

func (s *enumMetadataTestSuite) TestReceiverEnums_OnCustomEnumMetadataAndMultipleEnums_ThenReturnAll() {
	// Arrange
	// Act
	actualEnums := red.Enums()

	// Assert
	assert.Equal(s.T(), []testColor{green, red, blue, yellow}, actualEnums)
}

func (s *enumMetadataTestSuite) TestReceiverMarshalText_OnCustomEnumMetadata_ThenReturnName() {
	// Arrange
	// Act
	actualTextBytes, err := red.MarshalText()
	require.NoError(s.T(), err)

	// Assert
	assert.Equal(s.T(), []byte("RedSuffix"), actualTextBytes)
}

func (s *enumMetadataTestSuite) TestReceiverMarshalText_OnCustomEnumMetadataAndJsonMarshal_ThenReturnJsonEncoded() {
	// Arrange
	// Act
	actualJsonBytes, err := json.Marshal(red)
	require.NoError(s.T(), err)

	// Assert
	assert.Equal(s.T(), []byte("\"RedSuffix\""), actualJsonBytes)
}

func (s *enumMetadataTestSuite) TestStrings_OnCustomEnumMetadataAndMultipleEnums_ThenReturnDifferentStrings() {
	// Arrange
	// Act
	actualNames := Strings[testColor]()

	// Assert
	assert.Equal(s.T(), []string{"PrefixGreenSuffix", "PrefixRedSuffix", "PrefixBlueSuffix", "PrefixYellowSuffix"}, actualNames)
}

func (s *enumMetadataTestSuite) TestNames_OnCustomEnumMetadataAndMultipleEnums_ThenReturnDifferentNames() {
	// Arrange
	// Act
	actualNames := Names[testColor]()

	// Assert
	assert.Equal(s.T(), []string{"GreenSuffix", "RedSuffix", "BlueSuffix", "YellowSuffix"}, actualNames)
}

func (s *enumMetadataTestSuite) TestParse_OnCustomEnumMetadataAndExistingEnumName_ThenReturnEnum() {
	// Arrange
	// Act
	actualEnum, err := Parse[testColor]("Red")
	require.NoError(s.T(), err)

	// Assert
	assert.Equal(s.T(), red, actualEnum)
}

func (s *enumMetadataTestSuite) TestParse_OnCustomEnumMetadataAndNonExistingEnumName_ThenReturnError() {
	// Arrange
	// Act
	_, err := Parse[testColor]("RedSuffix")

	// Assert
	assert.Error(s.T(), err)
}

func (s *enumMetadataTestSuite) TestParse_OnCaseInsensitiveAndExists_ThenReturnEnum() {
	// Arrange
	SetOptions(CaseInsensitive(true))
	defer SetOptions(CaseInsensitive(false))

	// Act
	actualEnum, err := Parse[testColor]("BLUE")
	require.NoError(s.T(), err)

	// Assert
	assert.Equal(s.T(), blue, actualEnum)
}

func (s *enumMetadataTestSuite) TestParse_OnCaseInsensitiveAndMissing_ThenReturnError() {
	// Arrange
	SetOptions(CaseInsensitive(true))
	defer SetOptions(CaseInsensitive(false))

	// Act
	_, err := Parse[testColor]("NOT_COLOR")

	// Assert
	assert.Error(s.T(), err)
}

func (s *enumMetadataTestSuite) TestEnums_OnCustomEnumMetadataAndMultipleEnums_ThenReturnAll() {
	// Arrange
	// Act
	actualEnums := Enums[testColor]()

	// Assert
	assert.Equal(s.T(), []testColor{green, red, blue, yellow}, actualEnums)
}
