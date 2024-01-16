package domain

type Character struct {
	Name string
}

// TODO: make it changeable for admins
var characters = []*Character{
	{Name: "Achilles"},
	{Name: "Alice"},
	{Name: "Angel"},
	{Name: "Annie Christmas"},
	{Name: "Beowulf"},
	{Name: "Bigfoot"},
	{Name: "Black Panther"},
	{Name: "Black Widow"},
	{Name: "Bloody Mary"},
	{Name: "Bruce Lee"},
	{Name: "Buffy"},
	{Name: "Bullseye"},
	{Name: "Cloak and Dagger"},
	{Name: "Daredevil"},
	{Name: "Deadpool"},
	{Name: "Doctor Strange"},
	{Name: "Dr. Sattler"},
	{Name: "Dr. Jill Trent"},
	{Name: "Dracula"},
	{Name: "Elektra"},
	{Name: "Ghost Rider"},
	{Name: "Golden Bat"},
	{Name: "Houdini"},
	{Name: "InGen (Robert Muldoon)"},
	{Name: "Invisible Man"},
	{Name: "Jekyll & Hyde"},
	{Name: "King Arthur"},
	{Name: "Little Red Riding Hood"},
	{Name: "Luke Cage"},
	{Name: "Medusa"},
	{Name: "Moon Knight"},
	{Name: "Ms. Marvel"},
	{Name: "Nikola Tesla"},
	{Name: "Oda Nobunaga"},
	{Name: "Raptors"},
	{Name: "Robin Hood"},
	{Name: "She-Hulk"},
	{Name: "Sherlock Holmes"},
	{Name: "Sinbad"},
	{Name: "Spider-Man"},
	{Name: "Spike"},
	{Name: "Squirrel Girl"},
	{Name: "Sun Wukong"},
	{Name: "T. Rex"},
	{Name: "The Genie"},
	{Name: "Tomoe Gozen"},
	{Name: "Willow"},
	{Name: "Winter Soldier"},
	{Name: "Yennenga"},
}

func GetAllCharacters() []*Character {
	return characters
}
