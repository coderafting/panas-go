package sentiment

/*
Base data, as per the paper "PANAS-t: A Pychometric Scale for Measuring Sentiments on Twitter",
which can be accessed at https://arxiv.org/abs/1308.1857.
*/

// SelfReferences indicates the references to the subject in a sentence, as recognized by the PANAS-t paper.
var SelfReferences = []string{"I am", "I'm", "I", "am", "feeling", "me", "myself"}

// StatesColl is a collection of all the specific sentiment states that can appear in a text,
// as recognized by the PANAS-t paper.
var StatesColl = []string{
	// jovility states
	"happy", "joyful", "delighted", "cheerful", "excited", "enthusiastic", "lively", "energetic",
	// selfAssurance states
	"proud", "strong", "confident", "bold", "daring", "fearless",
	// attentiveness states
	"alert", "attentiveness", "concentrating", "determined",
	// fear states
	"afraid", "scared", "frightened", "nervous", "jittery", "shaky",
	// hostility states
	"angry", "hostile", "irritable", "scornful", "disgusted", "loathing",
	// guilt states
	"guilty", "ashamed", "blameworthy", "angry at self", "disgusted with self", "dissatisfied with self",
	// sadness states
	"sad", "blue", "downhearted", "alone", "lonely",
	// shyness states
	"shy", "bashful", "sheepish", "timid",
	// fatigue states
	"sleepy", "tired", "sluggish", "drowsy",
	// serenity states
	"calm", "relaxed", "at ease",
	// surprise states
	"amazed", "surprised", "astonished",
}

// CategoriesMap is map of available sentiment categories, as recognized by the pANAS-t paper
var CategoriesMap = map[string]bool{
	"jovility":      true,
	"selfAssurance": true,
	"attentiveness": true,
	"fear":          true,
	"hostility":     true,
	"guilt":         true,
	"sadness":       true,
	"shyness":       true,
	"fatigue":       true,
	"serenity":      true,
	"surprise":      true,
}

// StateC contains a its specific state-category and its positive/negative direction.
type StateC struct {
	Category  string
	Direction string
}

// StatesCategories is a map of states and their corresponding categories and overall positive/negative emotion.
var StatesCategories = map[string]StateC{
	"happy":                  StateC{Category: "jovility", Direction: "positive"},
	"joyful":                 StateC{Category: "jovility", Direction: "positive"},
	"delighted":              StateC{Category: "jovility", Direction: "positive"},
	"cheerful":               StateC{Category: "jovility", Direction: "positive"},
	"excited":                StateC{Category: "jovility", Direction: "positive"},
	"enthusiastic":           StateC{Category: "jovility", Direction: "positive"},
	"lively":                 StateC{Category: "jovility", Direction: "positive"},
	"energetic":              StateC{Category: "jovility", Direction: "positive"},
	"proud":                  StateC{Category: "selfAssurance", Direction: "positive"},
	"strong":                 StateC{Category: "selfAssurance", Direction: "positive"},
	"confident":              StateC{Category: "selfAssurance", Direction: "positive"},
	"bold":                   StateC{Category: "selfAssurance", Direction: "positive"},
	"daring":                 StateC{Category: "selfAssurance", Direction: "positive"},
	"fearless":               StateC{Category: "selfAssurance", Direction: "positive"},
	"alert":                  StateC{Category: "attentiveness", Direction: "positive"},
	"attentiveness":          StateC{Category: "attentiveness", Direction: "positive"},
	"concentrating":          StateC{Category: "attentiveness", Direction: "positive"},
	"determined":             StateC{Category: "attentiveness", Direction: "positive"},
	"afraid":                 StateC{Category: "fear", Direction: "negative"},
	"scared":                 StateC{Category: "fear", Direction: "negative"},
	"frightened":             StateC{Category: "fear", Direction: "negative"},
	"nervous":                StateC{Category: "fear", Direction: "negative"},
	"jittery":                StateC{Category: "fear", Direction: "negative"},
	"shaky":                  StateC{Category: "fear", Direction: "negative"},
	"angry":                  StateC{Category: "hostility", Direction: "negative"},
	"hostile":                StateC{Category: "hostility", Direction: "negative"},
	"irritable":              StateC{Category: "hostility", Direction: "negative"},
	"scornful":               StateC{Category: "hostility", Direction: "negative"},
	"disgusted":              StateC{Category: "hostility", Direction: "negative"},
	"loathing":               StateC{Category: "hostility", Direction: "negative"},
	"guilty":                 StateC{Category: "guilt", Direction: "negative"},
	"ashamed":                StateC{Category: "guilt", Direction: "negative"},
	"blameworthy":            StateC{Category: "guilt", Direction: "negative"},
	"angry at self":          StateC{Category: "guilt", Direction: "negative"},
	"disgusted with self":    StateC{Category: "guilt", Direction: "negative"},
	"dissatisfied with self": StateC{Category: "guilt", Direction: "negative"},
	"sad":                    StateC{Category: "sadness", Direction: "negative"},
	"blue":                   StateC{Category: "sadness", Direction: "negative"},
	"downhearted":            StateC{Category: "sadness", Direction: "negative"},
	"alone":                  StateC{Category: "sadness", Direction: "negative"},
	"lonely":                 StateC{Category: "sadness", Direction: "negative"},
	"shy":                    StateC{Category: "shyness", Direction: "other"},
	"bashful":                StateC{Category: "shyness", Direction: "other"},
	"sheepish":               StateC{Category: "shyness", Direction: "other"},
	"timid":                  StateC{Category: "shyness", Direction: "other"},
	"sleepy":                 StateC{Category: "fatigue", Direction: "other"},
	"tired":                  StateC{Category: "fatigue", Direction: "other"},
	"sluggish":               StateC{Category: "fatigue", Direction: "other"},
	"drowsy":                 StateC{Category: "fatigue", Direction: "other"},
	"calm":                   StateC{Category: "serenity", Direction: "other"},
	"relaxed":                StateC{Category: "serenity", Direction: "other"},
	"at ease":                StateC{Category: "serenity", Direction: "other"},
	"amazed":                 StateC{Category: "surprise", Direction: "other"},
	"surprised":              StateC{Category: "surprise", Direction: "other"},
	"astonished":             StateC{Category: "surprise", Direction: "other"},
}

// GeneralPosNegStates is a map of general states and their corresponding categories and overall positive/negative emotion.
var GeneralPosNegStates = map[string]StateC{
	"active":       StateC{Category: "general", Direction: "positive"},
	"alert":        StateC{Category: "general", Direction: "positive"},
	"attentive":    StateC{Category: "general", Direction: "positive"},
	"determined":   StateC{Category: "general", Direction: "positive"},
	"enthusiastic": StateC{Category: "general", Direction: "positive"},
	"excited":      StateC{Category: "general", Direction: "positive"},
	"inspired":     StateC{Category: "general", Direction: "positive"},
	"interested":   StateC{Category: "general", Direction: "positive"},
	"proud":        StateC{Category: "general", Direction: "positive"},
	"strong":       StateC{Category: "general", Direction: "positive"},
	"afraid":       StateC{Category: "general", Direction: "negative"},
	"scared":       StateC{Category: "general", Direction: "negative"},
	"nervous":      StateC{Category: "general", Direction: "negative"},
	"jittery":      StateC{Category: "general", Direction: "negative"},
	"irritable":    StateC{Category: "general", Direction: "negative"},
	"hostile":      StateC{Category: "general", Direction: "negative"},
	"guilty":       StateC{Category: "general", Direction: "negative"},
	"ashamed":      StateC{Category: "general", Direction: "negative"},
	"upset":        StateC{Category: "general", Direction: "negative"},
	"distressed":   StateC{Category: "general", Direction: "negative"},
}

// WorldBaseline is based on about 3.5 years (2009-13) of twitter data, about 0.48bn tweets.
// Please see the paper for further details.
var WorldBaseline = map[string]float64{
	// positive sentiments
	"joviality":     0.0182421,
	"selfAssurance": 0.0036012,
	"attentiveness": 0.0008997,
	// negative sentiments
	"fear":      0.0063791,
	"hostility": 0.0018225,
	"guilt":     0.0021756,
	"sadness":   0.0086279,
	// other sentiments
	"shyness":  0.0007608,
	"fatigue":  0.0240757,
	"surprise": 0.0084612,
	"serenity": 0.0022914,
	// overall (computed as the weighted averages of the positive, negative, and other sentiment values)
	"positive": 0.007581,
	"negative": 0.004751275,
	"other":    0.008897275,
}
