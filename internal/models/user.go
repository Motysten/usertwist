package models

import (
	"athelas/usertwist/internal/cesar"
	"fmt"
	"regexp"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type Users []User

const KEY = 4

var AuthorizedUsers = Users{
	{Username: "user1", Password: cesar.Rotate("randomPass1A", cesar.Right, KEY), Token: "8Mb19OuQAJ"},
	{Username: "user2", Password: cesar.Rotate("RandomPass2B", cesar.Right, KEY), Token: "1wVtw244Sg"},
	{Username: "admin", Password: cesar.Rotate("AdminSecret1C", cesar.Right, KEY), Token: "a7P47jXfFM"},
	{Username: "login", Password: cesar.Rotate("ADMINcodeD", cesar.Right, KEY), Token: "KlMnOpQrSt"},
	{Username: "Napoleon", Password: cesar.Rotate("Imperial42E", cesar.Right, KEY), Token: "4dRf9uIOL3"},
	{Username: "Cleopatra", Password: cesar.Rotate("Pharaoh123F", cesar.Right, KEY), Token: "hG7kLmN5Vx"},
	{Username: "Einstein", Password: cesar.Rotate("Relativity5G", cesar.Right, KEY), Token: "2F6YzpH1rD"},
	{Username: "Shakespeare", Password: cesar.Rotate("Bard4everH", cesar.Right, KEY), Token: "sQ8u7vPw3X"},
	{Username: "Gandhi", Password: cesar.Rotate("PeaceAnd1I", cesar.Right, KEY), Token: "9ZxNtJ6LkY"},
	{Username: "Aristotle", Password: cesar.Rotate("Philosophy9J", cesar.Right, KEY), Token: "0Bd3rFgH8q"},
	{Username: "Caesar", Password: cesar.Rotate("Julius8kL", cesar.Right, KEY), Token: "xU5mLvT2W1"},
	{Username: "Newton", Password: cesar.Rotate("Gravity7M", cesar.Right, KEY), Token: "q9Ef3KpJz4"},
	{Username: "Curie", Password: cesar.Rotate("Radium6N", cesar.Right, KEY), Token: "7YhF1nM5pX"},
	{Username: "Tesla", Password: cesar.Rotate("Electric5O", cesar.Right, KEY), Token: "z6R4yLtQ3K"},
	{Username: "Lincoln", Password: cesar.Rotate("Abe1861P", cesar.Right, KEY), Token: "fW8k2PvLx9"},
	{Username: "DaVinci", Password: cesar.Rotate("Renaissance4Q", cesar.Right, KEY), Token: "5gQ0NsVy7L"},
	{Username: "Columbus", Password: cesar.Rotate("Sail1492R", cesar.Right, KEY), Token: "3tB1kGmH9X"},
	{Username: "Plato", Password: cesar.Rotate("Ideal3S", cesar.Right, KEY), Token: "w8V6pLdM2Y"},
	{Username: "Beethoven", Password: cesar.Rotate("Symphony2T", cesar.Right, KEY), Token: "0xP3yRtF9Q"},
	{Username: "Mozart", Password: cesar.Rotate("Concerto1U", cesar.Right, KEY), Token: "z4F6mNjH1W"},
	{Username: "Galileo", Password: cesar.Rotate("Telescope0V", cesar.Right, KEY), Token: "5K2tQpL7Vx"},
	{Username: "Darwin", Password: cesar.Rotate("EvolutionZ", cesar.Right, KEY), Token: "r8W1yMnF4X"},
	{Username: "WrightBrothers", Password: cesar.Rotate("FlightY", cesar.Right, KEY), Token: "6zP5kJ3lTy"},
	{Username: "Edison", Password: cesar.Rotate("LightX", cesar.Right, KEY), Token: "f9B7vQ2pX4"},
	{Username: "Armstrong", Password: cesar.Rotate("MoonW", cesar.Right, KEY), Token: "0tL3mFy8Rk"},
	{Username: "Washington", Password: cesar.Rotate("GeneralV", cesar.Right, KEY), Token: "7P9r2WvTxY"},
	{Username: "Roosevelt", Password: cesar.Rotate("TeddyU", cesar.Right, KEY), Token: "1zF6yQ3mLp"},
	{Username: "Freud", Password: cesar.Rotate("PsychoT", cesar.Right, KEY), Token: "x4K8vNtJ9Y"},
	{Username: "Marx", Password: cesar.Rotate("ManifestS", cesar.Right, KEY), Token: "3W7kR5fMpQ"},
	{Username: "Socrates", Password: cesar.Rotate("QuestionR", cesar.Right, KEY), Token: "r6B9pFy2Lt"},
	{Username: "Hemingway", Password: cesar.Rotate("NovelQ", cesar.Right, KEY), Token: "0P3yW7nHxK"},
	{Username: "Austen", Password: cesar.Rotate("PrideP", cesar.Right, KEY), Token: "f8L5kV9Q3x"},
	{Username: "Kafka", Password: cesar.Rotate("MetamorphO", cesar.Right, KEY), Token: "4tB2rJ6mPy"},
	{Username: "Poe", Password: cesar.Rotate("RavenN", cesar.Right, KEY), Token: "x1F9kP3vL6"},
	{Username: "Hawking", Password: cesar.Rotate("BlackHoleM", cesar.Right, KEY), Token: "0W7yR4nF2p"},
	{Username: "Copernicus", Password: cesar.Rotate("HeliocentricL", cesar.Right, KEY), Token: "k9L6pQ3fJx"},
	{Username: "Lovelace", Password: cesar.Rotate("AlgorithmK", cesar.Right, KEY), Token: "3tB8vM5rP0"},
	{Username: "Franklin", Password: cesar.Rotate("ElectricJ", cesar.Right, KEY), Token: "f1W4pJ7kL2"},
	{Username: "Tesla", Password: cesar.Rotate("AlternatingI", cesar.Right, KEY), Token: "6zP3yV9rF5"},
	{Username: "Faraday", Password: cesar.Rotate("InductionH", cesar.Right, KEY), Token: "0tL1mN6yR8"},
	{Username: "Pasteur", Password: cesar.Rotate("GermTheoryG", cesar.Right, KEY), Token: "r4B9kP3vJ2"},
	{Username: "Fleming", Password: cesar.Rotate("PenicillinF", cesar.Right, KEY), Token: "f8W5tL1pQ7"},
	{Username: "Jobs", Password: cesar.Rotate("AppleE", cesar.Right, KEY), Token: "0zF2yN7mR6"},
	{Username: "Gates", Password: cesar.Rotate("MicrosoftD", cesar.Right, KEY), Token: "k3L9pJ5fV1"},
	{Username: "Zuckerberg", Password: cesar.Rotate("FacebookC", cesar.Right, KEY), Token: "x2B7vP6rF4"},
	{Username: "Musk", Password: cesar.Rotate("TeslaB", cesar.Right, KEY), Token: "5tW3yL9kN0"},
	{Username: "Bezos", Password: cesar.Rotate("AmazonA", cesar.Right, KEY), Token: "r1F8pJ6vQ3"},
	{Username: "Branson", Password: cesar.Rotate("VirginZ1", cesar.Right, KEY), Token: "0tP7mL3yV5"},
	{Username: "Buffett", Password: cesar.Rotate("BerkshireY2", cesar.Right, KEY), Token: "f9B2kN6rJ4"},
	{Username: "Picasso", Password: cesar.Rotate("GuernicaX3", cesar.Right, KEY), Token: "x5W3pQ1tL8"},
	{Username: "VanGogh", Password: cesar.Rotate("StarryNightW4", cesar.Right, KEY), Token: "4zP9yJ6vR2"},
	{Username: "Dali", Password: cesar.Rotate("PersistenceV5", cesar.Right, KEY), Token: "k8L1mN7fV3"},
	{Username: "Warhol", Password: cesar.Rotate("MarilynU6", cesar.Right, KEY), Token: "3F9rJ4vP5Q"},
	{Username: "Michelangelo", Password: cesar.Rotate("SistineT7", cesar.Right, KEY), Token: "0G7yP3nL8R"},
	{Username: "leonardodaVinci", Password: cesar.Rotate("secretCode123!", cesar.Right, KEY), Token: "34f5g6h7j8"},
	{Username: "queenElizabethI", Password: cesar.Rotate("password456*", cesar.Right, KEY), Token: "k9l8m7n6o5"},
	{Username: "sherlockHolmes", Password: cesar.Rotate("mysterySolve", cesar.Right, KEY), Token: "p1q2r3s4t5"},
	{Username: "marieCurie", Password: cesar.Rotate("scienceRocks!", cesar.Right, KEY), Token: "v6w7x8y9z0"},
	{Username: "napoleonBonaparte", Password: cesar.Rotate("empireBuilder", cesar.Right, KEY), Token: "b1c2d3e4f5"},
	{Username: "janeAusten", Password: cesar.Rotate("loveAndPride", cesar.Right, KEY), Token: "g6h7i8j9k0"},
	{Username: "albertEinstein", Password: cesar.Rotate("physicsGenius", cesar.Right, KEY), Token: "m1n2o3p4q5"},
	{Username: "cleopatra", Password: cesar.Rotate("queenOfNile", cesar.Right, KEY), Token: "s6t7u8v9w0"},
	{Username: "michaelangelo", Password: cesar.Rotate("artMasterpiece", cesar.Right, KEY), Token: "z1a2b3c4d5"},
	{Username: "plato", Password: cesar.Rotate("philosophyQuest", cesar.Right, KEY), Token: "f6g7h8i9j0"},
	{Username: "juliusCaesar", Password: cesar.Rotate("romanEmpire", cesar.Right, KEY), Token: "k1l2m3n4o5"},
	{Username: "alexanderTheGreat", Password: cesar.Rotate("conquerorWorld", cesar.Right, KEY), Token: "p6q7r8s9t0"},
	{Username: "confucius", Password: cesar.Rotate("wisdomTeacher", cesar.Right, KEY), Token: "v1w2x3y4z5"},
	{Username: "williamShakespeare", Password: cesar.Rotate("playwrightBard", cesar.Right, KEY), Token: "b6c7d8e9f0"},
	{Username: "galileoGalilei", Password: cesar.Rotate("astronomyStars", cesar.Right, KEY), Token: "g1h2i3j4k5"},
}

func (users *Users) GetToken(username string, password string) (string, error) {
	cipheredPassword := cesar.Rotate(password, cesar.Right, KEY)
	for _, u := range *users {
		if u.Username == username && u.Password == cipheredPassword {
			return u.Token, nil

		}
	}

	return "", fmt.Errorf("User not found")
}

func (users *Users) IsAuthorized(token string) bool {
	for _, u := range *users {
		if u.Token == token {
			return true
		}
	}

	return false
}

func (users *Users) Search(term string) (Users, error) {
	matches := make(Users, 0)

	regex, err := regexp.Compile(term)
	if err != nil {
		return Users{}, err
	}

	for _, u := range *users {
		if regex.Match([]byte(u.Username)) {
			matches = append(matches, u)
		}
	}

	return matches, nil
}
