package tarot

type CardMeaning struct {
	Name        string
	Heading     string
	Description string
	Love        string
	Career      string
}

var CardMeanings = map[string]CardMeaning{
	"The Fool": {
		"The Fool",
		"🎒 The Fool – New Beginnings & Adventure",
		"The Fool symbolizes new beginnings, adventure, and the spirit of taking a leap of faith. It represents the start of a journey filled with potential and discovery, driven by curiosity and trust in the unknown. This card encourages you to embrace change with an open heart and a sense of wonder.",
		"In love, The Fool suggests a fresh start, new romance, or the willingness to explore new emotional connections. It represents spontaneity and the joy of discovery in relationships.",
		"In career, this card symbolizes stepping into new opportunities with optimism and courage. It’s a time to take risks, explore new ventures, and trust your instincts.",
	},

	"The Magician": {
		"The Magician",
		"✨ The Magician – Manifestation & Skill",
		"The Magician represents manifestation, mastery, and the power to bring ideas into reality. It symbolizes the harnessing of one’s skills, creativity, and intellect to achieve goals. This card encourages you to recognize your own potential and act with confidence and purpose.",
		"In love, The Magician suggests the power to attract and manifest the kind of relationship you desire. It symbolizes charm, communication, and the ability to connect deeply.",
		"In career, this card represents resourcefulness, ambition, and the ability to make things happen. It’s a time to take initiative and turn your ideas into action.",
	},

	"The High Priestess": {
		"The High Priestess",
		"🌙 The High Priestess – Intuition & Mystery",
		"The High Priestess symbolizes intuition, hidden knowledge, and the mysteries of the subconscious. It represents the ability to access inner wisdom and understand truths that are not immediately visible. This card encourages you to trust your instincts and seek answers within yourself.",
		"In love, The High Priestess suggests deep emotional understanding and the need for patience and reflection. It may indicate unspoken feelings or the need for quiet understanding in the relationship.",
		"In career, this card represents the need for careful planning, intuition, and strategic thinking. It’s a time to listen to your inner voice and observe before acting.",
	},

	"The Empress": {
		"The Empress",
		"🌺 The Empress – Abundance & Nurturing",
		"The Empress symbolizes nurturing, creativity, and the abundance of life. It represents growth, fertility, and the ability to create and nurture ideas, projects, and relationships. This card encourages you to embrace compassion, care, and the natural flow of abundance.",
		"In love, The Empress suggests deep emotional connection, nurturing, and the potential for growth in the relationship. It’s a sign of comfort, affection, and strong bonds.",
		"In career, this card represents creative projects, growth opportunities, and the successful nurturing of business ventures. It’s a time to cultivate your ideas and watch them flourish.",
	},

	"The Emperor": {
		"The Emperor",
		"🏛️ The Emperor – Authority & Structure",
		"The Emperor symbolizes authority, structure, and the power of strategic planning. It represents stability, order, and the capacity to lead with confidence and control. This card encourages you to assert your authority and take a disciplined approach to achieving your goals.",
		"In love, The Emperor suggests stability, commitment, and the desire for a solid foundation in the relationship. It represents loyalty, protection, and strong support.",
		"In career, this card represents leadership, organization, and the ability to execute plans effectively. It’s a time to take control, implement strategies, and establish order.",
	},

	"The Hierophant": {
		"The Hierophant",
		"📜 The Hierophant – Tradition & Spiritual Guidance",
		"The Hierophant represents tradition, spiritual wisdom, and the pursuit of knowledge through structured learning. It symbolizes the importance of community, shared beliefs, and the guidance of trusted mentors. This card encourages you to seek wisdom in conventional practices and to respect established principles.",
		"In love, The Hierophant suggests commitment, tradition, and stability in relationships. It can also indicate marriage or the deepening of a bond through shared beliefs.",
		"In career, this card represents structured learning, mentorship, and adherence to established norms. It encourages you to follow proven methods and seek guidance from experienced individuals.",
	},

	"The Lovers": {
		"The Lovers",
		"💞 The Lovers – Union & Choices",
		"The Lovers symbolizes harmony, union, and the powerful connection between two souls. It represents meaningful relationships, deep connections, and the choices that come with commitment. This card encourages you to follow your heart and make decisions that align with your values.",
		"In love, The Lovers signifies soul-level connections, mutual understanding, and the power of choice in matters of the heart. It may also represent a crossroads where a decision must be made regarding a relationship.",
		"In career, this card suggests partnerships built on trust and shared vision. It can also indicate important decisions that require alignment with your core values.",
	},

	"The Chariot": {
		"The Chariot",
		"🚀 The Chariot – Willpower & Determination",
		"The Chariot symbolizes victory through determination, focus, and disciplined effort. It represents the ability to overcome obstacles through sheer willpower and the pursuit of your goals. This card encourages you to take control of your path and drive forward with confidence.",
		"In love, The Chariot suggests progress, determination to overcome challenges, and the power of commitment. It indicates moving forward together with shared goals and purpose.",
		"In career, this card represents ambition, success through hard work, and the power to achieve your objectives. It encourages you to stay focused and push through obstacles with unwavering determination.",
	},

	"Strength": {
		"Strength",
		"🦁 Strength – Courage & Inner Power",
		"Strength represents inner resilience, patience, and the mastery of one’s emotions. It symbolizes the power of calm confidence and compassion over brute force. This card encourages you to face challenges with grace and maintain composure even in the face of adversity.",
		"In love, Strength suggests enduring love, patience, and the ability to handle conflicts with understanding. It symbolizes compassion and the strength to support one another through difficulties.",
		"In career, this card represents perseverance, leadership through empathy, and the ability to overcome obstacles. It encourages you to approach challenges with confidence and steady resolve.",
	},

	"The Hermit": {
		"The Hermit",
		"🕯️ The Hermit – Introspection & Solitude",
		"The Hermit symbolizes the need for introspection, self-discovery, and retreat from distractions to seek clarity and wisdom. It represents a journey inward to find personal truth and a deeper understanding of one's path. This card encourages reflection, meditation, and solitude as means of personal growth.",
		"In love, The Hermit may indicate a period of reflection or emotional distance to understand personal needs and desires. It symbolizes the importance of understanding oneself before fully committing to another.",
		"In career, this card suggests taking time to reflect on your path and goals. It may indicate a need for strategic planning and independent thought to achieve clarity.",
	},

	"Wheel of Fortune": {
		"Wheel of Fortune",
		"🔄 Wheel of Fortune – Change & Destiny",
		"The Wheel of Fortune represents cycles of change, fate, and the ever-turning nature of life. It symbolizes both unexpected opportunities and unforeseen challenges, reminding you that nothing is permanent. This card encourages acceptance of change and trust in the process of transformation.",
		"In love, the Wheel of Fortune signifies shifts in your relationship status, such as new beginnings or transformative changes. It may indicate that fate is playing a role in your connections.",
		"In career, this card suggests a turning point, sudden opportunities, or a shift in direction. It encourages adaptability and openness to new possibilities as circumstances evolve.",
	},

	"Justice": {
		"Justice",
		"⚖️ Justice – Truth & Fairness",
		"Justice symbolizes truth, fairness, and the consequences of your actions. It represents the importance of accountability and the need for balanced decision-making. This card encourages you to act with integrity and consider the impact of your choices.",
		"In love, Justice suggests the need for honesty, balance, and mutual respect in relationships. It may also indicate important decisions that require fairness and transparency.",
		"In career, this card represents fair treatment, legal matters, and ethical decision-making. It encourages you to approach work with integrity and clear judgment.",
	},

	"The Hanged Man": {
		"The Hanged Man",
		"🔄 The Hanged Man – Perspective & Surrender",
		"The Hanged Man symbolizes the need to pause, reflect, and view situations from a new perspective. It represents surrender, patience, and the acceptance of moments of stillness. This card encourages letting go of control and embracing a new way of seeing things.",
		"In love, The Hanged Man suggests a period of reflection or waiting, often to gain clarity or shift perspective. It symbolizes patience and the need to release expectations.",
		"In career, this card represents taking a step back to evaluate your path and gain new insights. It encourages you to look at challenges from different angles to find creative solutions.",
	},

	"Death": {
		"Death",
		"💀 Death – Transformation & Rebirth",
		"Death symbolizes major transformation, endings, and the clearing away of old patterns to make way for new growth. It represents the end of a cycle and the opportunity for renewal and change. This card encourages you to embrace transformation rather than fear it.",
		"In love, Death signifies change within a relationship, either through growth or necessary endings. It can indicate letting go of old habits to allow space for something new.",
		"In career, this card suggests the end of a job, project, or phase, clearing the path for new opportunities. It encourages embracing change as a means of growth and evolution.",
	},

	"Temperance": {
		"Temperance",
		"🌈 Temperance – Balance & Harmony",
		"Temperance represents balance, moderation, and the blending of opposites to achieve harmony. It symbolizes patience, adaptability, and the pursuit of equilibrium in all aspects of life. This card encourages you to find middle ground and approach situations with calm and composure.",
		"In love, Temperance suggests a balanced, harmonious relationship that grows through mutual understanding and compromise. It symbolizes patience and the blending of energies.",
		"In career, this card represents finding balance between work and life, managing projects with patience, and blending skills to achieve success. It encourages adaptability and thoughtful planning.",
	},

	"The Devil": {
		"The Devil",
		"🩸 The Devil – Bondage & Temptation",
		"The Devil represents materialism, temptation, and the feeling of being trapped by desires or habits. It symbolizes illusions of powerlessness and the need to confront your inner shadows. This card encourages you to acknowledge your restraints and take action to break free from harmful cycles.",
		"In love, The Devil suggests passionate but possibly toxic attachments, obsession, or emotional entanglement. It’s a reminder to recognize unhealthy patterns and seek balance.",
		"In career, this card indicates feeling stuck or controlled by material concerns, work addiction, or negative influences. It urges you to identify what is holding you back and make conscious efforts to regain control.",
	},

	"The Tower": {
		"The Tower",
		"⚡ The Tower – Sudden Upheaval & Revelation",
		"The Tower symbolizes sudden change, upheaval, and the breaking down of old structures. It represents a moment of truth that clears away illusions and forces you to confront reality. This card encourages you to embrace transformation and rebuild stronger foundations.",
		"In love, The Tower can indicate sudden revelations, breakups, or major shifts in the relationship. It represents the breaking down of illusions to reveal the truth.",
		"In career, this card suggests dramatic change, unexpected challenges, or the collapse of a project or role. It urges you to adapt quickly and rebuild with greater awareness.",
	},

	"The Star": {
		"The Star",
		"⭐ The Star – Hope & Inspiration",
		"The Star represents hope, healing, and inspiration after challenging times. It symbolizes clarity, renewal, and a sense of purpose that guides you forward. This card encourages you to trust in the universe and remain optimistic about the future.",
		"In love, The Star suggests healing, harmony, and renewed hope in relationships. It represents openness and faith in the connection’s growth and potential.",
		"In career, this card signifies inspiration, creative vision, and a clear path forward. It encourages you to pursue your dreams with confidence and clarity.",
	},

	"The Moon": {
		"The Moon",
		"🌙 The Moon – Illusion & Intuition",
		"The Moon represents the subconscious, illusion, and the unknown. It symbolizes intuition, mystery, and the need to navigate through uncertainty with inner guidance. This card encourages you to trust your instincts and explore hidden truths.",
		"In love, The Moon suggests emotional confusion, hidden feelings, or the need to uncover deeper truths in the relationship. It’s a call to address insecurities and communicate openly.",
		"In career, this card represents uncertainty, unclear paths, or hidden challenges. It encourages you to trust your intuition and seek clarity before making decisions.",
	},

	"The Sun": {
		"The Sun",
		"🌞 The Sun – Success & Vitality",
		"The Sun represents joy, success, and a surge of positive energy. It symbolizes clarity, optimism, and the celebration of accomplishments. This card encourages you to embrace positivity, express yourself freely, and bask in the warmth of your achievements.",
		"In love, The Sun suggests happiness, mutual understanding, and a strong connection filled with joy and optimism. It signifies harmony and the fulfillment of shared dreams.",
		"In career, this card represents success, recognition, and a period of growth and prosperity. It encourages you to pursue your goals confidently, as success is within reach.",
	},

	"Judgement": {
		"Judgement",
		"🔔 Judgement – Awakening & Renewal",
		"Judgement symbolizes reflection, awakening, and the call to evaluate your life choices. It represents the need for self-assessment, acceptance of past actions, and the readiness to move forward with clarity and purpose. This card encourages you to embrace transformation and growth, free from the weight of the past.",
		"In love, Judgement suggests a time of reflection and the potential for renewal or closure. It indicates making decisions that align with your true values and emotional needs.",
		"In career, this card represents the opportunity for growth through self-evaluation and the recognition of past efforts. It encourages stepping up, taking responsibility, and moving forward with a clearer vision.",
	},

	"The World": {
		"The World",
		"🌎 The World – Completion & Fulfillment",
		"The World symbolizes completion, fulfillment, and the successful achievement of long-term goals. It represents the closing of one chapter and the beginning of another, marked by a sense of unity and accomplishment. This card encourages you to celebrate your successes and prepare for new adventures.",
		"In love, The World signifies harmony, deep connection, and the feeling of completeness within a relationship. It may also represent long-term commitments or achieving important milestones together.",
		"In career, this card suggests the successful completion of major projects, reaching career milestones, and recognition for your hard work. It encourages you to embrace your achievements and look forward to new possibilities.",
	},

	// Cups
	"Ace of Cups": {
		"Ace of Cups",
		"🌊 Ace of Cups – Overflowing Emotions & New Beginnings",
		"The Ace of Cups represents the purest essence of emotional energy, symbolizing new beginnings, deep connections, and the flow of love and compassion. It’s a card of emotional awakening, inviting you to embrace your feelings and allow your heart to open fully. Like a cup overflowing with water, this card signals an abundance of joy, creativity, and spiritual fulfillment. It often suggests that something beautiful is blossoming, leading to deep personal growth and emotional clarity.",
		"The Ace of Cups signals the start of a new romantic connection or the deepening of an existing one, marked by vulnerability and openness. It encourages you to express your feelings freely and embrace the flow of love that surrounds you.",
		"This card suggests the beginning of a passion-driven project or a workplace environment that nurtures emotional well-being and creativity. It may also indicate strong support and cooperation from colleagues, fostering a sense of unity and purpose.",
	},

	"Two of Cups": {
		"Two of Cups",
		"🤝 Two of Cups – Unity & Partnership",
		"The Two of Cups signifies harmonious partnerships and mutual understanding. It symbolizes the balance and flow of emotions between two individuals, whether in romance, friendship, or collaboration. This card highlights trust, equality, and the willingness to share openly, reflecting a deep emotional connection.",
		"In love, the Two of Cups indicates a powerful bond and mutual attraction, where both partners are equally invested and supportive. It’s a card of unity and shared joy, suggesting the beginning of a new relationship or the strengthening of an existing one.",
		"In career, it represents successful collaborations built on trust and understanding. It can also signify the start of a promising business partnership where both parties share the same vision and goals.",
	},

	"Three of Cups": {
		"Three of Cups",
		"🎉 Three of Cups – Celebration & Friendship",
		"The Three of Cups symbolizes celebration, friendship, and the joy of community. It’s a card of coming together with loved ones to share in good news, accomplishments, or simply the warmth of connection. This card invites you to cherish moments of happiness with those around you and appreciate the support network you have.",
		"In love, the Three of Cups suggests joyous occasions like engagements, anniversaries, or even reconciliations. It can also signify a sense of deep friendship and harmony within the relationship.",
		"In career, this card represents teamwork, group achievements, and collective success. It’s a time to celebrate milestones and recognize the contributions of others in your projects.",
	},

	"Four of Cups": {
		"Four of Cups",
		"🌑 Four of Cups – Contemplation & Apathy",
		"The Four of Cups represents introspection, emotional withdrawal, and a sense of dissatisfaction with what is currently offered. It symbolizes the need to step back and evaluate your feelings, recognizing moments of stagnation or disinterest. This card often suggests that opportunities are being overlooked due to a focus on what is missing rather than what is present.",
		"In love, the Four of Cups may indicate emotional distance or a feeling of being disconnected. It’s a call to reflect on your needs and desires, and whether they are being met in the relationship.",
		"In career, this card can suggest dissatisfaction with your current role or a lack of inspiration. It may be time to reassess your ambitions and consider what would truly bring fulfillment.",
	},

	"Five of Cups": {
		"Five of Cups",
		"💧 Five of Cups – Loss & Regret",
		"The Five of Cups symbolizes sorrow, regret, and focusing on what has been lost rather than what remains. It’s a card of emotional setback, but also a gentle reminder that not everything is gone; there is still hope and opportunity if you shift your perspective. This card invites you to acknowledge your grief but also to look for the hidden blessings still present.",
		"In love, the Five of Cups may represent heartbreak, separation, or disappointment. It encourages healing and finding closure while recognizing the potential for new emotional beginnings.",
		"In career, this card can indicate setbacks or the loss of an opportunity. However, it also urges you to learn from the experience and remain open to new paths that may emerge.",
	},

	"Six of Cups": {
		"Six of Cups",
		"🌼 Six of Cups – Nostalgia & Innocence",
		"The Six of Cups represents nostalgia, childhood memories, and the simple joys of life. It’s a card that brings a sense of innocence and warmth, inviting you to reconnect with your past or appreciate moments of kindness and generosity. This card often symbolizes a return to familiar places or the rekindling of old relationships.",
		"In love, the Six of Cups suggests a connection rooted in familiarity and comfort. It may also signify reunions or healing of past emotional wounds.",
		"In career, this card can represent a return to a past role, project, or collaboration that brought fulfillment. It also symbolizes a supportive environment where kindness and cooperation flourish.",
	},

	"Seven of Cups": {
		"Seven of Cups",
		"🌠 Seven of Cups – Choices & Illusions",
		"The Seven of Cups represents a time of many options and opportunities, often leading to feelings of confusion or indecision. It symbolizes dreams, fantasies, and the allure of possibilities that may not be grounded in reality. This card invites you to look closely at your options and consider which are truly attainable versus mere illusions.",
		"In love, the Seven of Cups can indicate idealizing a partner or being overwhelmed by choices in romantic prospects. It’s a reminder to ground your emotions and seek clarity in what you truly desire.",
		"In career, this card suggests multiple paths or projects, but warns of spreading yourself too thin. It’s important to prioritize and focus on what aligns with your long-term goals rather than short-lived fantasies.",
	},

	"Eight of Cups": {
		"Eight of Cups",
		"🌌 Eight of Cups – Moving On & Seeking Truth",
		"The Eight of Cups represents the decision to walk away from something unfulfilling in search of deeper meaning and emotional growth. It symbolizes leaving behind what no longer serves you, even if it’s difficult or uncertain. This card suggests the pursuit of authenticity and personal truth over superficial comfort.",
		"In love, the Eight of Cups often signifies the need to leave behind a stagnant or unfulfilling relationship. It can also mean moving on from emotional baggage to make space for healthier connections.",
		"In career, this card suggests reevaluating your path if it no longer brings you joy or purpose. It encourages you to seek work that resonates with your values and aspirations, even if it means stepping into the unknown.",
	},

	"Nine of Cups": {
		"Nine of Cups",
		"🍷 Nine of Cups – Contentment & Satisfaction",
		"The Nine of Cups is often called the 'wish card,' symbolizing emotional fulfillment, satisfaction, and the achievement of personal desires. It reflects a sense of accomplishment and contentment with where you are emotionally and materially. This card suggests that your efforts are coming to fruition, bringing joy and comfort into your life.",
		"In love, the Nine of Cups represents harmony, happiness, and mutual satisfaction in a relationship. It indicates that emotional needs are being met, and there is a sense of fulfillment with your partner.",
		"In career, this card suggests success and recognition for your hard work. It’s a sign that your goals are being realized, and you can take pride in your accomplishments.",
	},

	"Ten of Cups": {
		"Ten of Cups",
		"🏡 Ten of Cups – Harmony & Fulfillment",
		"The Ten of Cups represents emotional fulfillment, harmony, and lasting happiness within family and community. It symbolizes the ideal of peace and joy, where relationships are strong, supportive, and deeply satisfying. This card often signifies reaching a place of complete emotional contentment and unity.",
		"In love, the Ten of Cups suggests long-term commitment, emotional harmony, and a sense of completion with your partner. It represents shared dreams, stability, and mutual support.",
		"In career, this card signifies working in a harmonious environment where relationships are strong and supportive. It may also represent achieving career goals that bring a sense of pride and community involvement.",
	},

	"Page of Cups": {
		"Page of Cups",
		"🐟 Page of Cups – Imagination & New Emotional Beginnings",
		"The Page of Cups symbolizes the beginning of emotional exploration, creativity, and new possibilities. It’s a card of intuition, playfulness, and the willingness to dream big without limitations. This card often signals unexpected messages, emotional discoveries, or the blossoming of creative ideas.",
		"In love, the Page of Cups represents the spark of new feelings or a message of affection. It may indicate a budding romance or a renewal of emotional excitement in a current relationship.",
		"In career, this card suggests new creative projects, innovative ideas, and the exploration of untapped talents. It encourages approaching your work with curiosity and an open heart.",
	},

	"Knight of Cups": {
		"Knight of Cups",
		"🏇 Knight of Cups – Romance & Idealism",
		"The Knight of Cups represents the pursuit of dreams, romance, and emotional exploration. He is the messenger of love and inspiration, driven by the heart and guided by idealism. This card suggests taking action on emotional desires and moving forward with creative visions.",
		"In love, the Knight of Cups signifies romantic gestures, deep emotional connections, and the willingness to express feelings openly. It often heralds the arrival of someone passionate and emotionally expressive.",
		"In career, this card encourages following your heart and pursuing projects that inspire you. It suggests that creativity and vision will be key drivers of your success.",
	},

	"Queen of Cups": {
		"Queen of Cups",
		"👑 Queen of Cups – Compassion & Intuition",
		"The Queen of Cups symbolizes emotional maturity, deep intuition, and compassionate understanding. She represents the ability to navigate emotions with grace and empathy, offering support and nurturing to those around her. This card invites you to trust your instincts and embrace your emotional depth with kindness and sensitivity.",
		"In love, the Queen of Cups indicates deep emotional connection, empathy, and nurturing within the relationship. It suggests a partner who is caring, intuitive, and emotionally supportive.",
		"In career, this card represents compassionate leadership, emotional intelligence, and the ability to connect with others on a meaningful level. It may also suggest a career path that involves helping or healing others.",
	},

	"King of Cups": {
		"King of Cups",
		"⚓ King of Cups – Emotional Balance & Wisdom",
		"The King of Cups symbolizes emotional mastery, wisdom, and control over one's feelings. He represents the ability to remain calm and compassionate even in turbulent situations. This card suggests a deep understanding of emotions and the capacity to offer guidance and support with strength and sensitivity.",
		"In love, the King of Cups signifies emotional stability, maturity, and commitment. It represents a partner who is caring, balanced, and deeply empathetic.",
		"In career, this card suggests strong leadership through emotional intelligence and compassion. It reflects the ability to handle challenges with grace and maintain harmony in the workplace.",
	},

	// Wands
	"Ace of Wands": {
		"Ace of Wands",
		"🔥 Ace of Wands – Inspiration & New Beginnings",
		"The Ace of Wands represents the spark of inspiration, bold ideas, and the beginning of new ventures. It symbolizes raw potential and the urge to bring creative visions to life. This card encourages you to take action and trust in your instincts as you move forward with passion and enthusiasm.",
		"In love, the Ace of Wands signifies fresh energy, passionate connections, and the excitement of new possibilities. It suggests that a relationship may be entering a phase of growth and exploration.",
		"In career, this card indicates the start of a new project, business venture, or a surge of creative energy. It’s a time to pursue your ambitions with confidence and drive.",
	},

	"Two of Wands": {
		"Two of Wands",
		"🌐 Two of Wands – Planning & Decision-Making",
		"The Two of Wands represents planning, foresight, and the initial stages of manifesting your goals. It symbolizes the need to look beyond your current situation and consider the broader possibilities that lie ahead. This card encourages strategic thinking and taking the first steps toward long-term vision.",
		"In love, the Two of Wands suggests a period of reflection and decision-making about the future of the relationship. It may indicate the need to consider long-term compatibility and shared goals.",
		"In career, this card emphasizes planning, strategy, and setting long-term objectives. It’s a time to expand your horizons and explore new possibilities for growth.",
	},

	"Three of Wands": {
		"Three of Wands",
		"🌄 Three of Wands – Expansion & Progress",
		"The Three of Wands symbolizes expansion, progress, and the manifestation of plans. It represents looking forward with optimism and seeing the results of your efforts begin to take shape. This card encourages confidence in your path and anticipation for what is yet to come.",
		"In love, the Three of Wands indicates growth and exploration within the relationship. It suggests that mutual efforts are leading to deeper connection and shared experiences.",
		"In career, this card represents progress in projects and the expansion of business ventures. It’s a time to look ahead confidently and capitalize on the momentum you’ve built.",
	},

	"Four of Wands": {
		"Four of Wands",
		"🏡 Four of Wands – Celebration & Stability",
		"The Four of Wands symbolizes celebration, harmony, and the achievement of milestones. It represents a time of joy, unity, and a strong foundation built through collaboration and support. This card often signifies family gatherings, weddings, or the completion of an important project.",
		"In love, the Four of Wands suggests commitment, harmony, and celebration of love. It may indicate moving in together, engagements, or simply feeling deeply connected with your partner.",
		"In career, this card represents successful completion of projects and recognition for your efforts. It suggests a supportive environment where teamwork and unity are celebrated.",
	},

	"Five of Wands": {
		"Five of Wands",
		"⚔️ Five of Wands – Conflict & Competition",
		"The Five of Wands represents conflict, competition, and differing opinions. It symbolizes the struggle to be heard, clashing perspectives, or the need to assert oneself in a competitive environment. This card encourages resilience and finding ways to navigate disagreements constructively.",
		"In love, the Five of Wands may indicate misunderstandings or conflicts arising from differing expectations. It’s a reminder to communicate openly and work through challenges together.",
		"In career, this card suggests competition, workplace disagreements, or the need to stand out among peers. It’s a call to harness your strengths and prove your value through perseverance.",
	},

	"Six of Wands": {
		"Six of Wands",
		"🏆 Six of Wands – Victory & Recognition",
		"The Six of Wands symbolizes success, recognition, and public acknowledgment of your achievements. It represents triumph after hard work and the admiration of those around you. This card encourages you to celebrate your victories and take pride in your accomplishments.",
		"In love, the Six of Wands indicates a relationship that is flourishing and recognized for its strength and harmony. It suggests mutual pride and support for each other’s achievements.",
		"In career, this card represents promotions, awards, or public recognition for your contributions. It’s a sign that your efforts are being seen and appreciated by others.",
	},

	"Seven of Wands": {
		"Seven of Wands",
		"🛡️ Seven of Wands – Defense & Perseverance",
		"The Seven of Wands represents standing your ground, defending your position, and pushing back against opposition. It symbolizes the need to be assertive and protect what you’ve worked for, despite challenges or competition. This card encourages you to remain resilient and confident in your abilities.",
		"In love, the Seven of Wands suggests standing up for your relationship or defending it against outside influences. It indicates the need for loyalty and commitment in the face of challenges.",
		"In career, this card represents facing competition or obstacles that require determination and perseverance. It’s a call to protect your achievements and keep moving forward with confidence.",
	},

	"Eight of Wands": {
		"Eight of Wands",
		"💨 Eight of Wands – Swift Action & Momentum",
		"The Eight of Wands symbolizes rapid movement, swift decisions, and the acceleration of events. It represents the momentum building towards a goal and the quick unfolding of plans. This card encourages you to act decisively and embrace the flow of energy pushing you forward.",
		"In love, the Eight of Wands indicates quick developments, passionate communication, and moving forward with enthusiasm. It may signify the rapid deepening of a connection or sudden clarity in emotions.",
		"In career, this card suggests progress, quick decision-making, and the fast advancement of projects. It’s a time to capitalize on momentum and push forward with confidence.",
	},

	"Nine of Wands": {
		"Nine of Wands",
		"🛡️ Nine of Wands – Resilience & Endurance",
		"The Nine of Wands represents resilience, determination, and the final push toward a goal. It symbolizes perseverance in the face of obstacles and the ability to stand strong despite challenges. This card encourages you to hold your ground and stay focused, even if the journey has been difficult.",
		"In love, the Nine of Wands suggests protecting your relationship from outside pressures and being willing to fight for it. It represents loyalty and commitment through difficult times.",
		"In career, this card indicates overcoming challenges, maintaining focus, and pushing through obstacles. It’s a reminder to stay strong and not give up when facing adversity.",
	},

	"Ten of Wands": {
		"Ten of Wands",
		"💼 Ten of Wands – Burden & Responsibility",
		"The Ten of Wands symbolizes carrying heavy responsibilities, burdens, and the weight of obligations. It represents hard work, perseverance, and the feeling of being stretched thin. This card encourages you to assess your load and seek ways to lighten it if necessary.",
		"In love, the Ten of Wands can suggest feeling overwhelmed or burdened by expectations within the relationship. It’s a call to communicate openly and find balance together.",
		"In career, this card represents heavy workloads and the pressure of responsibilities. It may be time to delegate or seek support to avoid burnout.",
	},

	"Page of Wands": {
		"Page of Wands",
		"📣 Page of Wands – Enthusiasm & Exploration",
		"The Page of Wands symbolizes youthful energy, enthusiasm, and the desire to explore new possibilities. It represents a burst of inspiration and the willingness to step into the unknown with optimism. This card encourages embracing new adventures with curiosity and confidence.",
		"In love, the Page of Wands suggests the spark of excitement and a desire to explore new experiences with your partner. It may indicate the beginning of a passionate connection.",
		"In career, this card represents new projects, learning opportunities, and the eagerness to take on fresh challenges. It’s a time to embrace change and pursue creative ideas.",
	},

	"Knight of Wands": {
		"Knight of Wands",
		"🏇 Knight of Wands – Action & Boldness",
		"The Knight of Wands symbolizes bold action, passion, and the drive to pursue goals fearlessly. It represents movement, adventure, and the willingness to take risks. This card encourages you to charge forward with confidence and embrace new opportunities.",
		"In love, the Knight of Wands signifies passionate connections, adventurous spirits, and bold expressions of affection. It suggests moving forward with confidence and excitement.",
		"In career, this card represents ambition, fast-paced projects, and the pursuit of big goals. It’s a time to act decisively and push your ideas into reality.",
	},

	"Queen of Wands": {
		"Queen of Wands",
		"👑 Queen of Wands – Confidence & Independence",
		"The Queen of Wands symbolizes confidence, independence, and magnetic energy. She represents the embodiment of passion, creativity, and the ability to inspire others. This card encourages you to embrace your power and express your individuality boldly.",
		"In love, the Queen of Wands suggests a dynamic and passionate relationship, full of energy and mutual inspiration. It indicates confidence and assertiveness in expressing feelings.",
		"In career, this card represents leadership, creativity, and the ability to command respect. It’s a sign to step into your power and pursue your goals fearlessly.",
	},

	"King of Wands": {
		"King of Wands",
		"🔥 King of Wands – Leadership & Vision",
		"The King of Wands symbolizes visionary leadership, authority, and the ability to inspire others. He represents the mastery of ambition, creativity, and the drive to achieve big goals. This card encourages bold decision-making and leading with confidence.",
		"In love, the King of Wands signifies a passionate and committed partner who is driven and inspiring. It suggests strong leadership and confidence in the relationship.",
		"In career, this card represents entrepreneurial spirit, strong leadership, and the ability to turn ideas into action. It’s a time to take charge and pursue your ambitions fearlessly.",
	},

	// Swords

	"Ace of Swords": {
		"Ace of Swords",
		"⚔️ Ace of Swords – Clarity & Breakthrough",
		"The Ace of Swords represents mental clarity, breakthroughs, and the cutting away of confusion. It symbolizes the power of clear thinking, truth, and decisive action. This card encourages you to face challenges head-on with logic and strength, embracing new ideas and opportunities for growth.",
		"In love, the Ace of Swords suggests honest communication and the clearing away of misunderstandings. It may indicate a breakthrough in understanding between partners.",
		"In career, this card represents fresh ideas, sharp focus, and decisive action. It’s a time to cut through obstacles and move forward with clear objectives.",
	},

	"Two of Swords": {
		"Two of Swords",
		"⚖️ Two of Swords – Indecision & Stalemate",
		"The Two of Swords symbolizes a state of indecision, mental conflict, and a stalemate in making a choice. It represents the need to make a difficult decision while feeling uncertain or conflicted. This card encourages you to look within for clarity and to face the truth, even if it’s uncomfortable.",
		"In love, the Two of Swords may indicate a crossroads or a difficult decision that must be made. It suggests the need for open communication to break the deadlock.",
		"In career, this card represents feeling stuck between choices or struggling to make a clear decision. It’s a reminder to weigh your options carefully and consider all perspectives.",
	},

	"Three of Swords": {
		"Three of Swords",
		"💔 Three of Swords – Heartbreak & Sorrow",
		"The Three of Swords represents heartbreak, sorrow, and emotional pain. It symbolizes the sharp sting of loss, betrayal, or deep sadness. This card encourages you to acknowledge your pain, heal, and eventually find the strength to move forward.",
		"In love, the Three of Swords suggests emotional pain, conflict, or separation. It’s a call to process feelings of loss and seek healing.",
		"In career, this card indicates setbacks, disappointment, or conflicts that may feel deeply impactful. It’s important to allow space for healing and reflection during challenging times.",
	},

	"Four of Swords": {
		"Four of Swords",
		"🛌 Four of Swords – Rest & Recovery",
		"The Four of Swords represents rest, retreat, and the need for recovery after a period of struggle or conflict. It symbolizes taking a step back to restore your energy and reflect before moving forward. This card encourages you to find peace, solitude, and mental clarity through stillness.",
		"In love, the Four of Swords suggests a period of reflection or taking space to heal and understand each other's needs. It may indicate the need for patience and understanding.",
		"In career, this card represents a pause or a break to regroup and strategize. It’s a time to recharge and prepare for future challenges.",
	},

	"Five of Swords": {
		"Five of Swords",
		"⚔️ Five of Swords – Conflict & Tension",
		"The Five of Swords represents conflict, tension, and the potential fallout of aggressive behavior. It symbolizes the desire to win at all costs, even if it means sacrificing harmony and understanding. This card encourages you to reflect on your actions and consider the true cost of victory.",
		"In love, the Five of Swords suggests arguments, misunderstandings, or power struggles. It’s a reminder to communicate with respect and seek resolution instead of winning at any cost.",
		"In career, this card indicates workplace conflicts or power struggles. It’s important to consider long-term relationships and cooperation over short-term wins.",
	},

	"Six of Swords": {
		"Six of Swords",
		"🌊 Six of Swords – Transition & Moving Forward",
		"The Six of Swords symbolizes transition, moving away from conflict, and heading toward calmer waters. It represents leaving behind difficulties and seeking peace, often through reflection and understanding. This card encourages you to embrace change and move forward with clarity and purpose.",
		"In love, the Six of Swords suggests healing and moving on from past conflicts. It may indicate traveling together or starting fresh in a new chapter of the relationship.",
		"In career, this card represents moving away from stress or conflict at work, potentially through a job change or a shift in perspective. It signifies progress towards stability and peace.",
	},

	"Seven of Swords": {
		"Seven of Swords",
		"🗝️ Seven of Swords – Strategy & Deception",
		"The Seven of Swords represents strategy, stealth, and sometimes deception. It symbolizes the need to act carefully, avoid conflict, or manage situations with subtlety. This card encourages you to be mindful of your actions and the intentions of those around you.",
		"In love, the Seven of Swords suggests the need for honesty and transparency. It may indicate hidden feelings or a lack of openness in the relationship.",
		"In career, this card represents strategic thinking, but also warns against dishonesty or cutting corners. It’s important to maintain integrity and think carefully before taking action.",
	},

	"Eight of Swords": {
		"Eight of Swords",
		"🔗 Eight of Swords – Restriction & Self-Imposed Limitations",
		"The Eight of Swords symbolizes feeling trapped, restricted, or held back by circumstances, often through one's own mindset. It represents a sense of being stuck or unable to see a way out of a challenging situation. This card encourages you to recognize your own power to change your perspective and break free from mental barriers.",
		"In love, the Eight of Swords suggests feeling confined or trapped within a relationship. It may indicate fears or insecurities that are holding you back from fully expressing yourself.",
		"In career, this card represents feeling restricted or limited by circumstances, such as a demanding role or lack of opportunities. It’s a call to examine your situation and consider new approaches to overcome obstacles.",
	},

	"Nine of Swords": {
		"Nine of Swords",
		"🌌 Nine of Swords – Anxiety & Worry",
		"The Nine of Swords symbolizes anxiety, worry, and sleepless nights caused by overthinking or emotional distress. It represents the weight of fears and negative thoughts that may feel overwhelming. This card encourages you to confront these feelings and seek understanding and clarity to move forward.",
		"In love, the Nine of Swords suggests fear, insecurity, or worry about the relationship's future. It’s a call to address concerns openly and seek reassurance where needed.",
		"In career, this card indicates stress, overwork, or anxiety about your professional path. It’s important to identify the root causes and find ways to manage your mental well-being.",
	},

	"Ten of Swords": {
		"Ten of Swords",
		"⚰️ Ten of Swords – Endings & Rock Bottom",
		"The Ten of Swords symbolizes painful endings, betrayal, or hitting rock bottom. It represents the final stage of a difficult situation, where there is no choice but to accept the outcome and move forward. This card encourages you to acknowledge the end of a cycle and embrace the opportunity for renewal.",
		"In love, the Ten of Swords suggests the painful end of a relationship or a deep emotional wound. It’s a time to seek closure and begin the process of healing.",
		"In career, this card indicates the end of a job, project, or professional relationship. Though challenging, it marks the beginning of new opportunities and growth.",
	},

	"Page of Swords": {
		"Page of Swords",
		"📜 Page of Swords – Curiosity & New Ideas",
		"The Page of Swords symbolizes curiosity, mental agility, and the thirst for knowledge. It represents the beginning of intellectual exploration and the eagerness to learn new things. This card encourages you to ask questions, seek understanding, and embrace new ideas with enthusiasm.",
		"In love, the Page of Swords suggests open communication, curiosity, and the desire to understand your partner better. It may indicate the need for honest conversations and clear boundaries.",
		"In career, this card represents new projects, learning opportunities, and the beginning of intellectual pursuits. It’s a time to embrace new challenges and gather information with sharp focus.",
	},

	"Knight of Swords": {
		"Knight of Swords",
		"⚔️ Knight of Swords – Action & Determination",
		"The Knight of Swords symbolizes bold action, determination, and the relentless pursuit of goals. It represents moving forward with intensity and clarity, driven by strong principles and convictions. This card encourages you to act decisively and confidently, cutting through obstacles without hesitation.",
		"In love, the Knight of Swords suggests a direct and honest approach to relationships. It may indicate rapid developments and clear communication of intentions.",
		"In career, this card represents ambition, drive, and the willingness to overcome obstacles. It’s a call to pursue your goals with focus and determination.",
	},

	"Queen of Swords": {
		"Queen of Swords",
		"👑 Queen of Swords – Clarity & Independence",
		"The Queen of Swords symbolizes sharp intellect, clear communication, and emotional independence. She represents the ability to see through illusions and approach situations with honesty and logic. This card encourages you to be clear in your intentions and articulate your thoughts with precision.",
		"In love, the Queen of Swords suggests honesty, clarity, and the setting of healthy boundaries. It may indicate the need for open communication and mutual respect.",
		"In career, this card represents strong decision-making skills, strategic thinking, and effective communication. It’s a time to approach challenges with logic and confidence.",
	},

	"King of Swords": {
		"King of Swords",
		"⚔️ King of Swords – Authority & Truth",
		"The King of Swords symbolizes authority, truth, and intellectual mastery. He represents clear thinking, strong leadership, and the ability to make fair and logical decisions. This card encourages you to approach situations with integrity and assert your truth confidently.",
		"In love, the King of Swords suggests a rational and clear-minded approach to relationships. It may indicate a partner who is intellectually stimulating and values honesty.",
		"In career, this card represents leadership, strategic planning, and the ability to solve problems effectively. It’s a time to take charge with logic and clear thinking.",
	},

	"Ace of Pentacles": {
		"Ace of Pentacles",
		"🌿 Ace of Pentacles – Prosperity & New Opportunities",
		"The Ace of Pentacles symbolizes new beginnings in the realm of material wealth, financial growth, and tangible success. It represents the planting of seeds for future prosperity, inviting you to manifest abundance through effort and planning. This card encourages you to seize opportunities that can lead to long-term security and success.",
		"In love, the Ace of Pentacles suggests stability, commitment, and the potential for growth in the relationship. It’s a sign of building strong foundations and nurturing mutual goals.",
		"In career, this card represents new job opportunities, financial growth, and the potential for long-term success. It’s a time to invest in your future and build a solid path forward.",
	},

	"Two of Pentacles": {
		"Two of Pentacles",
		"⚖️ Two of Pentacles – Balance & Adaptability",
		"The Two of Pentacles symbolizes the need for balance, adaptability, and managing multiple responsibilities. It represents juggling priorities, financial decisions, and finding harmony amidst the chaos of everyday life. This card encourages you to stay flexible and maintain your equilibrium as you handle various demands.",
		"In love, the Two of Pentacles suggests finding balance between relationship responsibilities and personal obligations. It may indicate the need for better time management to nurture the connection.",
		"In career, this card represents multitasking, balancing projects, and managing financial decisions. It’s important to stay organized and adaptable as you navigate changes.",
	},

	"Three of Pentacles": {
		"Three of Pentacles",
		"🏗️ Three of Pentacles – Collaboration & Mastery",
		"The Three of Pentacles represents teamwork, collaboration, and the mastery of skills through effort and cooperation. It symbolizes the coming together of talents to create something meaningful and well-crafted. This card encourages you to value the contributions of others and work toward common goals.",
		"In love, the Three of Pentacles suggests building a strong foundation through cooperation and shared goals. It indicates mutual effort and commitment to growing together.",
		"In career, this card represents teamwork, collaboration, and recognition for your skills. It’s a time to work effectively with others to achieve common objectives.",
	},

	"Four of Pentacles": {
		"Four of Pentacles",
		"💰 Four of Pentacles – Security & Control",
		"The Four of Pentacles symbolizes stability, security, and the desire to hold onto material possessions. It represents the need to protect your assets and establish a firm foundation. However, it can also indicate holding on too tightly, leading to stagnation or fear of loss.",
		"In love, the Four of Pentacles suggests possessiveness or the need to feel secure in the relationship. It’s a reminder to find balance between protection and trust.",
		"In career, this card represents financial stability and the careful management of resources. It may also indicate a reluctance to take risks, favoring security over growth.",
	},

	"Five of Pentacles": {
		"Five of Pentacles",
		"🏚️ Five of Pentacles – Hardship & Financial Struggle",
		"The Five of Pentacles symbolizes financial hardship, loss, or feelings of isolation. It represents times of difficulty, where resources are scarce, and support may seem distant. This card encourages you to seek help where it is available and remember that challenges are temporary.",
		"In love, the Five of Pentacles suggests struggles within the relationship, possibly linked to financial stress or emotional distance. It’s a reminder to support each other through tough times.",
		"In career, this card indicates financial setbacks, job insecurity, or challenges in the workplace. It’s important to stay resilient and seek support to overcome obstacles.",
	},

	"Six of Pentacles": {
		"Six of Pentacles",
		"⚖️ Six of Pentacles – Generosity & Giving",
		"The Six of Pentacles symbolizes charity, generosity, and the balance of giving and receiving. It represents the flow of resources and the importance of supporting others when you are in a position to do so. This card encourages you to share your abundance and trust that generosity will come back to you.",
		"In love, the Six of Pentacles suggests mutual support, balance, and generosity between partners. It indicates a healthy flow of giving and receiving in the relationship.",
		"In career, this card represents financial help, collaborative support, or the sharing of resources. It may indicate assistance coming your way or the opportunity to help others succeed.",
	},

	"Seven of Pentacles": {
		"Seven of Pentacles",
		"🌱 Seven of Pentacles – Patience & Long-Term Vision",
		"The Seven of Pentacles symbolizes patience, reflection, and the evaluation of progress. It represents the need to assess your investments—whether financial, personal, or emotional—to determine if they are yielding the results you desire. This card encourages you to take a step back, reflect on your journey, and make adjustments if necessary.",
		"In love, the Seven of Pentacles suggests evaluating the relationship's progress and understanding its long-term potential. It may indicate a period of reflection and patience.",
		"In career, this card represents reviewing your work, assessing growth, and considering long-term goals. It’s a time to be patient and allow your efforts to bear fruit.",
	},

	"Eight of Pentacles": {
		"Eight of Pentacles",
		"🛠️ Eight of Pentacles – Mastery & Diligence",
		"The Eight of Pentacles symbolizes hard work, dedication, and the pursuit of mastery. It represents the commitment to honing your craft and improving your skills through persistence. This card encourages you to stay focused, detail-oriented, and disciplined in your efforts.",
		"In love, the Eight of Pentacles suggests working hard to improve the relationship, paying attention to the details that strengthen the bond. It’s a sign of dedication and commitment.",
		"In career, this card represents focused effort, skill-building, and attention to quality. It’s a time to perfect your craft and take pride in your work.",
	},

	"Nine of Pentacles": {
		"Nine of Pentacles",
		"🌻 Nine of Pentacles – Abundance & Independence",
		"The Nine of Pentacles symbolizes self-sufficiency, luxury, and the rewards of hard work. It represents the enjoyment of material wealth, personal achievement, and the confidence that comes from independence. This card encourages you to take pride in your accomplishments and enjoy the fruits of your labor.",
		"In love, the Nine of Pentacles suggests stability and self-assuredness within the relationship. It indicates mutual respect and the ability to thrive both together and independently.",
		"In career, this card represents financial independence, successful projects, and the attainment of long-term goals. It’s a sign that your efforts have brought you to a place of comfort and security.",
	},

	"Ten of Pentacles": {
		"Ten of Pentacles",
		"🏡 Ten of Pentacles – Legacy & Stability",
		"The Ten of Pentacles symbolizes long-term stability, wealth, and family legacy. It represents the culmination of hard work and the establishment of strong foundations that last for generations. This card encourages you to focus on building lasting success and cherishing connections with family and community.",
		"In love, the Ten of Pentacles suggests a deep commitment, long-term stability, and building a future together. It may indicate marriage, family growth, or the strengthening of bonds.",
		"In career, this card represents long-term success, financial security, and achievements that have lasting impact. It’s a time to focus on sustainability and future planning.",
	},

	"Page of Pentacles": {
		"Page of Pentacles",
		"📜 Page of Pentacles – Opportunity & Learning",
		"The Page of Pentacles symbolizes new opportunities, growth, and the exploration of practical skills. It represents the beginning stages of a financial or career endeavor, inviting you to learn, plan, and build for the future. This card encourages you to stay curious and dedicated as you explore new ventures.",
		"In love, the Page of Pentacles suggests the beginning of a stable and grounded connection. It represents learning about each other and building a solid foundation for the future.",
		"In career, this card represents the pursuit of new projects, educational growth, and financial planning. It’s a time to focus on developing skills and preparing for success.",
	},

	"Knight of Pentacles": {
		"Knight of Pentacles",
		"🐂 Knight of Pentacles – Diligence & Persistence",
		"The Knight of Pentacles symbolizes hard work, reliability, and methodical progress. He represents the dedication to achieve goals through persistence and careful planning. This card encourages you to stay committed to your path, even if the progress seems slow and steady.",
		"In love, the Knight of Pentacles suggests loyalty, reliability, and a slow but steady approach to building trust. It’s a sign of commitment and lasting stability.",
		"In career, this card represents hard work, consistency, and attention to detail. It’s a time to stay focused, be patient, and trust that your efforts will pay off in the long term.",
	},

	"Queen of Pentacles": {
		"Queen of Pentacles",
		"👑 Queen of Pentacles – Nurturing & Security",
		"The Queen of Pentacles symbolizes practicality, nurturing, and the ability to create abundance through grounded effort. She represents the balance of material success with care and generosity. This card encourages you to cultivate both financial security and a warm, welcoming environment.",
		"In love, the Queen of Pentacles suggests stability, nurturing, and deep care within the relationship. It symbolizes a partner who is supportive, loyal, and dedicated to building a loving home.",
		"In career, this card represents successful management of resources, financial security, and the ability to nurture projects into success. It’s a time to trust your instincts and invest in long-term growth.",
	},

	"King of Pentacles": {
		"King of Pentacles",
		"👑 King of Pentacles – Wealth & Stability",
		"The King of Pentacles symbolizes financial mastery, stability, and the successful management of wealth. He represents abundance achieved through patience, dedication, and strategic planning. This card encourages you to take pride in your accomplishments and continue building on your secure foundations.",
		"In love, the King of Pentacles suggests loyalty, commitment, and a focus on building a strong and secure relationship. It represents a partner who is protective, generous, and deeply invested in long-term happiness.",
		"In career, this card represents business success, financial security, and the wise management of resources. It’s a sign of achieving your goals through discipline and consistency.",
	},
}
