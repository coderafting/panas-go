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
	"happy":                  {Category: "jovility", Direction: "positive"},
	"joyful":                 {Category: "jovility", Direction: "positive"},
	"delighted":              {Category: "jovility", Direction: "positive"},
	"cheerful":               {Category: "jovility", Direction: "positive"},
	"excited":                {Category: "jovility", Direction: "positive"},
	"enthusiastic":           {Category: "jovility", Direction: "positive"},
	"lively":                 {Category: "jovility", Direction: "positive"},
	"energetic":              {Category: "jovility", Direction: "positive"},
	"proud":                  {Category: "selfAssurance", Direction: "positive"},
	"strong":                 {Category: "selfAssurance", Direction: "positive"},
	"confident":              {Category: "selfAssurance", Direction: "positive"},
	"bold":                   {Category: "selfAssurance", Direction: "positive"},
	"daring":                 {Category: "selfAssurance", Direction: "positive"},
	"fearless":               {Category: "selfAssurance", Direction: "positive"},
	"alert":                  {Category: "attentiveness", Direction: "positive"},
	"attentiveness":          {Category: "attentiveness", Direction: "positive"},
	"concentrating":          {Category: "attentiveness", Direction: "positive"},
	"determined":             {Category: "attentiveness", Direction: "positive"},
	"afraid":                 {Category: "fear", Direction: "negative"},
	"scared":                 {Category: "fear", Direction: "negative"},
	"frightened":             {Category: "fear", Direction: "negative"},
	"nervous":                {Category: "fear", Direction: "negative"},
	"jittery":                {Category: "fear", Direction: "negative"},
	"shaky":                  {Category: "fear", Direction: "negative"},
	"angry":                  {Category: "hostility", Direction: "negative"},
	"hostile":                {Category: "hostility", Direction: "negative"},
	"irritable":              {Category: "hostility", Direction: "negative"},
	"scornful":               {Category: "hostility", Direction: "negative"},
	"disgusted":              {Category: "hostility", Direction: "negative"},
	"loathing":               {Category: "hostility", Direction: "negative"},
	"guilty":                 {Category: "guilt", Direction: "negative"},
	"ashamed":                {Category: "guilt", Direction: "negative"},
	"blameworthy":            {Category: "guilt", Direction: "negative"},
	"angry at self":          {Category: "guilt", Direction: "negative"},
	"disgusted with self":    {Category: "guilt", Direction: "negative"},
	"dissatisfied with self": {Category: "guilt", Direction: "negative"},
	"sad":                    {Category: "sadness", Direction: "negative"},
	"blue":                   {Category: "sadness", Direction: "negative"},
	"downhearted":            {Category: "sadness", Direction: "negative"},
	"alone":                  {Category: "sadness", Direction: "negative"},
	"lonely":                 {Category: "sadness", Direction: "negative"},
	"shy":                    {Category: "shyness", Direction: "other"},
	"bashful":                {Category: "shyness", Direction: "other"},
	"sheepish":               {Category: "shyness", Direction: "other"},
	"timid":                  {Category: "shyness", Direction: "other"},
	"sleepy":                 {Category: "fatigue", Direction: "other"},
	"tired":                  {Category: "fatigue", Direction: "other"},
	"sluggish":               {Category: "fatigue", Direction: "other"},
	"drowsy":                 {Category: "fatigue", Direction: "other"},
	"calm":                   {Category: "serenity", Direction: "other"},
	"relaxed":                {Category: "serenity", Direction: "other"},
	"at ease":                {Category: "serenity", Direction: "other"},
	"amazed":                 {Category: "surprise", Direction: "other"},
	"surprised":              {Category: "surprise", Direction: "other"},
	"astonished":             {Category: "surprise", Direction: "other"},
}

// GeneralPosNegStates is a map of general states and their corresponding categories and overall positive/negative emotion.
var GeneralPosNegStates = map[string]StateC{
	"active":       {Category: "general", Direction: "positive"},
	"alert":        {Category: "general", Direction: "positive"},
	"attentive":    {Category: "general", Direction: "positive"},
	"determined":   {Category: "general", Direction: "positive"},
	"enthusiastic": {Category: "general", Direction: "positive"},
	"excited":      {Category: "general", Direction: "positive"},
	"inspired":     {Category: "general", Direction: "positive"},
	"interested":   {Category: "general", Direction: "positive"},
	"proud":        {Category: "general", Direction: "positive"},
	"strong":       {Category: "general", Direction: "positive"},
	"afraid":       {Category: "general", Direction: "negative"},
	"scared":       {Category: "general", Direction: "negative"},
	"nervous":      {Category: "general", Direction: "negative"},
	"jittery":      {Category: "general", Direction: "negative"},
	"irritable":    {Category: "general", Direction: "negative"},
	"hostile":      {Category: "general", Direction: "negative"},
	"guilty":       {Category: "general", Direction: "negative"},
	"ashamed":      {Category: "general", Direction: "negative"},
	"upset":        {Category: "general", Direction: "negative"},
	"distressed":   {Category: "general", Direction: "negative"},
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
