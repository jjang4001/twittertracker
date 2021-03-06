package parser

// Stopwords map
var Stopwords = map[string]bool{
	"i":          true,
	"i'm":        true,
	"r":          true,
	"gotta":      true,
	"hey":        true,
	"me":         true,
	"my":         true,
	"myself":     true,
	"we":         true,
	"we've":      true,
	"rt":         true,
	"our":        true,
	"ours":       true,
	"ourselves":  true,
	"you":        true,
	"you're":     true,
	"you've":     true,
	"you'll":     true,
	"you'd":      true,
	"your":       true,
	"yours":      true,
	"yourself":   true,
	"yourselves": true,
	"he":         true,
	"him":        true,
	"his":        true,
	"himself":    true,
	"she":        true,
	"she's":      true,
	"her":        true,
	"hers":       true,
	"herself":    true,
	"it":         true,
	"it's":       true,
	"its":        true,
	"itself":     true,
	"they":       true,
	"them":       true,
	"their":      true,
	"theirs":     true,
	"themselves": true,
	"what":       true,
	"which":      true,
	"who":        true,
	"whom":       true,
	"this":       true,
	"that":       true,
	"that'll":    true,
	"these":      true,
	"those":      true,
	"am":         true,
	"is":         true,
	"are":        true,
	"was":        true,
	"were":       true,
	"be":         true,
	"been":       true,
	"being":      true,
	"have":       true,
	"has":        true,
	"had":        true,
	"having":     true,
	"do":         true,
	"does":       true,
	"did":        true,
	"doing":      true,
	"a":          true,
	"an":         true,
	"the":        true,
	"and":        true,
	"but":        true,
	"if":         true,
	"or":         true,
	"because":    true,
	"as":         true,
	"until":      true,
	"while":      true,
	"of":         true,
	"at":         true,
	"by":         true,
	"for":        true,
	"with":       true,
	"about":      true,
	"against":    true,
	"between":    true,
	"into":       true,
	"through":    true,
	"during":     true,
	"before":     true,
	"after":      true,
	"above":      true,
	"below":      true,
	"to":         true,
	"from":       true,
	"up":         true,
	"down":       true,
	"in":         true,
	"out":        true,
	"on":         true,
	"off":        true,
	"over":       true,
	"under":      true,
	"again":      true,
	"further":    true,
	"then":       true,
	"once":       true,
	"here":       true,
	"there":      true,
	"when":       true,
	"where":      true,
	"why":        true,
	"how":        true,
	"all":        true,
	"any":        true,
	"both":       true,
	"each":       true,
	"few":        true,
	"more":       true,
	"most":       true,
	"other":      true,
	"some":       true,
	"such":       true,
	"no":         true,
	"nor":        true,
	"not":        true,
	"only":       true,
	"own":        true,
	"same":       true,
	"so":         true,
	"than":       true,
	"too":        true,
	"very":       true,
	"s":          true,
	"t":          true,
	"can":        true,
	"will":       true,
	"just":       true,
	"don":        true,
	"don't":      true,
	"should":     true,
	"should've":  true,
	"now":        true,
	"d":          true,
	"ll":         true,
	"m":          true,
	"o":          true,
	"re":         true,
	"ve":         true,
	"y":          true,
	"ain":        true,
	"aren":       true,
	"aren't":     true,
	"couldn":     true,
	"couldn't":   true,
	"didn":       true,
	"didn't":     true,
	"doesn":      true,
	"doesn't":    true,
	"hadn":       true,
	"hadn't":     true,
	"hasn":       true,
	"hasn't":     true,
	"haven":      true,
	"haven't":    true,
	"isn":        true,
	"isn't":      true,
	"ma":         true,
	"mightn":     true,
	"mightn't":   true,
	"mustn":      true,
	"mustn't":    true,
	"needn":      true,
	"needn't":    true,
	"shan":       true,
	"shan't":     true,
	"shouldn":    true,
	"shouldn't":  true,
	"wasn":       true,
	"wasn't":     true,
	"weren":      true,
	"weren't":    true,
	"won":        true,
	"won't":      true,
	"wouldn":     true,
	"wouldn't":   true,
	"\t":         true,
	"\n":         true,
	"":           true}
