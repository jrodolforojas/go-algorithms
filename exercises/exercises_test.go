package exercises

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type exercisesSuite struct {
	exe Exercises
	suite.Suite
}

func (suite *exercisesSuite) SetupTest() {
	suite.exe = Exercises{}
}

func TestExercises(t *testing.T) {
	es := new(exercisesSuite)
	suite.Run(t, es)
}

func (suite *exercisesSuite) TestFizzBuzz() {
	results := suite.exe.Fizzbuzz([]int{3, 5, 4, 15, 32})
	assert.NotNil(suite.T(), results)
	assert.Equal(suite.T(), []string{"fizz", "buzz", "4", "fizzbuzz", "32"}, results)
}

func (suite *exercisesSuite) TestHackerLanguage() {
	first := suite.exe.HackerLanguage("hola mundo")
	assert.NotEmpty(suite.T(), first)
	assert.Equal(suite.T(), "h0l4 mund0", first)

	second := suite.exe.HackerLanguage("cuodlibetal")
	assert.NotEmpty(suite.T(), second)
	assert.Equal(suite.T(), "cu0dl1b3t4l", second)
}

func (suite *exercisesSuite) TestPasswordGenerator() {
	pass := suite.exe.passwordGenerator(12, true, true, true)
	assert.NotEmpty(suite.T(), pass)

	pass = suite.exe.passwordGenerator(5, true, true, true)
	assert.Empty(suite.T(), pass)
}

func (suite *exercisesSuite) TestPrimoFibonacciPar() {
	result := suite.exe.primoFibonacciPar(2)
	fmt.Println(result)
	array := strings.Split(result, " ")
	assert.Equal(suite.T(), array[1], "true")
	assert.Equal(suite.T(), array[2], "true")
	assert.Equal(suite.T(), array[3], "true")

	seven := suite.exe.primoFibonacciPar(7)
	fmt.Println(seven)
	sevenArray := strings.Split(seven, " ")
	assert.Equal(suite.T(), sevenArray[1], "false")
	assert.Equal(suite.T(), sevenArray[2], "false")
	assert.Equal(suite.T(), sevenArray[3], "false")
}

func (suite *exercisesSuite) TestRockPaperScissors() {
	winner := suite.exe.rockPaperScissors([]Play{
		{player1: "🪨", player2: "🪨"},  // tie
		{player1: "🪨", player2: "📄"},  // player 2
		{player1: "✂️", player2: "🪨"}, // player 2
	})
	assert.Equal(suite.T(), "player 2", winner)

	winner = suite.exe.rockPaperScissors([]Play{
		{player1: "🪨", player2: "🪨"},  // tie
		{player1: "🪨", player2: "🪨"},  // tie
		{player1: "🪨", player2: "✂️"}, // player 1
	})
	assert.Equal(suite.T(), "player 1", winner)

	winner = suite.exe.rockPaperScissors([]Play{
		{player1: "🪨", player2: "🪨"},  // tie
		{player1: "📄", player2: "🪨"},  // player 1
		{player1: "✂️", player2: "🪨"}, // player 2
	})
	assert.Equal(suite.T(), "tie", winner)
}

func (suite *exercisesSuite) TestHeterogramIsogramPangram() {
	result := suite.exe.isHeterogram("centrifugado")
	assert.Equal(suite.T(), true, result)

	result = suite.exe.isHeterogram("rodolfo")
	assert.Equal(suite.T(), false, result)

	result = suite.exe.isIsogram("anna")
	assert.Equal(suite.T(), true, result)

	result = suite.exe.isIsogram("foo")
	assert.Equal(suite.T(), false, result)

	result = suite.exe.isPangram("Waltz bad nymph for quick jigs vex")
	assert.Equal(suite.T(), true, result)

	result = suite.exe.isPangram("Pac my box with five dozen liquor jugs")
	assert.Equal(suite.T(), false, result)
}

func (suite *exercisesSuite) TestGetCharacters() {
	characters, err := suite.exe.getRickAndMortyCharacters()
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), characters)
	assert.Len(suite.T(), characters.Results, 20)
}
