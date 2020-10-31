package sentiment

import (
	"github.com/amarjeet000/panas-go/internal/text"
)

/*
Utility functions used to build `SelfRefSoundexIndex` and `StatesSoundexIndex` indexes.
*/

// BuildSoundexIndex generates a map of Soundex codes with their corresponding original strings.
func BuildSoundexIndex(words []string) map[string][]string {
	res := map[string][]string{}
	for _, w := range words {
		code := text.Soundex(w)
		if res[code] != nil {
			v := append(res[code], w)
			res[code] = v
		} else {
			res[code] = []string{w}
		}
	}
	return res
}

// BuildSelfRefSoundexIndex generates a map of Soundex codes with their corresponding original self-ref strings.
func BuildSelfRefSoundexIndex() map[string][]string {
	return BuildSoundexIndex(SelfReferences)
}

// BuildStatesSoundexIndex generates a map of Soundex codes with their corresponding original state strings.
func BuildStatesSoundexIndex() map[string][]string {
	return BuildSoundexIndex(SelfReferences)
}

/*
Generated indexes (maps) below. Since the base data is immutable, we can pre-generate and
keep an immutable copy of the required indexes.
We use Soundex algorithm to generate keys of the index maps. The values are a collection of strings,
which helps in the case of a Soundex-code collision.
*/

// SelfRefSoundexIndex is a map of Soundex codes and their corresponding self-ref strings.
var SelfRefSoundexIndex = map[string][]string{
	"I000": []string{"I"},
	"I500": []string{"I'm", "I am"},
	"A500": []string{"am"},
	"F452": []string{"feeling"},
	"M000": []string{"me"},
	"M241": []string{"myself"},
}

// StatesSoundexIndex is a map of Soundex codes and their corresponding state strings.
var StatesSoundexIndex = map[string][]string{
	"A163": []string{"afraid"},
	"A235": []string{"astonished"},
	"A253": []string{"ashamed"},
	"A320": []string{"at ease"},
	"A353": []string{"attentiveness"},
	"A450": []string{"alone"},
	"A463": []string{"alert"},
	"A523": []string{"amazed"},
	"A526": []string{"angry", "angry at self"},
	"B214": []string{"bashful"},
	"B400": []string{"blue"},
	"B430": []string{"bold"},
	"B456": []string{"blameworthy"},
	"C450": []string{"calm"},
	"C513": []string{"confident"},
	"C525": []string{"concentrating"},
	"C614": []string{"cheerful"},
	"D222": []string{"disgusted", "disgusted with self"},
	"D232": []string{"dissatisfied with self"},
	"D365": []string{"determined"},
	"D423": []string{"delighted"},
	"D563": []string{"downhearted"},
	"D620": []string{"drowsy"},
	"D652": []string{"daring"},
	"E223": []string{"excited"},
	"E532": []string{"enthusiastic"},
	"E562": []string{"energetic"},
	"F623": []string{"frightened"},
	"F642": []string{"fearless"},
	"G430": []string{"guilty"},
	"H100": []string{"happy"},
	"H234": []string{"hostile"},
	"I631": []string{"irritable"},
	"J140": []string{"joyful"},
	"J360": []string{"jittery"},
	"L140": []string{"lively"},
	"L352": []string{"loathing"},
	"L540": []string{"lonely"},
	"N612": []string{"nervous"},
	"P630": []string{"proud"},
	"R423": []string{"relaxed"},
	"S000": []string{"shy"},
	"S120": []string{"sheepish"},
	"S200": []string{"shaky"},
	"S263": []string{"scared"},
	"S265": []string{"scornful"},
	"S300": []string{"sad"},
	"S365": []string{"strong"},
	"S410": []string{"sleepy"},
	"S422": []string{"sluggish"},
	"S616": []string{"surprised"},
	"T530": []string{"timid"},
	"T630": []string{"tired"},
}
