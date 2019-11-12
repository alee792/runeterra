package datadragon

import proto "github.com/alee792/runeterra/proto"

// Auto-generated from Data Dragon en_us.
var (
	// Keywords
	KeywordObliterate = proto.Keyword{
		Description: "Completely removed from the game. Doesn't cause Last Breath and can't be revived.",
		ID: &proto.ID{
			Name:    "Obliterate",
			NameRef: "Obliterate",
		},
	}

	KeywordSkill = proto.Keyword{
		Description: "A spell-like effect created and cast by unit.",
		ID: &proto.ID{
			Name:    "Skill",
			NameRef: "Skill",
		},
	}

	KeywordDoubleStrike = proto.Keyword{
		Description: "While attacking, it strikes both before AND at the same time as its blocker.",
		ID: &proto.ID{
			Name:    "Double Attack",
			NameRef: "Double Strike",
		},
	}

	KeywordWeakest = proto.Keyword{
		Description: "Lowest Power, with ties broken by lowest Health then lowest Cost",
		ID: &proto.ID{
			Name:    "Weakest",
			NameRef: "Weakest",
		},
	}

	KeywordElusive = proto.Keyword{
		Description: "Can only be blocked by an Elusive unit.",
		ID: &proto.ID{
			Name:    "Elusive",
			NameRef: "Elusive",
		},
	}

	KeywordDrain = proto.Keyword{
		Description: "Heal our Nexus for the amount of damage dealt",
		ID: &proto.ID{
			Name:    "Drain",
			NameRef: "Drain",
		},
	}

	KeywordStun = proto.Keyword{
		Description: "Remove a unit from combat. It can't attack or block for the rest of the round.",
		ID: &proto.ID{
			Name:    "Stun",
			NameRef: "Stun",
		},
	}

	KeywordAutoplay = proto.Keyword{
		Description: "Attaches to another card, trapping it. When the trapped card is drawn, perform the effect.",
		ID: &proto.ID{
			Name:    "Trap",
			NameRef: "Autoplay",
		},
	}

	KeywordPiltoverZaun = proto.Keyword{
		Description: " ",
		ID: &proto.ID{
			Name:    "Piltover & Zaun",
			NameRef: "PiltoverZaun",
		},
	}

	KeywordDemacia = proto.Keyword{
		Description: " ",
		ID: &proto.ID{
			Name:    "Demacia",
			NameRef: "Demacia",
		},
	}

	KeywordShadowIsles = proto.Keyword{
		Description: " ",
		ID: &proto.ID{
			Name:    "Shadow Isles",
			NameRef: "ShadowIsles",
		},
	}

	KeywordSpellOverwhelm = proto.Keyword{
		Description: "Inflicts damage beyond what would kill the target(s) to the enemy Nexus.",
		ID: &proto.ID{
			Name:    "Overwhelm",
			NameRef: "SpellOverwhelm",
		},
	}

	KeywordBarrier = proto.Keyword{
		Description: "Negates the next damage the unit would take. Lasts one round.",
		ID: &proto.ID{
			Name:    "Barrier",
			NameRef: "Barrier",
		},
	}

	KeywordCapture = proto.Keyword{
		Description: "A Captured card is removed from the game. It returns when the Capturing unit leaves play.",
		ID: &proto.ID{
			Name:    "Capture",
			NameRef: "Capture",
		},
	}

	KeywordFrostbite = proto.Keyword{
		Description: "Set a unit's Power to 0 this round (it can be changed after).",
		ID: &proto.ID{
			Name:    "Frostbite",
			NameRef: "Frostbite",
		},
	}

	KeywordBurst = proto.Keyword{
		Description: "Burst spells resolve instantly. The enemy can't act before it finishes.",
		ID: &proto.ID{
			Name:    "Burst",
			NameRef: "Burst",
		},
	}

	KeywordFleeting = proto.Keyword{
		Description: "Fleeting cards discard from hand when the round ends.",
		ID: &proto.ID{
			Name:    "Fleeting",
			NameRef: "Fleeting",
		},
	}

	KeywordFast = proto.Keyword{
		Description: "Fast spells can be played at any time, but allow the opponent to respond.",
		ID: &proto.ID{
			Name:    "Fast",
			NameRef: "Fast",
		},
	}

	KeywordOverwhelm = proto.Keyword{
		Description: "Excess damage I deal to my blocker is dealt to the enemy Nexus.",
		ID: &proto.ID{
			Name:    "Overwhelm",
			NameRef: "Overwhelm",
		},
	}

	KeywordQuickStrike = proto.Keyword{
		Description: "While attacking, strikes before its blocker.",
		ID: &proto.ID{
			Name:    "Quick Attack",
			NameRef: "Quick Strike",
		},
	}

	KeywordTough = proto.Keyword{
		Description: "Takes 1 less damage from all sources.",
		ID: &proto.ID{
			Name:    "Tough",
			NameRef: "Tough",
		},
	}

	KeywordRecall = proto.Keyword{
		Description: "Return a unit to hand and remove all effects applied to it.",
		ID: &proto.ID{
			Name:    "Recall",
			NameRef: "Recall",
		},
	}

	KeywordIonia = proto.Keyword{
		Description: " ",
		ID: &proto.ID{
			Name:    "Ionia",
			NameRef: "Ionia",
		},
	}

	KeywordRegeneration = proto.Keyword{
		Description: "Heals fully at the start of each round.",
		ID: &proto.ID{
			Name:    "Regeneration",
			NameRef: "Regeneration",
		},
	}

	KeywordLifesteal = proto.Keyword{
		Description: "Damage this unit deals heals its Nexus that amount.",
		ID: &proto.ID{
			Name:    "Lifesteal",
			NameRef: "Lifesteal",
		},
	}

	KeywordEnlightened = proto.Keyword{
		Description: "You're Enlightened when you have 10 max mana.",
		ID: &proto.ID{
			Name:    "Enlightened",
			NameRef: "Enlightened",
		},
	}

	KeywordSlow = proto.Keyword{
		Description: "Slow spells can be cast outside of combat and other casting. The enemy can respond.",
		ID: &proto.ID{
			Name:    "Slow",
			NameRef: "Slow",
		},
	}

	KeywordNoxus = proto.Keyword{
		Description: " ",
		ID: &proto.ID{
			Name:    "Noxus",
			NameRef: "Noxus",
		},
	}

	KeywordEphemeral = proto.Keyword{
		Description: "This unit dies when it strikes or when the round ends.",
		ID: &proto.ID{
			Name:    "Ephemeral",
			NameRef: "Ephemeral",
		},
	}

	KeywordFreljord = proto.Keyword{
		Description: " ",
		ID: &proto.ID{
			Name:    "Freljord",
			NameRef: "Freljord",
		},
	}

	KeywordLastBreath = proto.Keyword{
		Description: "These abilities take effect when the unit dies.",
		ID: &proto.ID{
			Name:    "Last Breath",
			NameRef: "Last Breath",
		},
	}

	KeywordChallenger = proto.Keyword{
		Description: "Can choose which enemy unit blocks.",
		ID: &proto.ID{
			Name:    "Challenger",
			NameRef: "Challenger",
		},
	}

	KeywordImbue = proto.Keyword{
		Description: "These abilities trigger when you resolve a spell.",
		ID: &proto.ID{
			Name:    "Imbue",
			NameRef: "Imbue",
		},
	}

	KeywordFearsome = proto.Keyword{
		Description: "Can only be blocked by enemies with 3 or more Attack.",
		ID: &proto.ID{
			Name:    "Fearsome",
			NameRef: "Fearsome",
		},
	}

	KeywordCantBlock = proto.Keyword{
		Description: " ",
		ID: &proto.ID{
			Name:    "Can't Block",
			NameRef: "CantBlock",
		},
	}

	// Regions
	RegionNeutral = proto.Region{
		Abbreviation: "NE",
		ID: &proto.ID{
			Name:    "Neutral",
			NameRef: "Neutral",
		},
		IconAbsolutePath: "",
	}

	RegionNoxus = proto.Region{
		Abbreviation: "NX",
		ID: &proto.ID{
			Name:    "Noxus",
			NameRef: "Noxus",
		},
		IconAbsolutePath: "",
	}

	RegionDemacia = proto.Region{
		Abbreviation: "DE",
		ID: &proto.ID{
			Name:    "Demacia",
			NameRef: "Demacia",
		},
		IconAbsolutePath: "",
	}

	RegionFreljord = proto.Region{
		Abbreviation: "FR",
		ID: &proto.ID{
			Name:    "Freljord",
			NameRef: "Freljord",
		},
		IconAbsolutePath: "",
	}

	RegionShadowIsles = proto.Region{
		Abbreviation: "SI",
		ID: &proto.ID{
			Name:    "Shadow Isles",
			NameRef: "ShadowIsles",
		},
		IconAbsolutePath: "",
	}

	RegionIonia = proto.Region{
		Abbreviation: "IO",
		ID: &proto.ID{
			Name:    "Ionia",
			NameRef: "Ionia",
		},
		IconAbsolutePath: "",
	}

	RegionPiltoverZaun = proto.Region{
		Abbreviation: "PZ",
		ID: &proto.ID{
			Name:    "Piltover & Zaun",
			NameRef: "PiltoverZaun",
		},
		IconAbsolutePath: "",
	}

	// Spell Speeds
	SpellSpeedSlow = proto.SpellSpeed{
		ID: &proto.ID{
			Name:    "Slow",
			NameRef: "Slow",
		},
	}

	SpellSpeedBurst = proto.SpellSpeed{
		ID: &proto.ID{
			Name:    "Burst",
			NameRef: "Burst",
		},
	}

	SpellSpeedFast = proto.SpellSpeed{
		ID: &proto.ID{
			Name:    "Fast",
			NameRef: "Fast",
		},
	}

	// Rarities
	RarityCommon = proto.Rarity{
		ID: &proto.ID{
			Name:    "Common",
			NameRef: "Common",
		},
	}

	RarityRare = proto.Rarity{
		ID: &proto.ID{
			Name:    "Rare",
			NameRef: "Rare",
		},
	}

	RarityEpic = proto.Rarity{
		ID: &proto.ID{
			Name:    "Epic",
			NameRef: "Epic",
		},
	}

	RarityChampion = proto.Rarity{
		ID: &proto.ID{
			Name:    "Champion",
			NameRef: "Champion",
		},
	}

	RarityNone = proto.Rarity{
		ID: &proto.ID{
			Name:    "None",
			NameRef: "None",
		},
	}
)
