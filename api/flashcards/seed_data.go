package flashcards

func seedPaths() []PathDocument {
	return []PathDocument{
		{Language: "en", Level: "beginner", Title: "First Steps", Description: "Essential phrasal verbs for everyday communication", Order: 1, TotalCards: 10, Icon: "🌱"},
		{Language: "en", Level: "intermediate", Title: "Building fluency", Description: "Trickier phrasal verbs and collocations", Order: 2, TotalCards: 10, Icon: "🌿"},
		{Language: "en", Level: "advanced", Title: "Native edges", Description: "Idioms and nuanced expressions", Order: 3, TotalCards: 10, Icon: "🌳"},
		{Language: "es", Level: "beginner", Title: "Primeros pasos", Description: "Falsos amigos y trampas comunes para quien habla portugués", Order: 1, TotalCards: 10, Icon: "🌱"},
		{Language: "es", Level: "intermediate", Title: "Más allá del básico", Description: "Uso real en conversación y registro", Order: 2, TotalCards: 10, Icon: "🌿"},
		{Language: "es", Level: "advanced", Title: "Matices", Description: "Expresiones finas y registro formal", Order: 3, TotalCards: 10, Icon: "🌳"},
	}
}

func mkCard(lang, path string, diff int, word, trans, typ, desc string, tags []string, ex []string) FlashcardDocument {
	if tags == nil {
		tags = []string{}
	}
	return FlashcardDocument{
		Word: word, Translation: trans, Type: typ, Language: lang, Path: path,
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
		mkCard("en", "beginner", 1, "fill up", "encher / completar", "phrasal_verb", "To fill something completely to the top, or to make something full.", []string{"daily-use", "travel"},
			[]string{"Can you fill up the gas tank before we leave?", "She filled up her journal with beautiful drawings."}),
		mkCard("en", "beginner", 1, "look up", "buscar (informação)", "phrasal_verb", "To search for information, often in a book or online.", []string{"study"},
			[]string{"Look up the word in the dictionary.", "I'll look up the train times for you."}),
		mkCard("en", "beginner", 2, "give up", "desistir", "phrasal_verb", "To stop trying or stop doing something.", []string{"motivation"},
			[]string{"Don't give up on your dreams.", "He gave up smoking last year."}),
		mkCard("en", "beginner", 2, "pick up", "pegar / aprender informalmente", "phrasal_verb", "To collect someone or something, or to learn a skill casually.", []string{"travel", "daily-use"},
			[]string{"I picked up some Spanish on my trip.", "I'll pick you up at eight."}),
		mkCard("en", "beginner", 2, "turn off", "desligar", "phrasal_verb", "To stop a machine or light from working.", []string{"home"},
			[]string{"Turn off the lights when you leave.", "She turned off the alarm."}),
		mkCard("en", "beginner", 3, "put off", "adiar", "phrasal_verb", "To delay something to a later time.", []string{"work"},
			[]string{"Let's not put off the dentist any longer.", "Rain put off the match."}),
		mkCard("en", "beginner", 3, "get up", "levantar-se", "phrasal_verb", "To leave your bed and start your day.", []string{"daily-use"},
			[]string{"I get up at six on weekdays.", "Get up slowly if you feel dizzy."}),
		mkCard("en", "beginner", 3, "come back", "voltar", "phrasal_verb", "To return to a place.", []string{"travel"},
			[]string{"Come back soon!", "What time did you come back last night?"}),
		mkCard("en", "beginner", 4, "run out of", "ficar sem", "phrasal_verb", "To use all of something so none is left.", []string{"daily-use"},
			[]string{"We've run out of milk.", "Don't run out of patience."}),
		mkCard("en", "beginner", 4, "wake up", "acordar", "phrasal_verb", "To stop sleeping.", []string{"daily-use"},
			[]string{"I wake up to birdsong.", "Wake me up at seven, please."}),
	}
}

func enIntermediate() []FlashcardDocument {
	return []FlashcardDocument{
		mkCard("en", "intermediate", 2, "come across", "encontrar por acaso / parecer", "phrasal_verb", "To find by chance, or to give an impression.", []string{"conversation"},
			[]string{"I came across an old photo.", "She comes across as very confident."}),
		mkCard("en", "intermediate", 2, "look forward to", "ter expectativa ansiosa por", "phrasal_verb", "To be excited about a future event.", []string{"email", "social"},
			[]string{"I'm looking forward to the weekend.", "We look forward to hearing from you."}),
		mkCard("en", "intermediate", 3, "put up with", "aguentar / tolerar", "phrasal_verb", "To tolerate something unpleasant.", []string{"relationships"},
			[]string{"How do you put up with the noise?", "I won't put up with rudeness."}),
		mkCard("en", "intermediate", 3, "get over", "superar (problema, doença)", "phrasal_verb", "To recover from illness or emotional difficulty.", []string{"health"},
			[]string{"It took weeks to get over the flu.", "She never got over losing her dog."}),
		mkCard("en", "intermediate", 3, "break down", "quebrar / ter crise emocional", "phrasal_verb", "When a machine stops working, or someone loses emotional control.", []string{"cars", "emotions"},
			[]string{"The car broke down on the highway.", "He broke down in tears."}),
		mkCard("en", "intermediate", 4, "carry on", "continuar", "phrasal_verb", "To continue doing something.", []string{"work"},
			[]string{"Carry on without me.", "They carried on talking all night."}),
		mkCard("en", "intermediate", 4, "end up", "acabar por", "phrasal_verb", "To finally be in a situation, often unexpectedly.", []string{"storytelling"},
			[]string{"We ended up staying home.", "How did you end up in Lisbon?"}),
		mkCard("en", "intermediate", 4, "figure out", "descobrir / resolver", "phrasal_verb", "To understand or solve something.", []string{"problem-solving"},
			[]string{"I can't figure out this puzzle.", "She figured out the answer quickly."}),
		mkCard("en", "intermediate", 5, "give in", "ceder", "phrasal_verb", "To finally agree after resisting.", []string{"negotiation"},
			[]string{"He gave in to pressure.", "Don't give in too easily."}),
		mkCard("en", "intermediate", 5, "turn down", "recusar / abaixar (volume)", "phrasal_verb", "To reject an offer, or lower volume/heat.", []string{"social", "home"},
			[]string{"She turned down the job.", "Could you turn down the music?"}),
	}
}

func enAdvanced() []FlashcardDocument {
	return []FlashcardDocument{
		mkCard("en", "advanced", 2, "take with a grain of salt", "não levar 100% a sério", "expression", "Be skeptical; don't believe everything completely.", []string{"media"},
			[]string{"Take rumors with a grain of salt.", "I'd take that review with a grain of salt."}),
		mkCard("en", "advanced", 2, "beat around the bush", "enrolar / não ir direto ao ponto", "expression", "Avoid speaking directly about a sensitive topic.", []string{"conversation"},
			[]string{"Stop beating around the bush—what happened?", "She beat around the bush for ten minutes."}),
		mkCard("en", "advanced", 3, "cut corners", "fazer pelo mais fácil (com má qualidade)", "expression", "Do something poorly or cheaply to save time or money.", []string{"work"},
			[]string{"Cutting corners on safety is dangerous.", "They cut corners and the paint peeled."}),
		mkCard("en", "advanced", 3, "get the ball rolling", "dar o pontapé inicial", "expression", "Start an activity or process.", []string{"meetings"},
			[]string{"I'll ask a question to get the ball rolling.", "A joke got the ball rolling."}),
		mkCard("en", "advanced", 3, "read between the lines", "ler nas entrelinhas", "expression", "Understand implied meaning, not only literal words.", []string{"literature"},
			[]string{"Read between the lines—she's unhappy.", "The email was polite, but read between the lines."}),
		mkCard("en", "advanced", 4, "hit the nail on the head", "acertar em cheio", "expression", "Describe something exactly right.", []string{"feedback"},
			[]string{"You hit the nail on the head with that comment.", "Her analysis hit the nail on the head."}),
		mkCard("en", "advanced", 4, "pull strings", "puxar os fios / usar influência", "expression", "Use hidden influence to get an advantage.", []string{"work"},
			[]string{"He pulled strings to get the ticket.", "It's not what you know, it's who pulls strings."}),
		mkCard("en", "advanced", 4, "a blessing in disguise", "um mal que veio para bem", "expression", "Something bad that later brings good results.", []string{"life"},
			[]string{"Losing that job was a blessing in disguise.", "The delay was a blessing in disguise—we missed the storm."}),
		mkCard("en", "advanced", 5, "the tip of the iceberg", "só a ponta do iceberg", "expression", "A small visible part of a much larger problem.", []string{"news"},
			[]string{"These errors are just the tip of the iceberg.", "One complaint may be the tip of the iceberg."}),
		mkCard("en", "advanced", 5, "bite off more than you can chew", "querer abraçar o mundo", "expression", "Take on more responsibility than you can handle.", []string{"work"},
			[]string{"Don't bite off more than you can chew this semester.", "We bit off more than we could chew with the renovation."}),
	}
}

func esBeginner() []FlashcardDocument {
	return []FlashcardDocument{
		mkCard("es", "beginner", 1, "embarazada", "grávida (NÃO \"envergonhada\")", "adjective", "Means pregnant. For 'embarrassed' use avergonzada / tener vergüenza.", []string{"false-friend", "pt-br"},
			[]string{"Ella está embarazada de tres meses.", "No confundas embarazada con embarrassed."}),
		mkCard("es", "beginner", 1, "borracha", "bêbada (NÃO borracha de apagar)", "noun", "Often means 'drunk woman'. Rubber/eraser is goma or goma de borrar.", []string{"false-friend", "pt-br"},
			[]string{"No manejes si estás borracha.", "Necesito una goma de borrar, no 'una borracha'."}),
		mkCard("es", "beginner", 2, "polvo", "poeira; polvo marino = polvo do mar", "noun", "Dust; marine 'polvo' can mean sea sediment. Chicken is pollo, not polvo.", []string{"false-friend", "pt-br"},
			[]string{"Limpia el polvo de los muebles.", "Para pollo (frango), di pollo—no polvo."}),
		mkCard("es", "beginner", 2, "largo", "comprido / longo (NÃO 'grande')", "adjective", "Long in length. Big is grande.", []string{"false-friend", "pt-br"},
			[]string{"Tiene el pelo largo y oscuro.", "Una mesa larga, no 'larga' por 'grande'."}),
		mkCard("es", "beginner", 2, "en absoluto", "de jeito nenhum / de forma alguma", "expression", "Often means 'not at all', opposite of English 'absolutely' agreeing.", []string{"false-friend", "pt-br"},
			[]string{"¿Te molesta? —En absoluto.", "No me gusta en absoluto."}),
		mkCard("es", "beginner", 3, "constipado", "resfriado (NÃO constipado intestinal)", "adjective", "Having a cold. For constipation use estreñido.", []string{"false-friend", "pt-br"},
			[]string{"Estoy constipado y me duele la garganta.", "Para el estreñimiento, mejor decir estreñido."}),
		mkCard("es", "beginner", 3, "éxito", "sucesso (NÃO saída)", "noun", "Success. Exit/salida is salida.", []string{"false-friend", "pt-br"},
			[]string{"Tuvo mucho éxito con su libro.", "La salida está allí—no 'el éxito'."}),
		mkCard("es", "beginner", 3, "actual", "atual / presente", "adjective", "Current/present, not 'actual' as in English 'real'.", []string{"false-friend", "pt-br"},
			[]string{"El presidente actual hablará hoy.", "En realidad (really), no actualmente en ese sentido."}),
		mkCard("es", "beginner", 4, "carpeta", "pasta de papéis (NÃO carpete)", "noun", "Folder. Carpet is alfombra or moqueta.", []string{"false-friend", "pt-br"},
			[]string{"Guarda el contrato en la carpeta.", "La alfombra es suave—no carpeta."}),
		mkCard("es", "beginner", 4, "emocionado", "emocionado / animado (pode ser 'nervioso' leve)", "adjective", "Excited or moved; nuance differs from Portuguese 'emocionado' alone.", []string{"nuance", "pt-br"},
			[]string{"Estoy emocionado por el viaje.", "Estaba tan emocionado que temblaba."}),
	}
}

func esIntermediate() []FlashcardDocument {
	return []FlashcardDocument{
		mkCard("es", "intermediate", 2, "ropa", "roupa", "noun", "Clothing. Rope is cuerda or soga.", []string{"vocabulary"},
			[]string{"Colgué la ropa en el tendedero.", "Necesito comprar ropa de invierno."}),
		mkCard("es", "intermediate", 2, "sopa", "sopa", "noun", "Soup—not 'soap' (jabón).", []string{"false-friend", "pt-br"},
			[]string{"La sopa de lentejas está deliciosa.", "El jabón está en el baño, no en el plato."}),
		mkCard("es", "intermediate", 3, "grabar", "gravar (áudio/vídeo)", "verb", "To record. Engrave can be grabar too—context clarifies.", []string{"media"},
			[]string{"Voy a grabar la entrevista.", "Grabaron un disco en vivo."}),
		mkCard("es", "intermediate", 3, "asistir", "comparecer / estar presente", "verb", "To attend an event. To help is ayudar.", []string{"false-friend", "pt-br"},
			[]string{"Asistiré a la conferencia mañana.", "¿Puedes ayudarme?—no 'asistirme'."}),
		mkCard("es", "intermediate", 3, "contestar", "responder", "verb", "To answer. To contest/dispute is impugnar or disputar.", []string{"false-friend", "pt-br"},
			[]string{"Contesta el teléfono, por favor.", "Contestó con educación a la pregunta difícil."}),
		mkCard("es", "intermediate", 4, "carrera", "curso universitário / corrida", "noun", "Degree/career or a race—context.", []string{"education"},
			[]string{"Estudia la carrera de medicina.", "Ganó la carrera de los cien metros."}),
		mkCard("es", "intermediate", 4, "ésta / esta", "esta (demostrativo)", "determiner", "Accent on demonstratives is largely obsolete in modern print; esta ciudad.", []string{"spelling"},
			[]string{"Esta ciudad me encanta.", "Esto es importante para mí."}),
		mkCard("es", "intermediate", 4, "por favor", "por favor", "expression", "Please—place flexibly: Dame agua, por favor.", []string{"politeness"},
			[]string{"Pase, por favor.", "¿Podrías repetir, por favor?"}),
		mkCard("es", "intermediate", 5, "subir", "subir / aumentar", "verb", "To go up, upload, or increase.", []string{"daily-use"},
			[]string{"Subimos la montaña al amanecer.", "Subieron los precios del transporte."}),
		mkCard("es", "intermediate", 5, "vale", "ok / tá bom (ES peninsular)", "expression", "Informal agreement—very common in Spain.", []string{"conversation"},
			[]string{"—¿Vemos a las ocho? —Vale.", "Vale, lo dejamos así."}),
	}
}

func esAdvanced() []FlashcardDocument {
	return []FlashcardDocument{
		mkCard("es", "advanced", 2, "ser pan comido", "ser muito fácil", "expression", "Something very easy.", []string{"idiom"},
			[]string{"Para ella, el examen fue pan comido.", "Con práctica, esto será pan comido."}),
		mkCard("es", "advanced", 2, "echar una mano", "dar uma mão", "expression", "To help someone informally.", []string{"idiom"},
			[]string{"¿Te echo una mano con las cajas?", "Me echó una mano cuando más lo necesitaba."}),
		mkCard("es", "advanced", 3, "quedarse en nada", "não se concretizar", "expression", "When plans fail to materialize.", []string{"idiom"},
			[]string{"El proyecto se quedó en nada.", "La reunión quedó en nada tras el cambio de jefe."}),
		mkCard("es", "advanced", 3, "tener mano izquierda", "tertúlia fina / jeito social", "expression", "Skill in handling people tactfully.", []string{"work"},
			[]string{"Un buen líder tiene mano izquierda.", "Falta mano izquierda en ese feedback."}),
		mkCard("es", "advanced", 3, "no tener pelos en la lengua", "falar sem rodeios", "expression", "To speak bluntly.", []string{"personality"},
			[]string{"Ella no tiene pelos en la lengua.", "Me gusta la gente sin pelos en la lengua."}),
		mkCard("es", "advanced", 4, "más vale tarde que nunca", "antes tarde do que nunca", "expression", "Better late than never.", []string{"proverb"},
			[]string{"Llegó más vale tarde que nunca.", "Aprender idiomas: más vale tarde que nunca."}),
		mkCard("es", "advanced", 4, "a buen entendedor, pocas palabras bastan", "a palavra dada...", "expression", "A hint is enough for a smart listener.", []string{"proverb"},
			[]string{"A buen entendedor, pocas palabras.", "No diré más: a buen entendedor..."}),
		mkCard("es", "advanced", 4, "estar en el aire", "estar indefinido / em suspenso", "expression", "Uncertain, not decided yet.", []string{"work"},
			[]string{"El contrato sigue en el aire.", "Todo quedó en el aire tras la reunión."}),
		mkCard("es", "advanced", 5, "dar largas", "empurrar com a barriga / protelar", "expression", "To stall or put something off repeatedly.", []string{"idiom"},
			[]string{"Me dan largas con el reembolso.", "Deja de dar largas y decide."}),
		mkCard("es", "advanced", 5, "ponerse las pilas", "animar-se / focar", "expression", "To get serious and focus—like 'get your act together'.", []string{"informal"},
			[]string{"Ponte las pilas con el estudio.", "Hay que ponerse las pilas este mes."}),
	}
}
