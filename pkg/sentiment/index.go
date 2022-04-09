package sentiment

import (
	"github.com/coderafting/panas-go/internal/text"
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
	"I000": {"I"},
	"I500": {"I'm", "I am"},
	"A500": {"am"},
	"F452": {"feeling"},
	"M000": {"me"},
	"M241": {"myself"},
}

// StatesSoundexIndex is a map of Soundex codes and their corresponding state strings.
var StatesSoundexIndex = map[string][]string{
	"A163": {"afraid"},
	"A235": {"astonished"},
	"A253": {"ashamed"},
	"A320": {"at ease"},
	"A353": {"attentiveness"},
	"A450": {"alone"},
	"A463": {"alert"},
	"A523": {"amazed"},
	"A526": {"angry", "angry at self"},
	"B214": {"bashful"},
	"B400": {"blue"},
	"B430": {"bold"},
	"B456": {"blameworthy"},
	"C450": {"calm"},
	"C513": {"confident"},
	"C525": {"concentrating"},
	"C614": {"cheerful"},
	"D222": {"disgusted", "disgusted with self"},
	"D232": {"dissatisfied with self"},
	"D365": {"determined"},
	"D423": {"delighted"},
	"D563": {"downhearted"},
	"D620": {"drowsy"},
	"D652": {"daring"},
	"E223": {"excited"},
	"E532": {"enthusiastic"},
	"E562": {"energetic"},
	"F623": {"frightened"},
	"F642": {"fearless"},
	"G430": {"guilty"},
	"H100": {"happy"},
	"H234": {"hostile"},
	"I631": {"irritable"},
	"J140": {"joyful"},
	"J360": {"jittery"},
	"L140": {"lively"},
	"L352": {"loathing"},
	"L540": {"lonely"},
	"N612": {"nervous"},
	"P630": {"proud"},
	"R423": {"relaxed"},
	"S000": {"shy"},
	"S120": {"sheepish"},
	"S200": {"shaky"},
	"S263": {"scared"},
	"S265": {"scornful"},
	"S300": {"sad"},
	"S365": {"strong"},
	"S410": {"sleepy"},
	"S422": {"sluggish"},
	"S616": {"surprised"},
	"T530": {"timid"},
	"T630": {"tired"},
}
