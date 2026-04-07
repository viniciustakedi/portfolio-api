package flashcards

func seedPaths() []PathDocument {
	return []PathDocument{
		{Language: "en", Level: "beginner", Title: "First Steps", Description: "Essential phrasal verbs for everyday communication", Order: 1, TotalCards: 10, Icon: "🌱"},
		{Language: "en", Level: "intermediate", Title: "Building fluency", Description: "Trickier phrasal verbs and collocations", Order: 2, TotalCards: 10, Icon: "🌿"},
		{Language: "en", Level: "advanced", Title: "Native edges", Description: "Idioms and nuanced expressions", Order: 3, TotalCards: 10, Icon: "🌳"},
		{Language: "es", Level: "beginner", Title: "Primeros pasos", Description: "Falsos amigos y matices útiles para cualquier hablante de L2", Order: 1, TotalCards: 10, Icon: "🌱"},
		{Language: "es", Level: "intermediate", Title: "Más allá del básico", Description: "Uso real en conversación y registro", Order: 2, TotalCards: 10, Icon: "🌿"},
		{Language: "es", Level: "advanced", Title: "Matices", Description: "Expresiones finas y registro formal", Order: 3, TotalCards: 10, Icon: "🌳"},
	}
}

// mkCard builds a seed card. synonyms are English (or Spanish) near-synonyms / related words—no single L1 gloss.
func mkCard(lang, path string, diff int, word, typ, desc string, synonyms []string, tags []string, ex []string) FlashcardDocument {
	if tags == nil {
		tags = []string{}
	}
	return FlashcardDocument{
		Word: word, Synonyms: synonyms, Type: typ, Language: lang, Path: path,
		Difficulty: diff, Description: desc, Examples: ex, Tags: tags,
	}
}

func seedCards() []FlashcardDocument {
	var out []FlashcardDocument
	out = append(out, enBeginner()...)
	out = append(out, enIntermediate()...)
	out = append(out, enAdvanced()...)
	out = append(out, esBeginner()...)
	out = append(out, esIntermediate()...)
	out = append(out, esAdvanced()...)
	return out
}

func enBeginner() []FlashcardDocument {
	return []FlashcardDocument{
		mkCard("en", "beginner", 1, "fill up", "phrasal_verb", "To fill something completely to the top, or to make something full.",
			[]string{"top off", "replenish", "make full"}, []string{"daily-use", "travel"},
			[]string{"Can you fill up the gas tank before we leave?", "She filled up her journal with beautiful drawings."}),
		mkCard("en", "beginner", 1, "look up", "phrasal_verb", "To search for information, often in a book or online.",
			[]string{"search for", "find", "check"}, []string{"study"},
			[]string{"Look up the word in the dictionary.", "I'll look up the train times for you."}),
		mkCard("en", "beginner", 2, "give up", "phrasal_verb", "To stop trying or stop doing something.",
			[]string{"quit", "abandon", "surrender"}, []string{"motivation"},
			[]string{"Don't give up on your dreams.", "He gave up smoking last year."}),
		mkCard("en", "beginner", 2, "pick up", "phrasal_verb", "To collect someone or something, or to learn a skill casually.",
			[]string{"collect", "fetch", "acquire casually"}, []string{"travel", "daily-use"},
			[]string{"I picked up some Spanish on my trip.", "I'll pick you up at eight."}),
		mkCard("en", "beginner", 2, "turn off", "phrasal_verb", "To stop a machine or light from working.",
			[]string{"switch off", "power down", "shut off"}, []string{"home"},
			[]string{"Turn off the lights when you leave.", "She turned off the alarm."}),
		mkCard("en", "beginner", 3, "put off", "phrasal_verb", "To delay something to a later time.",
			[]string{"postpone", "defer", "delay"}, []string{"work"},
			[]string{"Let's not put off the dentist any longer.", "Rain put off the match."}),
		mkCard("en", "beginner", 3, "get up", "phrasal_verb", "To leave your bed and start your day.",
			[]string{"rise", "wake and rise", "leave the bed"}, []string{"daily-use"},
			[]string{"I get up at six on weekdays.", "Get up slowly if you feel dizzy."}),
		mkCard("en", "beginner", 3, "come back", "phrasal_verb", "To return to a place.",
			[]string{"return", "go back", "revisit"}, []string{"travel"},
			[]string{"Come back soon!", "What time did you come back last night?"}),
		mkCard("en", "beginner", 4, "run out of", "phrasal_verb", "To use all of something so none is left.",
			[]string{"use up", "be depleted", "have none left"}, []string{"daily-use"},
			[]string{"We've run out of milk.", "Don't run out of patience."}),
		mkCard("en", "beginner", 4, "wake up", "phrasal_verb", "To stop sleeping.",
			[]string{"awaken", "rouse", "come awake"}, []string{"daily-use"},
			[]string{"I wake up to birdsong.", "Wake me up at seven, please."}),
	}
}

func enIntermediate() []FlashcardDocument {
	return []FlashcardDocument{
		mkCard("en", "intermediate", 2, "come across", "phrasal_verb", "To find by chance, or to give an impression.",
			[]string{"stumble upon", "seem", "appear"}, []string{"conversation"},
			[]string{"I came across an old photo.", "She comes across as very confident."}),
		mkCard("en", "intermediate", 2, "look forward to", "phrasal_verb", "To be excited about a future event.",
			[]string{"anticipate", "await eagerly", "can't wait for"}, []string{"email", "social"},
			[]string{"I'm looking forward to the weekend.", "We look forward to hearing from you."}),
		mkCard("en", "intermediate", 3, "put up with", "phrasal_verb", "To tolerate something unpleasant.",
			[]string{"tolerate", "endure", "bear"}, []string{"relationships"},
			[]string{"How do you put up with the noise?", "I won't put up with rudeness."}),
		mkCard("en", "intermediate", 3, "get over", "phrasal_verb", "To recover from illness or emotional difficulty.",
			[]string{"recover from", "move past", "bounce back from"}, []string{"health"},
			[]string{"It took weeks to get over the flu.", "She never got over losing her dog."}),
		mkCard("en", "intermediate", 3, "break down", "phrasal_verb", "When a machine stops working, or someone loses emotional control.",
			[]string{"fail", "collapse", "lose composure"}, []string{"cars", "emotions"},
			[]string{"The car broke down on the highway.", "He broke down in tears."}),
		mkCard("en", "intermediate", 4, "carry on", "phrasal_verb", "To continue doing something.",
			[]string{"continue", "keep going", "persist"}, []string{"work"},
			[]string{"Carry on without me.", "They carried on talking all night."}),
		mkCard("en", "intermediate", 4, "end up", "phrasal_verb", "To finally be in a situation, often unexpectedly.",
			[]string{"wind up", "finish up", "land up"}, []string{"storytelling"},
			[]string{"We ended up staying home.", "How did you end up in Lisbon?"}),
		mkCard("en", "intermediate", 4, "figure out", "phrasal_verb", "To understand or solve something.",
			[]string{"work out", "solve", "grasp"}, []string{"problem-solving"},
			[]string{"I can't figure out this puzzle.", "She figured out the answer quickly."}),
		mkCard("en", "intermediate", 5, "give in", "phrasal_verb", "To finally agree after resisting.",
			[]string{"yield", "relent", "concede"}, []string{"negotiation"},
			[]string{"He gave in to pressure.", "Don't give in too easily."}),
		mkCard("en", "intermediate", 5, "turn down", "phrasal_verb", "To reject an offer, or lower volume/heat.",
			[]string{"reject", "refuse", "lower (volume)"}, []string{"social", "home"},
			[]string{"She turned down the job.", "Could you turn down the music?"}),
	}
}

func enAdvanced() []FlashcardDocument {
	return []FlashcardDocument{
		mkCard("en", "advanced", 2, "take with a grain of salt", "expression", "Be skeptical; don't believe everything completely.",
			[]string{"be skeptical", "don't take literally", "question the source"}, []string{"media"},
			[]string{"Take rumors with a grain of salt.", "I'd take that review with a grain of salt."}),
		mkCard("en", "advanced", 2, "beat around the bush", "expression", "Avoid speaking directly about a sensitive topic.",
			[]string{"evade the point", "talk indirectly", "hem and haw"}, []string{"conversation"},
			[]string{"Stop beating around the bush—what happened?", "She beat around the bush for ten minutes."}),
		mkCard("en", "advanced", 3, "cut corners", "expression", "Do something poorly or cheaply to save time or money.",
			[]string{"skimp", "take shortcuts", "do the minimum"}, []string{"work"},
			[]string{"Cutting corners on safety is dangerous.", "They cut corners and the paint peeled."}),
		mkCard("en", "advanced", 3, "get the ball rolling", "expression", "Start an activity or process.",
			[]string{"kick off", "start things off", "initiate"}, []string{"meetings"},
			[]string{"I'll ask a question to get the ball rolling.", "A joke got the ball rolling."}),
		mkCard("en", "advanced", 3, "read between the lines", "expression", "Understand implied meaning, not only literal words.",
			[]string{"infer", "sense the subtext", "read the implication"}, []string{"literature"},
			[]string{"Read between the lines—she's unhappy.", "The email was polite, but read between the lines."}),
		mkCard("en", "advanced", 4, "hit the nail on the head", "expression", "Describe something exactly right.",
			[]string{"get it exactly right", "sum it up perfectly", "spot-on"}, []string{"feedback"},
			[]string{"You hit the nail on the head with that comment.", "Her analysis hit the nail on the head."}),
		mkCard("en", "advanced", 4, "pull strings", "expression", "Use hidden influence to get an advantage.",
			[]string{"use connections", "leverage influence", "work behind the scenes"}, []string{"work"},
			[]string{"He pulled strings to get the ticket.", "It's not what you know, it's who pulls strings."}),
		mkCard("en", "advanced", 4, "a blessing in disguise", "expression", "Something bad that later brings good results.",
			[]string{"silver lining", "hidden benefit", "good from bad luck"}, []string{"life"},
			[]string{"Losing that job was a blessing in disguise.", "The delay was a blessing in disguise—we missed the storm."}),
		mkCard("en", "advanced", 5, "the tip of the iceberg", "expression", "A small visible part of a much larger problem.",
			[]string{"just the start", "small visible fraction", "surface issue"}, []string{"news"},
			[]string{"These errors are just the tip of the iceberg.", "One complaint may be the tip of the iceberg."}),
		mkCard("en", "advanced", 5, "bite off more than you can chew", "expression", "Take on more responsibility than you can handle.",
			[]string{"overcommit", "take on too much", "overextend"}, []string{"work"},
			[]string{"Don't bite off more than you can chew this semester.", "We bit off more than we could chew with the renovation."}),
	}
}

func esBeginner() []FlashcardDocument {
	return []FlashcardDocument{
		mkCard("es", "beginner", 1, "embarazada", "adjective", "Means pregnant. For 'embarrassed' use avergonzada / tener vergüenza.",
			[]string{"pregnant", "expecting"}, []string{"false-friend"},
			[]string{"Ella está embarazada de tres meses.", "No confundas embarazada con 'embarrassed' en inglés."}),
		mkCard("es", "beginner", 1, "borracha", "noun", "Often means 'drunk woman'. Rubber/eraser is goma or goma de borrar.",
			[]string{"drunk (woman)", "intoxicated"}, []string{"false-friend"},
			[]string{"No manejes si estás borracha.", "Necesito una goma de borrar, no 'una borracha'."}),
		mkCard("es", "beginner", 2, "polvo", "noun", "Dust; context can also mean marine sediment. Chicken is pollo.",
			[]string{"dust", "powder (fine particles)"}, []string{"false-friend"},
			[]string{"Limpia el polvo de los muebles.", "Para 'chicken', di pollo—no polvo."}),
		mkCard("es", "beginner", 2, "largo", "adjective", "Long in length. Big is grande.",
			[]string{"long", "lengthy"}, []string{"false-friend"},
			[]string{"Tiene el pelo largo y oscuro.", "Una mesa larga, no 'grande' por longitud."}),
		mkCard("es", "beginner", 2, "en absoluto", "expression", "Often means 'not at all'—not the same as English 'absolutely' agreeing.",
			[]string{"not at all", "by no means", "definitely not"}, []string{"false-friend"},
			[]string{"¿Te molesta? —En absoluto.", "No me gusta en absoluto."}),
		mkCard("es", "beginner", 3, "constipado", "adjective", "Having a cold. For constipation use estreñido.",
			[]string{"having a cold", "stuffed up"}, []string{"false-friend"},
			[]string{"Estoy constipado y me duele la garganta.", "Para estreñimiento: estreñido."}),
		mkCard("es", "beginner", 3, "éxito", "noun", "Success. Exit is salida.",
			[]string{"success", "hit", "achievement"}, []string{"false-friend"},
			[]string{"Tuvo mucho éxito con su libro.", "La salida está allí—no 'el éxito'."}),
		mkCard("es", "beginner", 3, "actual", "adjective", "Current, present—not 'actual' as in English 'real'.",
			[]string{"current", "present-day", "today's"}, []string{"false-friend"},
			[]string{"El presidente actual hablará hoy.", "En realidad = really / actually (different word)."}),
		mkCard("es", "beginner", 4, "carpeta", "noun", "Folder. Carpet is alfombra or moqueta.",
			[]string{"folder", "file (binder)"}, []string{"false-friend"},
			[]string{"Guarda el contrato en la carpeta.", "La alfombra es suave—no carpeta."}),
		mkCard("es", "beginner", 4, "emocionado", "adjective", "Excited or moved; register can overlap with 'nervous' in intensity.",
			[]string{"excited", "moved", "thrilled"}, []string{"nuance"},
			[]string{"Estoy emocionado por el viaje.", "Estaba tan emocionado que temblaba."}),
	}
}

func esIntermediate() []FlashcardDocument {
	return []FlashcardDocument{
		mkCard("es", "intermediate", 2, "ropa", "noun", "Clothing. Rope is cuerda or soga.",
			[]string{"clothes", "clothing", "garments"}, []string{"vocabulary"},
			[]string{"Colgué la ropa en el tendedero.", "Necesito comprar ropa de invierno."}),
		mkCard("es", "intermediate", 2, "sopa", "noun", "Soup—not soap (jabón).",
			[]string{"soup", "broth"}, []string{"false-friend"},
			[]string{"La sopa de lentejas está deliciosa.", "El jabón está en el baño, no en el plato."}),
		mkCard("es", "intermediate", 3, "grabar", "verb", "To record audio/video. Engrave can be grabar too—context clarifies.",
			[]string{"record", "tape", "film"}, []string{"media"},
			[]string{"Voy a grabar la entrevista.", "Grabaron un disco en vivo."}),
		mkCard("es", "intermediate", 3, "asistir", "verb", "To attend an event. To help is ayudar.",
			[]string{"attend", "be present at"}, []string{"false-friend"},
			[]string{"Asistiré a la conferencia mañana.", "¿Puedes ayudarme?—no 'asistirme'."}),
		mkCard("es", "intermediate", 3, "contestar", "verb", "To answer. To contest/dispute is impugnar or disputar.",
			[]string{"answer", "reply", "pick up (phone)"}, []string{"false-friend"},
			[]string{"Contesta el teléfono, por favor.", "Contestó con educación a la pregunta difícil."}),
		mkCard("es", "intermediate", 4, "carrera", "noun", "University degree track or a race—context.",
			[]string{"degree", "major", "race (sport)"}, []string{"education"},
			[]string{"Estudia la carrera de medicina.", "Ganó la carrera de los cien metros."}),
		mkCard("es", "intermediate", 4, "ésta / esta", "determiner", "This; accent on demonstratives is largely obsolete in modern print.",
			[]string{"this", "these (f.)"}, []string{"spelling"},
			[]string{"Esta ciudad me encanta.", "Esto es importante para mí."}),
		mkCard("es", "intermediate", 4, "por favor", "expression", "Please—word order is flexible in requests.",
			[]string{"please", "if you would"}, []string{"politeness"},
			[]string{"Pase, por favor.", "¿Podrías repetir, por favor?"}),
		mkCard("es", "intermediate", 5, "subir", "verb", "To go up, upload, or increase.",
			[]string{"go up", "rise", "upload"}, []string{"daily-use"},
			[]string{"Subimos la montaña al amanecer.", "Subieron los precios del transporte."}),
		mkCard("es", "intermediate", 5, "vale", "expression", "Informal OK / deal—common in Spain.",
			[]string{"OK", "sure", "deal"}, []string{"conversation"},
			[]string{"—¿Vemos a las ocho? —Vale.", "Vale, lo dejamos así."}),
	}
}

func esAdvanced() []FlashcardDocument {
	return []FlashcardDocument{
		mkCard("es", "advanced", 2, "ser pan comido", "expression", "Something very easy.",
			[]string{"a piece of cake", "dead easy", "no sweat"}, []string{"idiom"},
			[]string{"Para ella, el examen fue pan comido.", "Con práctica, esto será pan comido."}),
		mkCard("es", "advanced", 2, "echar una mano", "expression", "To help someone informally.",
			[]string{"lend a hand", "help out", "give a hand"}, []string{"idiom"},
			[]string{"¿Te echo una mano con las cajas?", "Me echó una mano cuando más lo necesitaba."}),
		mkCard("es", "advanced", 3, "quedarse en nada", "expression", "When plans fail to materialize.",
			[]string{"fizzle out", "come to nothing", "fall through"}, []string{"idiom"},
			[]string{"El proyecto se quedó en nada.", "La reunión quedó en nada tras el cambio de jefe."}),
		mkCard("es", "advanced", 3, "tener mano izquierda", "expression", "Skill in handling people tactfully.",
			[]string{"tact", "diplomacy", "people skills"}, []string{"work"},
			[]string{"Un buen líder tiene mano izquierda.", "Falta mano izquierda en ese feedback."}),
		mkCard("es", "advanced", 3, "no tener pelos en la lengua", "expression", "To speak bluntly.",
			[]string{"speak plainly", "not mince words", "be outspoken"}, []string{"personality"},
			[]string{"Ella no tiene pelos en la lengua.", "Me gusta la gente sin pelos en la lengua."}),
		mkCard("es", "advanced", 4, "más vale tarde que nunca", "expression", "Better late than never.",
			[]string{"better late than never"}, []string{"proverb"},
			[]string{"Llegó más vale tarde que nunca.", "Aprender idiomas: más vale tarde que nunca."}),
		mkCard("es", "advanced", 4, "a buen entendedor, pocas palabras bastan", "expression", "A hint is enough for a smart listener.",
			[]string{"a word to the wise", "enough said"}, []string{"proverb"},
			[]string{"A buen entendedor, pocas palabras.", "No diré más: a buen entendedor..."}),
		mkCard("es", "advanced", 4, "estar en el aire", "expression", "Uncertain, not decided yet.",
			[]string{"up in the air", "unsettled", "pending"}, []string{"work"},
			[]string{"El contrato sigue en el aire.", "Todo quedó en el aire tras la reunión."}),
		mkCard("es", "advanced", 5, "dar largas", "expression", "To stall or put something off repeatedly.",
			[]string{"stall", "string along", "drag one's feet"}, []string{"idiom"},
			[]string{"Me dan largas con el reembolso.", "Deja de dar largas y decide."}),
		mkCard("es", "advanced", 5, "ponerse las pilas", "expression", "To get serious and focus.",
			[]string{"get it together", "shape up", "focus up"}, []string{"informal"},
			[]string{"Ponte las pilas con el estudio.", "Hay que ponerse las pilas este mes."}),
	}
}
