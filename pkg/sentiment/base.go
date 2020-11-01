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

type state struct {
	category  string
	direction string
}

// StatesCategories is a map of states and their corresponding categories and overall positive/negative emotion.
var StatesCategories = map[string]state{
	"happy":                  state{category: "jovility", direction: "positive"},
	"joyful":                 state{category: "jovility", direction: "positive"},
	"delighted":              state{category: "jovility", direction: "positive"},
	"cheerful":               state{category: "jovility", direction: "positive"},
	"excited":                state{category: "jovility", direction: "positive"},
	"enthusiastic":           state{category: "jovility", direction: "positive"},
	"lively":                 state{category: "jovility", direction: "positive"},
	"energetic":              state{category: "jovility", direction: "positive"},
	"proud":                  state{category: "selfAssurance", direction: "positive"},
	"strong":                 state{category: "selfAssurance", direction: "positive"},
	"confident":              state{category: "selfAssurance", direction: "positive"},
	"bold":                   state{category: "selfAssurance", direction: "positive"},
	"daring":                 state{category: "selfAssurance", direction: "positive"},
	"fearless":               state{category: "selfAssurance", direction: "positive"},
	"alert":                  state{category: "attentiveness", direction: "positive"},
	"attentiveness":          state{category: "attentiveness", direction: "positive"},
	"concentrating":          state{category: "attentiveness", direction: "positive"},
	"determined":             state{category: "attentiveness", direction: "positive"},
	"afraid":                 state{category: "fear", direction: "negative"},
	"scared":                 state{category: "fear", direction: "negative"},
	"frightened":             state{category: "fear", direction: "negative"},
	"nervous":                state{category: "fear", direction: "negative"},
	"jittery":                state{category: "fear", direction: "negative"},
	"shaky":                  state{category: "fear", direction: "negative"},
	"angry":                  state{category: "hostility", direction: "negative"},
	"hostile":                state{category: "hostility", direction: "negative"},
	"irritable":              state{category: "hostility", direction: "negative"},
	"scornful":               state{category: "hostility", direction: "negative"},
	"disgusted":              state{category: "hostility", direction: "negative"},
	"loathing":               state{category: "hostility", direction: "negative"},
	"guilty":                 state{category: "guilt", direction: "negative"},
	"ashamed":                state{category: "guilt", direction: "negative"},
	"blameworthy":            state{category: "guilt", direction: "negative"},
	"angry at self":          state{category: "guilt", direction: "negative"},
	"disgusted with self":    state{category: "guilt", direction: "negative"},
	"dissatisfied with self": state{category: "guilt", direction: "negative"},
	"sad":                    state{category: "sadness", direction: "negative"},
	"blue":                   state{category: "sadness", direction: "negative"},
	"downhearted":            state{category: "sadness", direction: "negative"},
	"alone":                  state{category: "sadness", direction: "negative"},
	"lonely":                 state{category: "sadness", direction: "negative"},
	"shy":                    state{category: "shyness", direction: "other"},
	"bashful":                state{category: "shyness", direction: "other"},
	"sheepish":               state{category: "shyness", direction: "other"},
	"timid":                  state{category: "shyness", direction: "other"},
	"sleepy":                 state{category: "fatigue", direction: "other"},
	"tired":                  state{category: "fatigue", direction: "other"},
	"sluggish":               state{category: "fatigue", direction: "other"},
	"drowsy":                 state{category: "fatigue", direction: "other"},
	"calm":                   state{category: "serenity", direction: "other"},
	"relaxed":                state{category: "serenity", direction: "other"},
	"at ease":                state{category: "serenity", direction: "other"},
	"amazed":                 state{category: "surprise", direction: "other"},
	"surprised":              state{category: "surprise", direction: "other"},
	"astonished":             state{category: "surprise", direction: "other"},
}

// GeneralPosNegStates is a map of general states and their corresponding categories and overall positive/negative emotion.
var GeneralPosNegStates = map[string]state{
	"active":       state{category: "general", direction: "positive"},
	"alert":        state{category: "general", direction: "positive"},
	"attentive":    state{category: "general", direction: "positive"},
	"determined":   state{category: "general", direction: "positive"},
	"enthusiastic": state{category: "general", direction: "positive"},
	"excited":      state{category: "general", direction: "positive"},
	"inspired":     state{category: "general", direction: "positive"},
	"interested":   state{category: "general", direction: "positive"},
	"proud":        state{category: "general", direction: "positive"},
	"strong":       state{category: "general", direction: "positive"},
	"afraid":       state{category: "general", direction: "negative"},
	"scared":       state{category: "general", direction: "negative"},
	"nervous":      state{category: "general", direction: "negative"},
	"jittery":      state{category: "general", direction: "negative"},
	"irritable":    state{category: "general", direction: "negative"},
	"hostile":      state{category: "general", direction: "negative"},
	"guilty":       state{category: "general", direction: "negative"},
	"ashamed":      state{category: "general", direction: "negative"},
	"upset":        state{category: "general", direction: "negative"},
	"distressed":   state{category: "general", direction: "negative"},
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
